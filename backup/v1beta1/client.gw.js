// Code generated by protoc-gen-grpc-js-client
// source: client.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function clientsReconfigure(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/apps/bacula/actions/reconfigure'
    delete p['cluster']
    return xhr(path, 'PUT', conf, null, p);
}

module.exports = {
    clients: {
        reconfigure: clientsReconfigure
    }
};
