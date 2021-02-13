// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var gmailapi_gmailapi_pb = require('../gmailapi/gmailapi_pb.js');

function serialize_gmailapi_ClientSentimentRequest(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.ClientSentimentRequest)) {
    throw new Error('Expected argument of type notificationservice.ClientSentimentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_ClientSentimentRequest(buffer_arg) {
  return gmailapi_gmailapi_pb.ClientSentimentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gmailapi_ClientSentimentResponse(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.ClientSentimentResponse)) {
    throw new Error('Expected argument of type notificationservice.ClientSentimentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_ClientSentimentResponse(buffer_arg) {
  return gmailapi_gmailapi_pb.ClientSentimentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gmailapi_ConfirmationResponse(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.ConfirmationResponse)) {
    throw new Error('Expected argument of type notificationservice.ConfirmationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_ConfirmationResponse(buffer_arg) {
  return gmailapi_gmailapi_pb.ConfirmationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gmailapi_LightStreamerSubRequest(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.LightStreamerSubRequest)) {
    throw new Error('Expected argument of type notificationservice.LightStreamerSubRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_LightStreamerSubRequest(buffer_arg) {
  return gmailapi_gmailapi_pb.LightStreamerSubRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gmailapi_LightStreamerSubResponse(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.LightStreamerSubResponse)) {
    throw new Error('Expected argument of type notificationservice.LightStreamerSubResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_LightStreamerSubResponse(buffer_arg) {
  return gmailapi_gmailapi_pb.LightStreamerSubResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gmailapi_OTCOrderRequest(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.OTCOrderRequest)) {
    throw new Error('Expected argument of type notificationservice.OTCOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_OTCOrderRequest(buffer_arg) {
  return gmailapi_gmailapi_pb.OTCOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gmailapi_OTCOrderResponse(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.OTCOrderResponse)) {
    throw new Error('Expected argument of type notificationservice.OTCOrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_OTCOrderResponse(buffer_arg) {
  return gmailapi_gmailapi_pb.OTCOrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gmailapi_SubscriptionRequest(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.SubscriptionRequest)) {
    throw new Error('Expected argument of type notificationservice.SubscriptionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_SubscriptionRequest(buffer_arg) {
  return gmailapi_gmailapi_pb.SubscriptionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gmailapi_SubscriptionResponse(arg) {
  if (!(arg instanceof gmailapi_gmailapi_pb.SubscriptionResponse)) {
    throw new Error('Expected argument of type notificationservice.SubscriptionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gmailapi_SubscriptionResponse(buffer_arg) {
  return gmailapi_gmailapi_pb.SubscriptionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var GmailapiSvcService = exports.GmailapiSvcService = {
  subscribeToMail: {
    path: '/notificationservice.GmailapiSvc/subscribeToMail',
    requestStream: false,
    responseStream: false,
    requestType: gmailapi_gmailapi_pb.SubscriptionRequest,
    responseType: gmailapi_gmailapi_pb.SubscriptionResponse,
    requestSerialize: serialize_gmailapi_SubscriptionRequest,
    requestDeserialize: deserialize_gmailapi_SubscriptionRequest,
    responseSerialize: serialize_gmailapi_SubscriptionResponse,
    responseDeserialize: deserialize_gmailapi_SubscriptionResponse,
  },
  openOTCOrder: {
    path: '/notificationservice.GmailapiSvc/openOTCOrder',
    requestStream: false,
    responseStream: false,
    requestType: gmailapi_gmailapi_pb.OTCOrderRequest,
    responseType: gmailapi_gmailapi_pb.OTCOrderResponse,
    requestSerialize: serialize_gmailapi_OTCOrderRequest,
    requestDeserialize: deserialize_gmailapi_OTCOrderRequest,
    responseSerialize: serialize_gmailapi_OTCOrderResponse,
    responseDeserialize: deserialize_gmailapi_OTCOrderResponse,
  },
  getConfirmationDetails: {
    path: '/notificationservice.GmailapiSvc/getConfirmationDetails',
    requestStream: false,
    responseStream: false,
    requestType: gmailapi_gmailapi_pb.OTCOrderResponse,
    responseType: gmailapi_gmailapi_pb.ConfirmationResponse,
    requestSerialize: serialize_gmailapi_OTCOrderResponse,
    requestDeserialize: deserialize_gmailapi_OTCOrderResponse,
    responseSerialize: serialize_gmailapi_ConfirmationResponse,
    responseDeserialize: deserialize_gmailapi_ConfirmationResponse,
  },
  getClientSentiment: {
    path: '/notificationservice.GmailapiSvc/getClientSentiment',
    requestStream: false,
    responseStream: false,
    requestType: gmailapi_gmailapi_pb.ClientSentimentRequest,
    responseType: gmailapi_gmailapi_pb.ClientSentimentResponse,
    requestSerialize: serialize_gmailapi_ClientSentimentRequest,
    requestDeserialize: deserialize_gmailapi_ClientSentimentRequest,
    responseSerialize: serialize_gmailapi_ClientSentimentResponse,
    responseDeserialize: deserialize_gmailapi_ClientSentimentResponse,
  },
  openLightStreamerSubscription: {
    path: '/notificationservice.GmailapiSvc/openLightStreamerSubscription',
    requestStream: false,
    responseStream: true,
    requestType: gmailapi_gmailapi_pb.LightStreamerSubRequest,
    responseType: gmailapi_gmailapi_pb.LightStreamerSubResponse,
    requestSerialize: serialize_gmailapi_LightStreamerSubRequest,
    requestDeserialize: deserialize_gmailapi_LightStreamerSubRequest,
    responseSerialize: serialize_gmailapi_LightStreamerSubResponse,
    responseDeserialize: deserialize_gmailapi_LightStreamerSubResponse,
  },
};

exports.GmailapiSvcClient = grpc.makeGenericClientConstructor(GmailapiSvcService);
