// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var things_pb = require(require("path").join(__dirname, './things_pb.js'));
var google_api_annotations_pb = require(require("path").join(__dirname, './google/api/annotations_pb.js'));
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_things_Thing(arg) {
  if (!(arg instanceof things_pb.Thing)) {
    throw new Error('Expected argument of type things.Thing');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_things_Thing(buffer_arg) {
  return things_pb.Thing.deserializeBinary(new Uint8Array(buffer_arg));
}


var ThingsService = exports.ThingsService = {
  get: {
    path: '/things.Things/Get',
    requestStream: false,
    responseStream: false,
    requestType: things_pb.Thing,
    responseType: things_pb.Thing,
    requestSerialize: serialize_things_Thing,
    requestDeserialize: deserialize_things_Thing,
    responseSerialize: serialize_things_Thing,
    responseDeserialize: deserialize_things_Thing,
  },
  post: {
    path: '/things.Things/Post',
    requestStream: false,
    responseStream: false,
    requestType: things_pb.Thing,
    responseType: things_pb.Thing,
    requestSerialize: serialize_things_Thing,
    requestDeserialize: deserialize_things_Thing,
    responseSerialize: serialize_things_Thing,
    responseDeserialize: deserialize_things_Thing,
  },
};

exports.ThingsClient = grpc.makeGenericClientConstructor(ThingsService);
