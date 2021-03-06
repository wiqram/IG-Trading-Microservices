// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: igapi.proto

package com.iggroup.webapi.samples.client.grpc;

/**
 * <pre>
 * ClientSentimentRequest - Response for client sentiment
 * </pre>
 *
 * Protobuf type {@code demo_igapi.ClientSentimentRequest}
 */
public final class ClientSentimentRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:demo_igapi.ClientSentimentRequest)
    ClientSentimentRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use ClientSentimentRequest.newBuilder() to construct.
  private ClientSentimentRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private ClientSentimentRequest() {
    marketId_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new ClientSentimentRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private ClientSentimentRequest(
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
    return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest.class, com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest.Builder.class);
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
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest)) {
      return super.equals(obj);
    }
    com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest other = (com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest) obj;

    if (!getMarketId()
        .equals(other.getMarketId())) return false;
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
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parseFrom(
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
  public static Builder newBuilder(com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest prototype) {
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
   * ClientSentimentRequest - Response for client sentiment
   * </pre>
   *
   * Protobuf type {@code demo_igapi.ClientSentimentRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:demo_igapi.ClientSentimentRequest)
      com.iggroup.webapi.samples.client.grpc.ClientSentimentRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest.class, com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest.Builder.class);
    }

    // Construct using com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest.newBuilder()
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

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentRequest_descriptor;
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest getDefaultInstanceForType() {
      return com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest.getDefaultInstance();
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest build() {
      com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest buildPartial() {
      com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest result = new com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest(this);
      result.marketId_ = marketId_;
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
      if (other instanceof com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest) {
        return mergeFrom((com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest other) {
      if (other == com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest.getDefaultInstance()) return this;
      if (!other.getMarketId().isEmpty()) {
        marketId_ = other.marketId_;
        onChanged();
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
      com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest) e.getUnfinishedMessage();
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


    // @@protoc_insertion_point(builder_scope:demo_igapi.ClientSentimentRequest)
  }

  // @@protoc_insertion_point(class_scope:demo_igapi.ClientSentimentRequest)
  private static final com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest();
  }

  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<ClientSentimentRequest>
      PARSER = new com.google.protobuf.AbstractParser<ClientSentimentRequest>() {
    @java.lang.Override
    public ClientSentimentRequest parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new ClientSentimentRequest(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<ClientSentimentRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<ClientSentimentRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.iggroup.webapi.samples.client.grpc.ClientSentimentRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

