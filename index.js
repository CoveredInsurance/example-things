'use strict';

const grpc = require('grpc');
const _ = require('lodash');
const { ThingsService } = require('./js/things_grpc_pb.js');
const { Thing } = require('./js/things_pb.js');

let get = (call, done) => {
  return done(null, call.request);
};

// createServer creates a new grpc Server with EDI endpoint handlers
let server = new grpc.Server();
server.addService(ThingsService, {
  get: get,
});

server.bind('0.0.0.0:23478', grpc.ServerCredentials.createInsecure());
console.log('Server started!');
server.start();
