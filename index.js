'use strict';

const grpc = require('grpc');
const _ = require('lodash');
const { ThingsService } = require('./js/things_grpc_pb.js');
const { Thing } = require('./js/things_pb.js');

let things = [
  {
    id: 1,
    name: 'test thing 1'
  },
  {
    id: 2,
    name: 'test thing 2'
  }
];

let get = (call, done) => {
  let thing = call.request;
  let id = thing.getId();
  let foundThing = _.find(things, {id});
  thing.setName(foundThing.name);
  return done(null, thing);
};

// createServer creates a new grpc Server with EDI endpoint handlers
let server = new grpc.Server();
server.addService(ThingsService, {
  get: get,  
});

server.bind('0.0.0.0:23478', grpc.ServerCredentials.createInsecure());
console.log('Server started!');
server.start();
