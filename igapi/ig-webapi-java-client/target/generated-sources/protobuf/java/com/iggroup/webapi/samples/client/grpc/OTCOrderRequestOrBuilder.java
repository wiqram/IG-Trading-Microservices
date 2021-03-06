// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: igapi.proto

package com.iggroup.webapi.samples.client.grpc;

public interface OTCOrderRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:demo_igapi.OTCOrderRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string epic = 1;</code>
   * @return The epic.
   */
  java.lang.String getEpic();
  /**
   * <code>string epic = 1;</code>
   * @return The bytes for epic.
   */
  com.google.protobuf.ByteString
      getEpicBytes();

  /**
   * <code>string level = 2;</code>
   * @return The level.
   */
  java.lang.String getLevel();
  /**
   * <code>string level = 2;</code>
   * @return The bytes for level.
   */
  com.google.protobuf.ByteString
      getLevelBytes();

  /**
   * <code>bool forceOpen = 3;</code>
   * @return The forceOpen.
   */
  boolean getForceOpen();

  /**
   * <code>string orderType = 4;</code>
   * @return The orderType.
   */
  java.lang.String getOrderType();
  /**
   * <code>string orderType = 4;</code>
   * @return The bytes for orderType.
   */
  com.google.protobuf.ByteString
      getOrderTypeBytes();

  /**
   * <code>string currencyCode = 5;</code>
   * @return The currencyCode.
   */
  java.lang.String getCurrencyCode();
  /**
   * <code>string currencyCode = 5;</code>
   * @return The bytes for currencyCode.
   */
  com.google.protobuf.ByteString
      getCurrencyCodeBytes();

  /**
   * <pre>
   * "BUY" or "SELL"
   * </pre>
   *
   * <code>string direction = 6;</code>
   * @return The direction.
   */
  java.lang.String getDirection();
  /**
   * <pre>
   * "BUY" or "SELL"
   * </pre>
   *
   * <code>string direction = 6;</code>
   * @return The bytes for direction.
   */
  com.google.protobuf.ByteString
      getDirectionBytes();

  /**
   * <code>string expiry = 7;</code>
   * @return The expiry.
   */
  java.lang.String getExpiry();
  /**
   * <code>string expiry = 7;</code>
   * @return The bytes for expiry.
   */
  com.google.protobuf.ByteString
      getExpiryBytes();

  /**
   * <code>float size = 8;</code>
   * @return The size.
   */
  float getSize();

  /**
   * <pre>
   * Pips
   * </pre>
   *
   * <code>string stopDistance = 9;</code>
   * @return The stopDistance.
   */
  java.lang.String getStopDistance();
  /**
   * <pre>
   * Pips
   * </pre>
   *
   * <code>string stopDistance = 9;</code>
   * @return The bytes for stopDistance.
   */
  com.google.protobuf.ByteString
      getStopDistanceBytes();

  /**
   * <pre>
   * Pips
   * </pre>
   *
   * <code>string stopLevel = 10;</code>
   * @return The stopLevel.
   */
  java.lang.String getStopLevel();
  /**
   * <pre>
   * Pips
   * </pre>
   *
   * <code>string stopLevel = 10;</code>
   * @return The bytes for stopLevel.
   */
  com.google.protobuf.ByteString
      getStopLevelBytes();

  /**
   * <pre>
   * Pips
   * </pre>
   *
   * <code>string limitDistance = 11;</code>
   * @return The limitDistance.
   */
  java.lang.String getLimitDistance();
  /**
   * <pre>
   * Pips
   * </pre>
   *
   * <code>string limitDistance = 11;</code>
   * @return The bytes for limitDistance.
   */
  com.google.protobuf.ByteString
      getLimitDistanceBytes();

  /**
   * <pre>
   * Pips
   * </pre>
   *
   * <code>string limitLevel = 12;</code>
   * @return The limitLevel.
   */
  java.lang.String getLimitLevel();
  /**
   * <pre>
   * Pips
   * </pre>
   *
   * <code>string limitLevel = 12;</code>
   * @return The bytes for limitLevel.
   */
  com.google.protobuf.ByteString
      getLimitLevelBytes();

  /**
   * <code>string quoteId = 13;</code>
   * @return The quoteId.
   */
  java.lang.String getQuoteId();
  /**
   * <code>string quoteId = 13;</code>
   * @return The bytes for quoteId.
   */
  com.google.protobuf.ByteString
      getQuoteIdBytes();

  /**
   * <pre>
   * "EXECUTE_AND_ELIMINATE" or "FILL_OR_KILL"
   * </pre>
   *
   * <code>string timeInForce = 14;</code>
   * @return The timeInForce.
   */
  java.lang.String getTimeInForce();
  /**
   * <pre>
   * "EXECUTE_AND_ELIMINATE" or "FILL_OR_KILL"
   * </pre>
   *
   * <code>string timeInForce = 14;</code>
   * @return The bytes for timeInForce.
   */
  com.google.protobuf.ByteString
      getTimeInForceBytes();

  /**
   * <code>bool trailingStop = 15;</code>
   * @return The trailingStop.
   */
  boolean getTrailingStop();

  /**
   * <code>string trailingStopIncrement = 16;</code>
   * @return The trailingStopIncrement.
   */
  java.lang.String getTrailingStopIncrement();
  /**
   * <code>string trailingStopIncrement = 16;</code>
   * @return The bytes for trailingStopIncrement.
   */
  com.google.protobuf.ByteString
      getTrailingStopIncrementBytes();

  /**
   * <code>bool guaranteedStop = 17;</code>
   * @return The guaranteedStop.
   */
  boolean getGuaranteedStop();

  /**
   * <code>string dealReference = 19;</code>
   * @return The dealReference.
   */
  java.lang.String getDealReference();
  /**
   * <code>string dealReference = 19;</code>
   * @return The bytes for dealReference.
   */
  com.google.protobuf.ByteString
      getDealReferenceBytes();
}
