package com.iggroup.webapi.samples.client.grpc;

import io.grpc.*;

public class GrpcServer
{
    public static void main( String[] args ) throws Exception
    {
        // Create a new server to listen on port
        Server server = ServerBuilder.forPort(50055)
                .addService(new IGAPIServiceImpl())
                .build();

        // Start the server
        server.start();

        // Server threads are running in the background.
        System.out.println("Server started and listening on port"+50055);
        // Don't exit the main thread. Wait until server is terminated.
        server.awaitTermination();
    }
}