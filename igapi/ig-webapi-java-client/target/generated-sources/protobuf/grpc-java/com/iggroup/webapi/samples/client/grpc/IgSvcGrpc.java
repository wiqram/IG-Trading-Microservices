package com.iggroup.webapi.samples.client.grpc;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.4.0)",
    comments = "Source: igapi.proto")
public final class IgSvcGrpc {

  private IgSvcGrpc() {}

  public static final String SERVICE_NAME = "demo_igapi.IgSvc";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.OpenTradeRequest,
      com.iggroup.webapi.samples.client.grpc.TradeResponse> METHOD_TRADE_ACTION =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.OpenTradeRequest, com.iggroup.webapi.samples.client.grpc.TradeResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "tradeAction"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.OpenTradeRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TradeResponse.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.TextRazorRequest,
      com.iggroup.webapi.samples.client.grpc.TextRazorResponse> METHOD_TEST_ACCOUNT_TEXT_RAZOR =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.TextRazorRequest, com.iggroup.webapi.samples.client.grpc.TextRazorResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "testAccountTextRazor"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TextRazorRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TextRazorResponse.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.TextRazorRequest,
      com.iggroup.webapi.samples.client.grpc.TextRazorResponse> METHOD_TEST_CLASSIFIER_TEXT_RAZOR =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.TextRazorRequest, com.iggroup.webapi.samples.client.grpc.TextRazorResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "testClassifierTextRazor"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TextRazorRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TextRazorResponse.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.TextRazorRequest,
      com.iggroup.webapi.samples.client.grpc.TextRazorResponse> METHOD_TEST_DICTIONARY_TEXT_RAZOR =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.TextRazorRequest, com.iggroup.webapi.samples.client.grpc.TextRazorResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "testDictionaryTextRazor"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TextRazorRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TextRazorResponse.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.TextRazorRequest,
      com.iggroup.webapi.samples.client.grpc.TextRazorResponse> METHOD_TEST_ANALYSIS_TEXT_RAZOR =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.TextRazorRequest, com.iggroup.webapi.samples.client.grpc.TextRazorResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "testAnalysisTextRazor"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TextRazorRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.TextRazorResponse.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.OTCOrderRequest,
      com.iggroup.webapi.samples.client.grpc.OTCOrderResponse> METHOD_CREATE_POSITION =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.OTCOrderRequest, com.iggroup.webapi.samples.client.grpc.OTCOrderResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "createPosition"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.OTCOrderRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.OTCOrderResponse.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.OTCOrderResponse,
      com.iggroup.webapi.samples.client.grpc.ConfirmationResponse> METHOD_GET_CONFIRMATION_DETAILS =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.OTCOrderResponse, com.iggroup.webapi.samples.client.grpc.ConfirmationResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "getConfirmationDetails"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.OTCOrderResponse.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.ConfirmationResponse.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest,
      com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse> METHOD_GET_CLIENT_SENTIMENT =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest, com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "getClientSentiment"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.iggroup.webapi.samples.client.grpc.LightStreamerSubRequest,
      com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse> METHOD_OPEN_LIGHT_STREAMER_SUBSCRIPTION =
      io.grpc.MethodDescriptor.<com.iggroup.webapi.samples.client.grpc.LightStreamerSubRequest, com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
          .setFullMethodName(generateFullMethodName(
              "demo_igapi.IgSvc", "openLightStreamerSubscription"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.LightStreamerSubRequest.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse.getDefaultInstance()))
          .build();

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static IgSvcStub newStub(io.grpc.Channel channel) {
    return new IgSvcStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static IgSvcBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new IgSvcBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static IgSvcFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new IgSvcFutureStub(channel);
  }

  /**
   */
  public static abstract class IgSvcImplBase implements io.grpc.BindableService {

    /**
     */
    public void tradeAction(com.iggroup.webapi.samples.client.grpc.OpenTradeRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TradeResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_TRADE_ACTION, responseObserver);
    }

    /**
     */
    public void testAccountTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_TEST_ACCOUNT_TEXT_RAZOR, responseObserver);
    }

    /**
     */
    public void testClassifierTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_TEST_CLASSIFIER_TEXT_RAZOR, responseObserver);
    }

    /**
     */
    public void testDictionaryTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_TEST_DICTIONARY_TEXT_RAZOR, responseObserver);
    }

    /**
     */
    public void testAnalysisTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_TEST_ANALYSIS_TEXT_RAZOR, responseObserver);
    }

    /**
     */
    public void createPosition(com.iggroup.webapi.samples.client.grpc.OTCOrderRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.OTCOrderResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_CREATE_POSITION, responseObserver);
    }

    /**
     */
    public void getConfirmationDetails(com.iggroup.webapi.samples.client.grpc.OTCOrderResponse request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.ConfirmationResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_CONFIRMATION_DETAILS, responseObserver);
    }

    /**
     */
    public void getClientSentiment(com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_CLIENT_SENTIMENT, responseObserver);
    }

    /**
     */
    public void openLightStreamerSubscription(com.iggroup.webapi.samples.client.grpc.LightStreamerSubRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_OPEN_LIGHT_STREAMER_SUBSCRIPTION, responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            METHOD_TRADE_ACTION,
            asyncUnaryCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.OpenTradeRequest,
                com.iggroup.webapi.samples.client.grpc.TradeResponse>(
                  this, METHODID_TRADE_ACTION)))
          .addMethod(
            METHOD_TEST_ACCOUNT_TEXT_RAZOR,
            asyncUnaryCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.TextRazorRequest,
                com.iggroup.webapi.samples.client.grpc.TextRazorResponse>(
                  this, METHODID_TEST_ACCOUNT_TEXT_RAZOR)))
          .addMethod(
            METHOD_TEST_CLASSIFIER_TEXT_RAZOR,
            asyncUnaryCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.TextRazorRequest,
                com.iggroup.webapi.samples.client.grpc.TextRazorResponse>(
                  this, METHODID_TEST_CLASSIFIER_TEXT_RAZOR)))
          .addMethod(
            METHOD_TEST_DICTIONARY_TEXT_RAZOR,
            asyncUnaryCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.TextRazorRequest,
                com.iggroup.webapi.samples.client.grpc.TextRazorResponse>(
                  this, METHODID_TEST_DICTIONARY_TEXT_RAZOR)))
          .addMethod(
            METHOD_TEST_ANALYSIS_TEXT_RAZOR,
            asyncUnaryCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.TextRazorRequest,
                com.iggroup.webapi.samples.client.grpc.TextRazorResponse>(
                  this, METHODID_TEST_ANALYSIS_TEXT_RAZOR)))
          .addMethod(
            METHOD_CREATE_POSITION,
            asyncUnaryCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.OTCOrderRequest,
                com.iggroup.webapi.samples.client.grpc.OTCOrderResponse>(
                  this, METHODID_CREATE_POSITION)))
          .addMethod(
            METHOD_GET_CONFIRMATION_DETAILS,
            asyncUnaryCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.OTCOrderResponse,
                com.iggroup.webapi.samples.client.grpc.ConfirmationResponse>(
                  this, METHODID_GET_CONFIRMATION_DETAILS)))
          .addMethod(
            METHOD_GET_CLIENT_SENTIMENT,
            asyncUnaryCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest,
                com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse>(
                  this, METHODID_GET_CLIENT_SENTIMENT)))
          .addMethod(
            METHOD_OPEN_LIGHT_STREAMER_SUBSCRIPTION,
            asyncServerStreamingCall(
              new MethodHandlers<
                com.iggroup.webapi.samples.client.grpc.LightStreamerSubRequest,
                com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse>(
                  this, METHODID_OPEN_LIGHT_STREAMER_SUBSCRIPTION)))
          .build();
    }
  }

  /**
   */
  public static final class IgSvcStub extends io.grpc.stub.AbstractStub<IgSvcStub> {
    private IgSvcStub(io.grpc.Channel channel) {
      super(channel);
    }

    private IgSvcStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected IgSvcStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new IgSvcStub(channel, callOptions);
    }

    /**
     */
    public void tradeAction(com.iggroup.webapi.samples.client.grpc.OpenTradeRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TradeResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_TRADE_ACTION, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void testAccountTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_TEST_ACCOUNT_TEXT_RAZOR, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void testClassifierTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_TEST_CLASSIFIER_TEXT_RAZOR, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void testDictionaryTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_TEST_DICTIONARY_TEXT_RAZOR, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void testAnalysisTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_TEST_ANALYSIS_TEXT_RAZOR, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createPosition(com.iggroup.webapi.samples.client.grpc.OTCOrderRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.OTCOrderResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_CREATE_POSITION, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getConfirmationDetails(com.iggroup.webapi.samples.client.grpc.OTCOrderResponse request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.ConfirmationResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_CONFIRMATION_DETAILS, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getClientSentiment(com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_CLIENT_SENTIMENT, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void openLightStreamerSubscription(com.iggroup.webapi.samples.client.grpc.LightStreamerSubRequest request,
        io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse> responseObserver) {
      asyncServerStreamingCall(
          getChannel().newCall(METHOD_OPEN_LIGHT_STREAMER_SUBSCRIPTION, getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class IgSvcBlockingStub extends io.grpc.stub.AbstractStub<IgSvcBlockingStub> {
    private IgSvcBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private IgSvcBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected IgSvcBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new IgSvcBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.iggroup.webapi.samples.client.grpc.TradeResponse tradeAction(com.iggroup.webapi.samples.client.grpc.OpenTradeRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_TRADE_ACTION, getCallOptions(), request);
    }

    /**
     */
    public com.iggroup.webapi.samples.client.grpc.TextRazorResponse testAccountTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_TEST_ACCOUNT_TEXT_RAZOR, getCallOptions(), request);
    }

    /**
     */
    public com.iggroup.webapi.samples.client.grpc.TextRazorResponse testClassifierTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_TEST_CLASSIFIER_TEXT_RAZOR, getCallOptions(), request);
    }

    /**
     */
    public com.iggroup.webapi.samples.client.grpc.TextRazorResponse testDictionaryTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_TEST_DICTIONARY_TEXT_RAZOR, getCallOptions(), request);
    }

    /**
     */
    public com.iggroup.webapi.samples.client.grpc.TextRazorResponse testAnalysisTextRazor(com.iggroup.webapi.samples.client.grpc.TextRazorRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_TEST_ANALYSIS_TEXT_RAZOR, getCallOptions(), request);
    }

    /**
     */
    public com.iggroup.webapi.samples.client.grpc.OTCOrderResponse createPosition(com.iggroup.webapi.samples.client.grpc.OTCOrderRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_CREATE_POSITION, getCallOptions(), request);
    }

    /**
     */
    public com.iggroup.webapi.samples.client.grpc.ConfirmationResponse getConfirmationDetails(com.iggroup.webapi.samples.client.grpc.OTCOrderResponse request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_CONFIRMATION_DETAILS, getCallOptions(), request);
    }

    /**
     */
    public com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse getClientSentiment(com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_CLIENT_SENTIMENT, getCallOptions(), request);
    }

    /**
     */
    public java.util.Iterator<com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse> openLightStreamerSubscription(
        com.iggroup.webapi.samples.client.grpc.LightStreamerSubRequest request) {
      return blockingServerStreamingCall(
          getChannel(), METHOD_OPEN_LIGHT_STREAMER_SUBSCRIPTION, getCallOptions(), request);
    }
  }

  /**
   */
  public static final class IgSvcFutureStub extends io.grpc.stub.AbstractStub<IgSvcFutureStub> {
    private IgSvcFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private IgSvcFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected IgSvcFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new IgSvcFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.iggroup.webapi.samples.client.grpc.TradeResponse> tradeAction(
        com.iggroup.webapi.samples.client.grpc.OpenTradeRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_TRADE_ACTION, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> testAccountTextRazor(
        com.iggroup.webapi.samples.client.grpc.TextRazorRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_TEST_ACCOUNT_TEXT_RAZOR, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> testClassifierTextRazor(
        com.iggroup.webapi.samples.client.grpc.TextRazorRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_TEST_CLASSIFIER_TEXT_RAZOR, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> testDictionaryTextRazor(
        com.iggroup.webapi.samples.client.grpc.TextRazorRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_TEST_DICTIONARY_TEXT_RAZOR, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.iggroup.webapi.samples.client.grpc.TextRazorResponse> testAnalysisTextRazor(
        com.iggroup.webapi.samples.client.grpc.TextRazorRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_TEST_ANALYSIS_TEXT_RAZOR, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.iggroup.webapi.samples.client.grpc.OTCOrderResponse> createPosition(
        com.iggroup.webapi.samples.client.grpc.OTCOrderRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_CREATE_POSITION, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.iggroup.webapi.samples.client.grpc.ConfirmationResponse> getConfirmationDetails(
        com.iggroup.webapi.samples.client.grpc.OTCOrderResponse request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_CONFIRMATION_DETAILS, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse> getClientSentiment(
        com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_CLIENT_SENTIMENT, getCallOptions()), request);
    }
  }

  private static final int METHODID_TRADE_ACTION = 0;
  private static final int METHODID_TEST_ACCOUNT_TEXT_RAZOR = 1;
  private static final int METHODID_TEST_CLASSIFIER_TEXT_RAZOR = 2;
  private static final int METHODID_TEST_DICTIONARY_TEXT_RAZOR = 3;
  private static final int METHODID_TEST_ANALYSIS_TEXT_RAZOR = 4;
  private static final int METHODID_CREATE_POSITION = 5;
  private static final int METHODID_GET_CONFIRMATION_DETAILS = 6;
  private static final int METHODID_GET_CLIENT_SENTIMENT = 7;
  private static final int METHODID_OPEN_LIGHT_STREAMER_SUBSCRIPTION = 8;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final IgSvcImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(IgSvcImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_TRADE_ACTION:
          serviceImpl.tradeAction((com.iggroup.webapi.samples.client.grpc.OpenTradeRequest) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TradeResponse>) responseObserver);
          break;
        case METHODID_TEST_ACCOUNT_TEXT_RAZOR:
          serviceImpl.testAccountTextRazor((com.iggroup.webapi.samples.client.grpc.TextRazorRequest) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse>) responseObserver);
          break;
        case METHODID_TEST_CLASSIFIER_TEXT_RAZOR:
          serviceImpl.testClassifierTextRazor((com.iggroup.webapi.samples.client.grpc.TextRazorRequest) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse>) responseObserver);
          break;
        case METHODID_TEST_DICTIONARY_TEXT_RAZOR:
          serviceImpl.testDictionaryTextRazor((com.iggroup.webapi.samples.client.grpc.TextRazorRequest) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse>) responseObserver);
          break;
        case METHODID_TEST_ANALYSIS_TEXT_RAZOR:
          serviceImpl.testAnalysisTextRazor((com.iggroup.webapi.samples.client.grpc.TextRazorRequest) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.TextRazorResponse>) responseObserver);
          break;
        case METHODID_CREATE_POSITION:
          serviceImpl.createPosition((com.iggroup.webapi.samples.client.grpc.OTCOrderRequest) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.OTCOrderResponse>) responseObserver);
          break;
        case METHODID_GET_CONFIRMATION_DETAILS:
          serviceImpl.getConfirmationDetails((com.iggroup.webapi.samples.client.grpc.OTCOrderResponse) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.ConfirmationResponse>) responseObserver);
          break;
        case METHODID_GET_CLIENT_SENTIMENT:
          serviceImpl.getClientSentiment((com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse>) responseObserver);
          break;
        case METHODID_OPEN_LIGHT_STREAMER_SUBSCRIPTION:
          serviceImpl.openLightStreamerSubscription((com.iggroup.webapi.samples.client.grpc.LightStreamerSubRequest) request,
              (io.grpc.stub.StreamObserver<com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static final class IgSvcDescriptorSupplier implements io.grpc.protobuf.ProtoFileDescriptorSupplier {
    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.getDescriptor();
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (IgSvcGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new IgSvcDescriptorSupplier())
              .addMethod(METHOD_TRADE_ACTION)
              .addMethod(METHOD_TEST_ACCOUNT_TEXT_RAZOR)
              .addMethod(METHOD_TEST_CLASSIFIER_TEXT_RAZOR)
              .addMethod(METHOD_TEST_DICTIONARY_TEXT_RAZOR)
              .addMethod(METHOD_TEST_ANALYSIS_TEXT_RAZOR)
              .addMethod(METHOD_CREATE_POSITION)
              .addMethod(METHOD_GET_CONFIRMATION_DETAILS)
              .addMethod(METHOD_GET_CLIENT_SENTIMENT)
              .addMethod(METHOD_OPEN_LIGHT_STREAMER_SUBSCRIPTION)
              .build();
        }
      }
    }
    return result;
  }
}
