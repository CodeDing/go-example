// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: example.proto

package example;

public final class Example {
  private Example() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ErrorStatusOrBuilder extends
      // @@protoc_insertion_point(interface_extends:example.ErrorStatus)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string message = 1;</code>
     */
    java.lang.String getMessage();
    /**
     * <code>string message = 1;</code>
     */
    com.google.protobuf.ByteString
        getMessageBytes();

    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    java.util.List<com.google.protobuf.Any> 
        getDetailsList();
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    com.google.protobuf.Any getDetails(int index);
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    int getDetailsCount();
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    java.util.List<? extends com.google.protobuf.AnyOrBuilder> 
        getDetailsOrBuilderList();
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    com.google.protobuf.AnyOrBuilder getDetailsOrBuilder(
        int index);
  }
  /**
   * Protobuf type {@code example.ErrorStatus}
   */
  public  static final class ErrorStatus extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:example.ErrorStatus)
      ErrorStatusOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use ErrorStatus.newBuilder() to construct.
    private ErrorStatus(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private ErrorStatus() {
      message_ = "";
      details_ = java.util.Collections.emptyList();
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private ErrorStatus(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      int mutable_bitField0_ = 0;
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

              message_ = s;
              break;
            }
            case 18: {
              if (!((mutable_bitField0_ & 0x00000002) == 0x00000002)) {
                details_ = new java.util.ArrayList<com.google.protobuf.Any>();
                mutable_bitField0_ |= 0x00000002;
              }
              details_.add(
                  input.readMessage(com.google.protobuf.Any.parser(), extensionRegistry));
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
        if (((mutable_bitField0_ & 0x00000002) == 0x00000002)) {
          details_ = java.util.Collections.unmodifiableList(details_);
        }
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return example.Example.internal_static_example_ErrorStatus_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return example.Example.internal_static_example_ErrorStatus_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              example.Example.ErrorStatus.class, example.Example.ErrorStatus.Builder.class);
    }

    private int bitField0_;
    public static final int MESSAGE_FIELD_NUMBER = 1;
    private volatile java.lang.Object message_;
    /**
     * <code>string message = 1;</code>
     */
    public java.lang.String getMessage() {
      java.lang.Object ref = message_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        message_ = s;
        return s;
      }
    }
    /**
     * <code>string message = 1;</code>
     */
    public com.google.protobuf.ByteString
        getMessageBytes() {
      java.lang.Object ref = message_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        message_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int DETAILS_FIELD_NUMBER = 2;
    private java.util.List<com.google.protobuf.Any> details_;
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    public java.util.List<com.google.protobuf.Any> getDetailsList() {
      return details_;
    }
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    public java.util.List<? extends com.google.protobuf.AnyOrBuilder> 
        getDetailsOrBuilderList() {
      return details_;
    }
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    public int getDetailsCount() {
      return details_.size();
    }
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    public com.google.protobuf.Any getDetails(int index) {
      return details_.get(index);
    }
    /**
     * <code>repeated .google.protobuf.Any details = 2;</code>
     */
    public com.google.protobuf.AnyOrBuilder getDetailsOrBuilder(
        int index) {
      return details_.get(index);
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
      if (!getMessageBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, message_);
      }
      for (int i = 0; i < details_.size(); i++) {
        output.writeMessage(2, details_.get(i));
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!getMessageBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, message_);
      }
      for (int i = 0; i < details_.size(); i++) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(2, details_.get(i));
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
      if (!(obj instanceof example.Example.ErrorStatus)) {
        return super.equals(obj);
      }
      example.Example.ErrorStatus other = (example.Example.ErrorStatus) obj;

      boolean result = true;
      result = result && getMessage()
          .equals(other.getMessage());
      result = result && getDetailsList()
          .equals(other.getDetailsList());
      result = result && unknownFields.equals(other.unknownFields);
      return result;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + MESSAGE_FIELD_NUMBER;
      hash = (53 * hash) + getMessage().hashCode();
      if (getDetailsCount() > 0) {
        hash = (37 * hash) + DETAILS_FIELD_NUMBER;
        hash = (53 * hash) + getDetailsList().hashCode();
      }
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static example.Example.ErrorStatus parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static example.Example.ErrorStatus parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static example.Example.ErrorStatus parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static example.Example.ErrorStatus parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static example.Example.ErrorStatus parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static example.Example.ErrorStatus parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static example.Example.ErrorStatus parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static example.Example.ErrorStatus parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static example.Example.ErrorStatus parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static example.Example.ErrorStatus parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static example.Example.ErrorStatus parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static example.Example.ErrorStatus parseFrom(
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
    public static Builder newBuilder(example.Example.ErrorStatus prototype) {
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
     * Protobuf type {@code example.ErrorStatus}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:example.ErrorStatus)
        example.Example.ErrorStatusOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return example.Example.internal_static_example_ErrorStatus_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return example.Example.internal_static_example_ErrorStatus_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                example.Example.ErrorStatus.class, example.Example.ErrorStatus.Builder.class);
      }

      // Construct using example.Example.ErrorStatus.newBuilder()
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
          getDetailsFieldBuilder();
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        message_ = "";

        if (detailsBuilder_ == null) {
          details_ = java.util.Collections.emptyList();
          bitField0_ = (bitField0_ & ~0x00000002);
        } else {
          detailsBuilder_.clear();
        }
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return example.Example.internal_static_example_ErrorStatus_descriptor;
      }

      @java.lang.Override
      public example.Example.ErrorStatus getDefaultInstanceForType() {
        return example.Example.ErrorStatus.getDefaultInstance();
      }

      @java.lang.Override
      public example.Example.ErrorStatus build() {
        example.Example.ErrorStatus result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public example.Example.ErrorStatus buildPartial() {
        example.Example.ErrorStatus result = new example.Example.ErrorStatus(this);
        int from_bitField0_ = bitField0_;
        int to_bitField0_ = 0;
        result.message_ = message_;
        if (detailsBuilder_ == null) {
          if (((bitField0_ & 0x00000002) == 0x00000002)) {
            details_ = java.util.Collections.unmodifiableList(details_);
            bitField0_ = (bitField0_ & ~0x00000002);
          }
          result.details_ = details_;
        } else {
          result.details_ = detailsBuilder_.build();
        }
        result.bitField0_ = to_bitField0_;
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
        if (other instanceof example.Example.ErrorStatus) {
          return mergeFrom((example.Example.ErrorStatus)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(example.Example.ErrorStatus other) {
        if (other == example.Example.ErrorStatus.getDefaultInstance()) return this;
        if (!other.getMessage().isEmpty()) {
          message_ = other.message_;
          onChanged();
        }
        if (detailsBuilder_ == null) {
          if (!other.details_.isEmpty()) {
            if (details_.isEmpty()) {
              details_ = other.details_;
              bitField0_ = (bitField0_ & ~0x00000002);
            } else {
              ensureDetailsIsMutable();
              details_.addAll(other.details_);
            }
            onChanged();
          }
        } else {
          if (!other.details_.isEmpty()) {
            if (detailsBuilder_.isEmpty()) {
              detailsBuilder_.dispose();
              detailsBuilder_ = null;
              details_ = other.details_;
              bitField0_ = (bitField0_ & ~0x00000002);
              detailsBuilder_ = 
                com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                   getDetailsFieldBuilder() : null;
            } else {
              detailsBuilder_.addAllMessages(other.details_);
            }
          }
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
        example.Example.ErrorStatus parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (example.Example.ErrorStatus) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }
      private int bitField0_;

      private java.lang.Object message_ = "";
      /**
       * <code>string message = 1;</code>
       */
      public java.lang.String getMessage() {
        java.lang.Object ref = message_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          message_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string message = 1;</code>
       */
      public com.google.protobuf.ByteString
          getMessageBytes() {
        java.lang.Object ref = message_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          message_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string message = 1;</code>
       */
      public Builder setMessage(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        message_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string message = 1;</code>
       */
      public Builder clearMessage() {
        
        message_ = getDefaultInstance().getMessage();
        onChanged();
        return this;
      }
      /**
       * <code>string message = 1;</code>
       */
      public Builder setMessageBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        message_ = value;
        onChanged();
        return this;
      }

      private java.util.List<com.google.protobuf.Any> details_ =
        java.util.Collections.emptyList();
      private void ensureDetailsIsMutable() {
        if (!((bitField0_ & 0x00000002) == 0x00000002)) {
          details_ = new java.util.ArrayList<com.google.protobuf.Any>(details_);
          bitField0_ |= 0x00000002;
         }
      }

      private com.google.protobuf.RepeatedFieldBuilderV3<
          com.google.protobuf.Any, com.google.protobuf.Any.Builder, com.google.protobuf.AnyOrBuilder> detailsBuilder_;

      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public java.util.List<com.google.protobuf.Any> getDetailsList() {
        if (detailsBuilder_ == null) {
          return java.util.Collections.unmodifiableList(details_);
        } else {
          return detailsBuilder_.getMessageList();
        }
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public int getDetailsCount() {
        if (detailsBuilder_ == null) {
          return details_.size();
        } else {
          return detailsBuilder_.getCount();
        }
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public com.google.protobuf.Any getDetails(int index) {
        if (detailsBuilder_ == null) {
          return details_.get(index);
        } else {
          return detailsBuilder_.getMessage(index);
        }
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder setDetails(
          int index, com.google.protobuf.Any value) {
        if (detailsBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          ensureDetailsIsMutable();
          details_.set(index, value);
          onChanged();
        } else {
          detailsBuilder_.setMessage(index, value);
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder setDetails(
          int index, com.google.protobuf.Any.Builder builderForValue) {
        if (detailsBuilder_ == null) {
          ensureDetailsIsMutable();
          details_.set(index, builderForValue.build());
          onChanged();
        } else {
          detailsBuilder_.setMessage(index, builderForValue.build());
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder addDetails(com.google.protobuf.Any value) {
        if (detailsBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          ensureDetailsIsMutable();
          details_.add(value);
          onChanged();
        } else {
          detailsBuilder_.addMessage(value);
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder addDetails(
          int index, com.google.protobuf.Any value) {
        if (detailsBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          ensureDetailsIsMutable();
          details_.add(index, value);
          onChanged();
        } else {
          detailsBuilder_.addMessage(index, value);
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder addDetails(
          com.google.protobuf.Any.Builder builderForValue) {
        if (detailsBuilder_ == null) {
          ensureDetailsIsMutable();
          details_.add(builderForValue.build());
          onChanged();
        } else {
          detailsBuilder_.addMessage(builderForValue.build());
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder addDetails(
          int index, com.google.protobuf.Any.Builder builderForValue) {
        if (detailsBuilder_ == null) {
          ensureDetailsIsMutable();
          details_.add(index, builderForValue.build());
          onChanged();
        } else {
          detailsBuilder_.addMessage(index, builderForValue.build());
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder addAllDetails(
          java.lang.Iterable<? extends com.google.protobuf.Any> values) {
        if (detailsBuilder_ == null) {
          ensureDetailsIsMutable();
          com.google.protobuf.AbstractMessageLite.Builder.addAll(
              values, details_);
          onChanged();
        } else {
          detailsBuilder_.addAllMessages(values);
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder clearDetails() {
        if (detailsBuilder_ == null) {
          details_ = java.util.Collections.emptyList();
          bitField0_ = (bitField0_ & ~0x00000002);
          onChanged();
        } else {
          detailsBuilder_.clear();
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public Builder removeDetails(int index) {
        if (detailsBuilder_ == null) {
          ensureDetailsIsMutable();
          details_.remove(index);
          onChanged();
        } else {
          detailsBuilder_.remove(index);
        }
        return this;
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public com.google.protobuf.Any.Builder getDetailsBuilder(
          int index) {
        return getDetailsFieldBuilder().getBuilder(index);
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public com.google.protobuf.AnyOrBuilder getDetailsOrBuilder(
          int index) {
        if (detailsBuilder_ == null) {
          return details_.get(index);  } else {
          return detailsBuilder_.getMessageOrBuilder(index);
        }
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public java.util.List<? extends com.google.protobuf.AnyOrBuilder> 
           getDetailsOrBuilderList() {
        if (detailsBuilder_ != null) {
          return detailsBuilder_.getMessageOrBuilderList();
        } else {
          return java.util.Collections.unmodifiableList(details_);
        }
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public com.google.protobuf.Any.Builder addDetailsBuilder() {
        return getDetailsFieldBuilder().addBuilder(
            com.google.protobuf.Any.getDefaultInstance());
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public com.google.protobuf.Any.Builder addDetailsBuilder(
          int index) {
        return getDetailsFieldBuilder().addBuilder(
            index, com.google.protobuf.Any.getDefaultInstance());
      }
      /**
       * <code>repeated .google.protobuf.Any details = 2;</code>
       */
      public java.util.List<com.google.protobuf.Any.Builder> 
           getDetailsBuilderList() {
        return getDetailsFieldBuilder().getBuilderList();
      }
      private com.google.protobuf.RepeatedFieldBuilderV3<
          com.google.protobuf.Any, com.google.protobuf.Any.Builder, com.google.protobuf.AnyOrBuilder> 
          getDetailsFieldBuilder() {
        if (detailsBuilder_ == null) {
          detailsBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
              com.google.protobuf.Any, com.google.protobuf.Any.Builder, com.google.protobuf.AnyOrBuilder>(
                  details_,
                  ((bitField0_ & 0x00000002) == 0x00000002),
                  getParentForChildren(),
                  isClean());
          details_ = null;
        }
        return detailsBuilder_;
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


      // @@protoc_insertion_point(builder_scope:example.ErrorStatus)
    }

    // @@protoc_insertion_point(class_scope:example.ErrorStatus)
    private static final example.Example.ErrorStatus DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new example.Example.ErrorStatus();
    }

    public static example.Example.ErrorStatus getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<ErrorStatus>
        PARSER = new com.google.protobuf.AbstractParser<ErrorStatus>() {
      @java.lang.Override
      public ErrorStatus parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new ErrorStatus(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<ErrorStatus> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<ErrorStatus> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public example.Example.ErrorStatus getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_example_ErrorStatus_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_example_ErrorStatus_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\rexample.proto\022\007example\032\031google/protobu" +
      "f/any.proto\"E\n\013ErrorStatus\022\017\n\007message\030\001 " +
      "\001(\t\022%\n\007details\030\002 \003(\0132\024.google.protobuf.A" +
      "nyb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.AnyProto.getDescriptor(),
        }, assigner);
    internal_static_example_ErrorStatus_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_example_ErrorStatus_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_example_ErrorStatus_descriptor,
        new java.lang.String[] { "Message", "Details", });
    com.google.protobuf.AnyProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
