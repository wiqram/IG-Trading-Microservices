// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: igapi.proto

package com.iggroup.webapi.samples.client.grpc;

public interface PositionsOrBuilder extends
    // @@protoc_insertion_point(interface_extends:demo_igapi.Positions)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.demo_igapi.MarketData marketData = 1;</code>
   * @return Whether the marketData field is set.
   */
  boolean hasMarketData();
  /**
   * <code>.demo_igapi.MarketData marketData = 1;</code>
   * @return The marketData.
   */
  com.iggroup.webapi.samples.client.grpc.MarketData getMarketData();
  /**
   * <code>.demo_igapi.MarketData marketData = 1;</code>
   */
  com.iggroup.webapi.samples.client.grpc.MarketDataOrBuilder getMarketDataOrBuilder();

  /**
   * <code>.demo_igapi.Position position = 2;</code>
   * @return Whether the position field is set.
   */
  boolean hasPosition();
  /**
   * <code>.demo_igapi.Position position = 2;</code>
   * @return The position.
   */
  com.iggroup.webapi.samples.client.grpc.Position getPosition();
  /**
   * <code>.demo_igapi.Position position = 2;</code>
   */
  com.iggroup.webapi.samples.client.grpc.PositionOrBuilder getPositionOrBuilder();
}
