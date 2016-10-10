// Code generated by protoc-gen-grpc-js-client
// source: agent.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function AgentsList(p, conf) {
    path = '/ci/v1beta1/agents'
    return xhr(path, 'GET', conf, p);
}

function AgentsDescribe(p, conf) {
    path = '/ci/v1beta1/agents/' + p['name']
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

function AgentsCreate(p, conf) {
    path = '/ci/v1beta1/agents'
    return xhr(path, 'POST', conf, null, p);
}

function AgentsDelete(p, conf) {
    path = '/ci/v1beta1/agents/' + p['name']
    delete p['name']
    return xhr(path, 'DELETE', conf, p);
}

function AgentsRestart(p, conf) {
    path = '/ci/v1beta1/agents/' + p['name'] + '/actions/reboot'
    delete p['name']
    return xhr(path, 'POST', conf, p);
}

module.exports = {
    Agents: {
        List: AgentsList,
        Describe: AgentsDescribe,
        Create: AgentsCreate,
        Delete: AgentsDelete,
        Restart: AgentsRestart
    }
};
