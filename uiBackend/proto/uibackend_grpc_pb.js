// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var uibackend_pb = require('./uibackend_pb.js');

function serialize_demo_uibackend_ProductsRequest(arg) {
  if (!(arg instanceof uibackend_pb.ProductsRequest)) {
    throw new Error('Expected argument of type demo_uibackend.ProductsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_demo_uibackend_ProductsRequest(buffer_arg) {
  return uibackend_pb.ProductsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_demo_uibackend_ProductsResponse(arg) {
  if (!(arg instanceof uibackend_pb.ProductsResponse)) {
    throw new Error('Expected argument of type demo_uibackend.ProductsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_demo_uibackend_ProductsResponse(buffer_arg) {
  return uibackend_pb.ProductsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_demo_uibackend_TransactionsRequest(arg) {
  if (!(arg instanceof uibackend_pb.TransactionsRequest)) {
    throw new Error('Expected argument of type demo_uibackend.TransactionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_demo_uibackend_TransactionsRequest(buffer_arg) {
  return uibackend_pb.TransactionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_demo_uibackend_TransactionsResponse(arg) {
  if (!(arg instanceof uibackend_pb.TransactionsResponse)) {
    throw new Error('Expected argument of type demo_uibackend.TransactionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_demo_uibackend_TransactionsResponse(buffer_arg) {
  return uibackend_pb.TransactionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var UibackendSvcService = exports.UibackendSvcService = {
  getProducts: {
    path: '/demo_uibackend.UibackendSvc/getProducts',
    requestStream: false,
    responseStream: false,
    requestType: uibackend_pb.ProductsRequest,
    responseType: uibackend_pb.ProductsResponse,
    requestSerialize: serialize_demo_uibackend_ProductsRequest,
    requestDeserialize: deserialize_demo_uibackend_ProductsRequest,
    responseSerialize: serialize_demo_uibackend_ProductsResponse,
    responseDeserialize: deserialize_demo_uibackend_ProductsResponse,
  },
  getTransactions: {
    path: '/demo_uibackend.UibackendSvc/getTransactions',
    requestStream: false,
    responseStream: false,
    requestType: uibackend_pb.TransactionsRequest,
    responseType: uibackend_pb.TransactionsResponse,
    requestSerialize: serialize_demo_uibackend_TransactionsRequest,
    requestDeserialize: deserialize_demo_uibackend_TransactionsRequest,
    responseSerialize: serialize_demo_uibackend_TransactionsResponse,
    responseDeserialize: deserialize_demo_uibackend_TransactionsResponse,
  },
};

exports.UibackendSvcClient = grpc.makeGenericClientConstructor(UibackendSvcService);
