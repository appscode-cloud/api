// Code generated by protoc-gen-grpc-js-client
// source: database.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function databasesList(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function databasesCreate(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases'
    delete p['cluster']
    return xhr(path, 'POST', conf, null, p);
}

function databasesScale(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/actions/scale'
    delete p['cluster']
    delete p['uid']
    return xhr(path, 'PUT', conf, null, p);
}

function databasesUpdate(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid']
    delete p['cluster']
    delete p['uid']
    return xhr(path, 'PUT', conf, null, p);
}

function databasesDescribe(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid']
    delete p['cluster']
    delete p['uid']
    return xhr(path, 'GET', conf, p);
}

function databasesDelete(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid']
    delete p['cluster']
    delete p['uid']
    return xhr(path, 'DELETE', conf, p);
}

module.exports = {
    databases: {
        list: databasesList,
        create: databasesCreate,
        scale: databasesScale,
        update: databasesUpdate,
        describe: databasesDescribe,
        delete: databasesDelete
    }
};
