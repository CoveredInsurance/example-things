'use strict';

const grpc = require('grpc');

const { ThingsService } = require('./js/things_grpc_pb.js');


let get = (call, done) => {
  return done({
    code: grpc.status.UNIMPLEMENTED,
    message: 'Unimplemented'
  });
};

// createServer creates a new grpc Server with EDI endpoint handlers
let server = new grpc.Server();
server.addService(ThingsService, {
  get: get
});

server.bind('0.0.0.0:23478', grpc.ServerCredentials.createInsecure());
server.start();
