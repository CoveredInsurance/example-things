'use strict';

const grpc = require('grpc');

const { ThingsClient } = require('./js/things_grpc_pb.js');
const { Thing } = require('./js/things_pb.js');

let client = new ThingsClient('0.0.0.0:23478', grpc.credentials.createInsecure());
let newThing = new Thing();
newThing.setId(2);

client.get(newThing, (err, res) => {
  console.log(err, res.toObject());
});
