// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: igapi.proto

package com.iggroup.webapi.samples.client.grpc;

/**
 * <pre>
 * LightStreamerSubResponse - Response for client sentiment
 * </pre>
 *
 * Protobuf type {@code demo_igapi.LightStreamerSubResponse}
 */
public final class LightStreamerSubResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:demo_igapi.LightStreamerSubResponse)
    LightStreamerSubResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use LightStreamerSubResponse.newBuilder() to construct.
  private LightStreamerSubResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private LightStreamerSubResponse() {
    marketId_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new LightStreamerSubResponse();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private LightStreamerSubResponse(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();

            marketId_ = s;
            break;
          }
          case 16: {

            time_ = input.readInt32();
            break;
          }
          case 29: {

            bid_ = input.readFloat();
            break;
          }
          case 37: {

            ask_ = input.readFloat();
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_LightStreamerSubResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_LightStreamerSubResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse.class, com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse.Builder.class);
  }

  public static final int MARKETID_FIELD_NUMBER = 1;
  private volatile java.lang.Object marketId_;
  /**
   * <code>string marketId = 1;</code>
   * @return The marketId.
   */
  @java.lang.Override
  public java.lang.String getMarketId() {
    java.lang.Object ref = marketId_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      marketId_ = s;
      return s;
    }
  }
  /**
   * <code>string marketId = 1;</code>
   * @return The bytes for marketId.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getMarketIdBytes() {
    java.lang.Object ref = marketId_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      marketId_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int TIME_FIELD_NUMBER = 2;
  private int time_;
  /**
   * <code>int32 time = 2;</code>
   * @return The time.
   */
  @java.lang.Override
  public int getTime() {
    return time_;
  }

  public static final int BID_FIELD_NUMBER = 3;
  private float bid_;
  /**
   * <code>float bid = 3;</code>
   * @return The bid.
   */
  @java.lang.Override
  public float getBid() {
    return bid_;
  }

  public static final int ASK_FIELD_NUMBER = 4;
  private float ask_;
  /**
   * <code>float ask = 4;</code>
   * @return The ask.
   */
  @java.lang.Override
  public float getAsk() {
    return ask_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!getMarketIdBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, marketId_);
    }
    if (time_ != 0) {
      output.writeInt32(2, time_);
    }
    if (bid_ != 0F) {
      output.writeFloat(3, bid_);
    }
    if (ask_ != 0F) {
      output.writeFloat(4, ask_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getMarketIdBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, marketId_);
    }
    if (time_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, time_);
    }
    if (bid_ != 0F) {
      size += com.google.protobuf.CodedOutputStream
        .computeFloatSize(3, bid_);
    }
    if (ask_ != 0F) {
      size += com.google.protobuf.CodedOutputStream
        .computeFloatSize(4, ask_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse)) {
      return super.equals(obj);
    }
    com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse other = (com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse) obj;

    if (!getMarketId()
        .equals(other.getMarketId())) return false;
    if (getTime()
        != other.getTime()) return false;
    if (java.lang.Float.floatToIntBits(getBid())
        != java.lang.Float.floatToIntBits(
            other.getBid())) return false;
    if (java.lang.Float.floatToIntBits(getAsk())
        != java.lang.Float.floatToIntBits(
            other.getAsk())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + MARKETID_FIELD_NUMBER;
    hash = (53 * hash) + getMarketId().hashCode();
    hash = (37 * hash) + TIME_FIELD_NUMBER;
    hash = (53 * hash) + getTime();
    hash = (37 * hash) + BID_FIELD_NUMBER;
    hash = (53 * hash) + java.lang.Float.floatToIntBits(
        getBid());
    hash = (37 * hash) + ASK_FIELD_NUMBER;
    hash = (53 * hash) + java.lang.Float.floatToIntBits(
        getAsk());
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * <pre>
   * LightStreamerSubResponse - Response for client sentiment
   * </pre>
   *
   * Protobuf type {@code demo_igapi.LightStreamerSubResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:demo_igapi.LightStreamerSubResponse)
      com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_LightStreamerSubResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_LightStreamerSubResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse.class, com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse.Builder.class);
    }

    // Construct using com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      marketId_ = "";

      time_ = 0;

      bid_ = 0F;

      ask_ = 0F;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_LightStreamerSubResponse_descriptor;
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse getDefaultInstanceForType() {
      return com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse.getDefaultInstance();
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse build() {
      com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse buildPartial() {
      com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse result = new com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse(this);
      result.marketId_ = marketId_;
      result.time_ = time_;
      result.bid_ = bid_;
      result.ask_ = ask_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse) {
        return mergeFrom((com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse other) {
      if (other == com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse.getDefaultInstance()) return this;
      if (!other.getMarketId().isEmpty()) {
        marketId_ = other.marketId_;
        onChanged();
      }
      if (other.getTime() != 0) {
        setTime(other.getTime());
      }
      if (other.getBid() != 0F) {
        setBid(other.getBid());
      }
      if (other.getAsk() != 0F) {
        setAsk(other.getAsk());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private java.lang.Object marketId_ = "";
    /**
     * <code>string marketId = 1;</code>
     * @return The marketId.
     */
    public java.lang.String getMarketId() {
      java.lang.Object ref = marketId_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        marketId_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string marketId = 1;</code>
     * @return The bytes for marketId.
     */
    public com.google.protobuf.ByteString
        getMarketIdBytes() {
      java.lang.Object ref = marketId_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        marketId_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string marketId = 1;</code>
     * @param value The marketId to set.
     * @return This builder for chaining.
     */
    public Builder setMarketId(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      marketId_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string marketId = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearMarketId() {
      
      marketId_ = getDefaultInstance().getMarketId();
      onChanged();
      return this;
    }
    /**
     * <code>string marketId = 1;</code>
     * @param value The bytes for marketId to set.
     * @return This builder for chaining.
     */
    public Builder setMarketIdBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      marketId_ = value;
      onChanged();
      return this;
    }

    private int time_ ;
    /**
     * <code>int32 time = 2;</code>
     * @return The time.
     */
    @java.lang.Override
    public int getTime() {
      return time_;
    }
    /**
     * <code>int32 time = 2;</code>
     * @param value The time to set.
     * @return This builder for chaining.
     */
    public Builder setTime(int value) {
      
      time_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 time = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearTime() {
      
      time_ = 0;
      onChanged();
      return this;
    }

    private float bid_ ;
    /**
     * <code>float bid = 3;</code>
     * @return The bid.
     */
    @java.lang.Override
    public float getBid() {
      return bid_;
    }
    /**
     * <code>float bid = 3;</code>
     * @param value The bid to set.
     * @return This builder for chaining.
     */
    public Builder setBid(float value) {
      
      bid_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>float bid = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearBid() {
      
      bid_ = 0F;
      onChanged();
      return this;
    }

    private float ask_ ;
    /**
     * <code>float ask = 4;</code>
     * @return The ask.
     */
    @java.lang.Override
    public float getAsk() {
      return ask_;
    }
    /**
     * <code>float ask = 4;</code>
     * @param value The ask to set.
     * @return This builder for chaining.
     */
    public Builder setAsk(float value) {
      
      ask_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>float ask = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearAsk() {
      
      ask_ = 0F;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:demo_igapi.LightStreamerSubResponse)
  }

  // @@protoc_insertion_point(class_scope:demo_igapi.LightStreamerSubResponse)
  private static final com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse();
  }

  public static com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<LightStreamerSubResponse>
      PARSER = new com.google.protobuf.AbstractParser<LightStreamerSubResponse>() {
    @java.lang.Override
    public LightStreamerSubResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new LightStreamerSubResponse(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<LightStreamerSubResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<LightStreamerSubResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.iggroup.webapi.samples.client.grpc.LightStreamerSubResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
