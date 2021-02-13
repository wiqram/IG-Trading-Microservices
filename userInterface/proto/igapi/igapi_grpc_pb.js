// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var igapi_igapi_pb = require('../igapi/igapi_pb.js');

function serialize_demo_igapi_OpenTradeRequest(arg) {
  if (!(arg instanceof igapi_igapi_pb.OpenTradeRequest)) {
    throw new Error('Expected argument of type demo_igapi.OpenTradeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_demo_igapi_OpenTradeRequest(buffer_arg) {
  return igapi_igapi_pb.OpenTradeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_demo_igapi_TradeResponse(arg) {
  if (!(arg instanceof igapi_igapi_pb.TradeResponse)) {
    throw new Error('Expected argument of type demo_igapi.TradeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_demo_igapi_TradeResponse(buffer_arg) {
  return igapi_igapi_pb.TradeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var IgSvcService = exports.IgSvcService = {
  tradeAction: {
    path: '/demo_igapi.IgSvc/tradeAction',
    requestStream: false,
    responseStream: false,
    requestType: igapi_igapi_pb.OpenTradeRequest,
    responseType: igapi_igapi_pb.TradeResponse,
    requestSerialize: serialize_demo_igapi_OpenTradeRequest,
    requestDeserialize: deserialize_demo_igapi_OpenTradeRequest,
    responseSerialize: serialize_demo_igapi_TradeResponse,
    responseDeserialize: deserialize_demo_igapi_TradeResponse,
  },
};

exports.IgSvcClient = grpc.makeGenericClientConstructor(IgSvcService);
