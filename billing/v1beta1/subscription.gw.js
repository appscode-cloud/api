// Code generated by protoc-gen-grpc-js-client
// source: subscription.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function subscriptionsSubscribe(p, conf) {
    path = '/billing/v1beta1/subscriptions'
    return xhr(path, 'POST', conf, null, p);
}

module.exports = {
    subscriptions: {
        subscribe: subscriptionsSubscribe
    }
};
