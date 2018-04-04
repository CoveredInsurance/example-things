'use strict';

const grpc = require('grpc');

const { ThingsService } = require('./js/things_grpc_pb.js');
const { Thing } = require('./js/things_pb.js');

let things = [
  {
    id: 1,
    name: 'Test Name',
    desc: 'Test Desc...'
  },
  {
    id: 2,
    name: 'Test Name2',
    desc: 'Test Desc2...'
  },
]

let get = (call, done) => {
  let thingRs = call.request;
  let id = thingRs.getId()
  let thingMatch;
  
  for (let thing of things) {
    if (thing.id === id) {
      thingMatch = thing;
    }
  }
  thingRs.setId(thingMatch.id);
  thingRs.setName(thingMatch.name);
  thingRs.setDesc(thingMatch.desc);
  return done(null, thingRs)
};

// createServer creates a new grpc Server with EDI endpoint handlers
let server = new grpc.Server();
server.addService(ThingsService, {
  get: get,
});

server.bind('0.0.0.0:23478', grpc.ServerCredentials.createInsecure());
console.log('Server started!!!!');
server.start();
