// Code generated by protoc-gen-grpc-js-client
// source: incident.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function IncidentsList(p, conf) {
    path = '/kubernetes/v1beta1/incidents'
    return xhr(path, 'GET', conf, p);
}

function IncidentsDescribe(p, conf) {
    path = '/kubernetes/v1beta1/incidents/' + p['phid']
    delete p['phid']
    return xhr(path, 'GET', conf, p);
}

function IncidentsNotify(p, conf) {
    path = '/kubernetes/v1beta1/alerts/' + p['alert_phid'] + '/actions/notify'
    delete p['alert_phid']
    return xhr(path, 'POST', conf, null, p);
}

function IncidentsCreateEvent(p, conf) {
    path = '/kubernetes/v1beta1/incidents/' + p['phid'] + '/events'
    delete p['phid']
    return xhr(path, 'POST', conf, null, p);
}

module.exports = {
    Incidents: {
        List: IncidentsList,
        Describe: IncidentsDescribe,
        Notify: IncidentsNotify,
        CreateEvent: IncidentsCreateEvent
    }
};
