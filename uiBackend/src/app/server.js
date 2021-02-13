const app = require('./app');
/*
require('dotenv').config();
const grpc = require('@grpc/grpc-js');
const services = require('../proto/uibackend_grpc_pb');

const server = new grpc.Server();
server.addService(services.UibackendSvcService, {
  getProducts: app.getProducts,
  getTransactions: app.getTransactions,
});
let address = process.env.HOST + ":" + process.env.PORT;
server.bindAsync(address, grpc.ServerCredentials.createInsecure(), () => {
  server.start();
  console.log("Server running at " + address);
});
*/



module.exports = app.listen(3070, err => {
  if (err) return console.error(`Failed to start server`, err);
  console.log(`GRPC Server started on  http://localhost:3070`);
});

