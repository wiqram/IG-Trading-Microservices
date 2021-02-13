// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: igapi.proto

package com.iggroup.webapi.samples.client.grpc;

/**
 * <pre>
 * ClientSentimentResponse - Response for client sentiment
 * </pre>
 *
 * Protobuf type {@code demo_igapi.ClientSentimentResponse}
 */
public final class ClientSentimentResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:demo_igapi.ClientSentimentResponse)
    ClientSentimentResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use ClientSentimentResponse.newBuilder() to construct.
  private ClientSentimentResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private ClientSentimentResponse() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new ClientSentimentResponse();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private ClientSentimentResponse(
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
          case 13: {

            longPositionPercentage_ = input.readFloat();
            break;
          }
          case 21: {

            shortPositionPercentage_ = input.readFloat();
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
    return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse.class, com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse.Builder.class);
  }

  public static final int LONGPOSITIONPERCENTAGE_FIELD_NUMBER = 1;
  private float longPositionPercentage_;
  /**
   * <code>float longPositionPercentage = 1;</code>
   * @return The longPositionPercentage.
   */
  @java.lang.Override
  public float getLongPositionPercentage() {
    return longPositionPercentage_;
  }

  public static final int SHORTPOSITIONPERCENTAGE_FIELD_NUMBER = 2;
  private float shortPositionPercentage_;
  /**
   * <code>float shortPositionPercentage = 2;</code>
   * @return The shortPositionPercentage.
   */
  @java.lang.Override
  public float getShortPositionPercentage() {
    return shortPositionPercentage_;
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
    if (longPositionPercentage_ != 0F) {
      output.writeFloat(1, longPositionPercentage_);
    }
    if (shortPositionPercentage_ != 0F) {
      output.writeFloat(2, shortPositionPercentage_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (longPositionPercentage_ != 0F) {
      size += com.google.protobuf.CodedOutputStream
        .computeFloatSize(1, longPositionPercentage_);
    }
    if (shortPositionPercentage_ != 0F) {
      size += com.google.protobuf.CodedOutputStream
        .computeFloatSize(2, shortPositionPercentage_);
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
    if (!(obj instanceof com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse)) {
      return super.equals(obj);
    }
    com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse other = (com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse) obj;

    if (java.lang.Float.floatToIntBits(getLongPositionPercentage())
        != java.lang.Float.floatToIntBits(
            other.getLongPositionPercentage())) return false;
    if (java.lang.Float.floatToIntBits(getShortPositionPercentage())
        != java.lang.Float.floatToIntBits(
            other.getShortPositionPercentage())) return false;
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
    hash = (37 * hash) + LONGPOSITIONPERCENTAGE_FIELD_NUMBER;
    hash = (53 * hash) + java.lang.Float.floatToIntBits(
        getLongPositionPercentage());
    hash = (37 * hash) + SHORTPOSITIONPERCENTAGE_FIELD_NUMBER;
    hash = (53 * hash) + java.lang.Float.floatToIntBits(
        getShortPositionPercentage());
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parseFrom(
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
  public static Builder newBuilder(com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse prototype) {
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
   * ClientSentimentResponse - Response for client sentiment
   * </pre>
   *
   * Protobuf type {@code demo_igapi.ClientSentimentResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:demo_igapi.ClientSentimentResponse)
      com.iggroup.webapi.samples.client.grpc.ClientSentimentResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse.class, com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse.Builder.class);
    }

    // Construct using com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse.newBuilder()
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
      longPositionPercentage_ = 0F;

      shortPositionPercentage_ = 0F;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.iggroup.webapi.samples.client.grpc.Igapi.internal_static_demo_igapi_ClientSentimentResponse_descriptor;
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse getDefaultInstanceForType() {
      return com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse.getDefaultInstance();
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse build() {
      com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse buildPartial() {
      com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse result = new com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse(this);
      result.longPositionPercentage_ = longPositionPercentage_;
      result.shortPositionPercentage_ = shortPositionPercentage_;
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
      if (other instanceof com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse) {
        return mergeFrom((com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse other) {
      if (other == com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse.getDefaultInstance()) return this;
      if (other.getLongPositionPercentage() != 0F) {
        setLongPositionPercentage(other.getLongPositionPercentage());
      }
      if (other.getShortPositionPercentage() != 0F) {
        setShortPositionPercentage(other.getShortPositionPercentage());
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
      com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private float longPositionPercentage_ ;
    /**
     * <code>float longPositionPercentage = 1;</code>
     * @return The longPositionPercentage.
     */
    @java.lang.Override
    public float getLongPositionPercentage() {
      return longPositionPercentage_;
    }
    /**
     * <code>float longPositionPercentage = 1;</code>
     * @param value The longPositionPercentage to set.
     * @return This builder for chaining.
     */
    public Builder setLongPositionPercentage(float value) {
      
      longPositionPercentage_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>float longPositionPercentage = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearLongPositionPercentage() {
      
      longPositionPercentage_ = 0F;
      onChanged();
      return this;
    }

    private float shortPositionPercentage_ ;
    /**
     * <code>float shortPositionPercentage = 2;</code>
     * @return The shortPositionPercentage.
     */
    @java.lang.Override
    public float getShortPositionPercentage() {
      return shortPositionPercentage_;
    }
    /**
     * <code>float shortPositionPercentage = 2;</code>
     * @param value The shortPositionPercentage to set.
     * @return This builder for chaining.
     */
    public Builder setShortPositionPercentage(float value) {
      
      shortPositionPercentage_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>float shortPositionPercentage = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearShortPositionPercentage() {
      
      shortPositionPercentage_ = 0F;
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


    // @@protoc_insertion_point(builder_scope:demo_igapi.ClientSentimentResponse)
  }

  // @@protoc_insertion_point(class_scope:demo_igapi.ClientSentimentResponse)
  private static final com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse();
  }

  public static com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<ClientSentimentResponse>
      PARSER = new com.google.protobuf.AbstractParser<ClientSentimentResponse>() {
    @java.lang.Override
    public ClientSentimentResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new ClientSentimentResponse(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<ClientSentimentResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<ClientSentimentResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.iggroup.webapi.samples.client.grpc.ClientSentimentResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

