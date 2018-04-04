'use strict';

const grpc = require('grpc');

const { ThingsClient } = require('./js/things_grpc_pb.js');
const { Thing } = require('./js/things_pb.js');

let client = new ThingsClient('0.0.0.0:23478', grpc.credentials.createInsecure());

let thingRq = new Thing();
thingRq.setId(2)

client.get(thingRq, (err, res) => {
  console.log(err, res)
  let id = res.getId();
  console.log(id)
});
