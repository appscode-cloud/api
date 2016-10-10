// Code generated by protoc-gen-grpc-js-client
// source: authentication.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function authenticationLogin(p, conf) {
    path = '/auth/v1beta1/login'
    return xhr(path, 'POST', conf, null, p);
}

function authenticationLogout(p, conf) {
    path = '/auth/v1beta1/logout'
    return xhr(path, 'POST', conf, null, p);
}

function authenticationToken(p, conf) {
    path = '/auth/v1beta1/token'
    return xhr(path, 'POST', conf, null, p);
}

module.exports = {
    authentication: {
        login: authenticationLogin,
        logout: authenticationLogout,
        token: authenticationToken
    }
};
