#!/usr/bin/env python

import fnmatch
import json
import os
import re
import subprocess
import sys
from os.path import expandvars

# binary -> GOOS -> GOARCH
# windows	386
# windows	amd64

# Debian package
# https://gist.github.com/rcrowley/3728417

API_ROOT = expandvars("$GOPATH/src/github.com/appscode/api")
VALID_FORMATS = ['date-time',
                 'email',
                 'hostname',
                 'ipv4',
                 'ipv6',
                 'uri']
PKG_REGEX = re.compile(ur'^package\s*(?P<pkg>.*);$')


def call(cmd, stdin=None, cwd=API_ROOT):
    print(cmd)
    subprocess.call([expandvars(cmd)], shell=True, stdin=stdin, cwd=cwd)


def write_file(name, content):
    dir = os.path.dirname(name)
    if not os.path.exists(dir):
        os.makedirs(dir)
    with open(name, 'w') as f:
        return f.write(content)


def append_file(name, content):
    with open(name, 'a') as f:
        return f.write(content)


# TODO: use unicode encoding
def read_json(name):
    try:
        with open(name, 'r') as f:
            return json.load(f)
    except IOError:
        return {}


def write_json(obj, name):
    with open(name, 'w') as f:
        return json.dump(obj, f, sort_keys=True, indent=2, separators=(',', ': '))


def swagger_defs(defs):
    stack = []
    result = {}
    for key in defs.keys():
        if key.endswith("Request"):
            stack.append(key)
    while stack:
        name = stack.pop()
        schema = defs[name]
        result[name] = schema
        if 'properties' in schema:
            for p, v in schema['properties'].iteritems():
                if '$ref' in v:
                    nw_obj = v['$ref'][len('#/definitions/'):]
                    if nw_obj not in result:
                        stack.append(nw_obj)
                if 'format' in v and not v['format'] in VALID_FORMATS:
                    v.pop('format', None)
                if 'items' in v:
                    if '$ref' in v['items']:
                        nw_obj = v['items']['$ref'][len('#/definitions/'):]
                        if nw_obj not in result:
                            stack.append(nw_obj)
                    if 'format' in v['items'] and not v['items']['format'] in VALID_FORMATS:
                        v['items'].pop('format', None)
    return result


def generate_json_schema():
    for root, dirnames, filenames in os.walk(API_ROOT):
        for filename in fnmatch.filter(filenames, '*.swagger.json'):
            swagger = os.path.join(root, filename)
            schema = os.path.join(root, filename.replace('.swagger.', '.schema.'))
            print schema
            defs = swagger_defs(read_json(swagger)['definitions'])
            if os.path.exists(schema):
                defs.update(read_json(schema)['definitions'])
            write_json({'definitions': defs}, schema)


def schema_go(pkg, defs):
    result = {}
    for key in defs.keys():
        if key.startswith(pkg) and key.endswith("Request"):
            schema = defs[key]
            result[key[len(pkg):]] = schema
            dep_defs = {}
            stack = []
            stack.append(schema)
            while stack:
                sch = stack.pop()
                if 'properties' in sch:
                    for p, v in sch['properties'].iteritems():
                        if '$ref' in v:
                            nw_obj = v['$ref'][len('#/definitions/'):]
                            if nw_obj not in dep_defs:
                                dep_defs[nw_obj] = defs[nw_obj]
                                stack.append(defs[nw_obj])
                        if 'items' in v and '$ref' in v['items']:
                            nw_obj = v['items']['$ref'][len('#/definitions/'):]
                            if nw_obj not in dep_defs:
                                dep_defs[nw_obj] = defs[nw_obj]
                                stack.append(defs[nw_obj])
            if dep_defs:
                schema['definitions'] = dep_defs
            schema['$schema'] = 'http://json-schema.org/draft-04/schema#'
    return result


def render_schema_go(pkg, schemas):
    contents = """package {}

import (
	"github.com/xeipuuv/gojsonschema"
	"log"
)

// Auto-generated. DO NOT EDIT.

""".format(pkg)
    for key in schemas.keys():
        contents += 'var {0}Schema *gojsonschema.Schema\n'.format(key[0:1].lower() + key[1:])
    contents += '\n'

    contents += """func init() {
	var err error
"""
    for key, sch in schemas.iteritems():
        var_name = key[0:1].lower() + key[1:]
        sch_str = json.dumps(sch, sort_keys=True, indent=2, separators=(',', ': '))
        contents += '	{0}Schema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{1}`))\n'.format(
            var_name, sch_str)
        contents += """	if err != nil {
		log.Fatal(err)
	}
"""
    contents += '}\n\n'

    for key in schemas.keys():
        contents += 'func (m *' + key + ') IsValid() (*gojsonschema.Result, error) {\n'
        contents += '	return {}Schema.Validate(gojsonschema.NewGoLoader(m))\n'.format(key[0:1].lower() + key[1:])
        contents += '}\n'
        contents += 'func (m *' + key + ') IsRequest() {}\n\n'

    return contents


def detect_pkg(schema):
    proto = schema[:-len('.schema.json')] + '.proto'
    with open(proto) as f:
        for line in f:
            line = line.strip()
            m = re.search(PKG_REGEX, line)
            if m:
                return m.group('pkg')


def generate_go_schema():
    for root, dirnames, filenames in os.walk(API_ROOT):
        for filename in fnmatch.filter(filenames, '*.schema.json'):
            schema = os.path.join(root, filename)
            go = schema[:-len('.json')] + '.go'
            print go
            pkg = detect_pkg(schema)
            defs = read_json(schema)['definitions']
            schemas = schema_go(pkg, defs)
            if schemas:
                write_file(go, render_schema_go(pkg, schemas))


def apply_naming_policy():
    for root, dirnames, filenames in os.walk(API_ROOT):
        for filename in fnmatch.filter(filenames, '*.schema.json'):
            schema = os.path.join(root, filename)
            print schema
            content = read_json(schema)
            for key, defs in content['definitions'].iteritems():
                if 'properties' in defs:
                    for p, v in defs['properties'].iteritems():
                        if p in [
                            'cluster_name',
                            'namespace', 'name',
                            'bucket_name',
                            'secret_name',
                            'snapshot_name',
                            'auth_secret_name',
                            'cloud_credential'
                        ]:
                            print '====>>>> ' + p
                            if 'maxLength' not in v:
                                v['maxLength'] = 63
                            if 'pattern' not in v:
                                v['pattern'] = "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$"
            write_json(content, schema)


if __name__ == "__main__":
    if len(sys.argv) > 1:
        # http://stackoverflow.com/a/834451
        # http://stackoverflow.com/a/817296
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        generate_json_schema()
        # apply_naming_policy()
        generate_go_schema()
