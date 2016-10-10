// Code generated by protoc-gen-grpc-js-client
// source: purchase.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function PurchasesBegin(p, conf) {
    path = '/billing/v1beta1/purchases'
    return xhr(path, 'POST', conf, null, p);
}

function PurchasesComplete(p, conf) {
    path = '/billing/v1beta1/purchases/' + p['phid']
    delete p['phid']
    return xhr(path, 'PUT', conf, null, p);
}

function PurchasesClose(p, conf) {
    path = '/billing/v1beta1/purchases/' + p['object_phid']
    delete p['object_phid']
    return xhr(path, 'DELETE', conf, p);
}

module.exports = {
    Purchases: {
        Begin: PurchasesBegin,
        Complete: PurchasesComplete,
        Close: PurchasesClose
    }
};
