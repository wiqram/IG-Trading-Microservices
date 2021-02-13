package com.iggroup.webapi.samples.client.grpc;

import com.textrazor.AnalysisException;
import com.textrazor.NetworkException;
import io.grpc.stub.StreamObserver;
import com.iggroup.webapi.samples.client.TextRazorAPI;


public class IGAPIServiceImpl extends IgSvcGrpc.IgSvcImplBase {
    @Override
    public void tradeAction(OpenTradeRequest request,
                         StreamObserver<TradeResponse> responseObserver) {
        // HelloRequest has toString auto-generated.
        System.out.println(request);

        // You must use a builder to construct a new Protobuffer object
        TradeResponse response = TradeResponse.newBuilder()
                .setTransactionId("1234")
                .build();

        // Use responseObserver to send a single response back
        responseObserver.onNext(response);

        // When you are done, you must call onCompleted.
        responseObserver.onCompleted();
    }
    @Override
    public void testAccountTextRazor(TextRazorRequest request,
                            StreamObserver<TextRazorResponse> responseObserver) {
        // HelloRequest has toString auto-generated.
       /* System.out.println("this we are in test account for text razor -- this is the text coming up ---- ");
        System.out.println(request.getText());*/
        try {
            TextRazorAPI.testAccount(request.getApiKey(), request.getText());
        } catch (NetworkException e) {
            e.printStackTrace();
        } catch (AnalysisException e) {
            e.printStackTrace();
        }

        // You must use a builder to construct a new Protobuffer object

        // Use responseObserver to send a single response back

        // When you are done, you must call onCompleted.
        responseObserver.onCompleted();
    }

    public void testAnalysisTextRazor(TextRazorRequest request,
                                     StreamObserver<TextRazorResponse> responseObserver) {
        // HelloRequest has toString auto-generated.
        try {
            TextRazorAPI.testAnalysis(request.getApiKey(), request.getText());
        } catch (NetworkException e) {
            e.printStackTrace();
        } catch (AnalysisException e) {
            e.printStackTrace();
        }

        // You must use a builder to construct a new Protobuffer object

        // Use responseObserver to send a single response back

        // When you are done, you must call onCompleted.
        responseObserver.onCompleted();
    }
}