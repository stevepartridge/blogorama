/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_api_annotations_pb = require('./google/api/annotations_pb.js');
goog.object.extend(proto, google_api_annotations_pb);
var protoc$gen$swagger_options_annotations_pb = require('./protoc-gen-swagger/options/annotations_pb.js');
goog.object.extend(proto, protoc$gen$swagger_options_annotations_pb);
goog.exportSymbol('proto.blog.Post', null, global);
goog.exportSymbol('proto.blog.PostRequest', null, global);
goog.exportSymbol('proto.blog.PostResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.blog.Post = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.blog.Post, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.blog.Post.displayName = 'proto.blog.Post';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.blog.PostRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.blog.PostRequest.repeatedFields_, null);
};
goog.inherits(proto.blog.PostRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.blog.PostRequest.displayName = 'proto.blog.PostRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.blog.PostResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.blog.PostResponse.repeatedFields_, null);
};
goog.inherits(proto.blog.PostResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.blog.PostResponse.displayName = 'proto.blog.PostResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.blog.Post.prototype.toObject = function(opt_includeInstance) {
  return proto.blog.Post.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.blog.Post} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blog.Post.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    title: jspb.Message.getFieldWithDefault(msg, 3, ""),
    content: jspb.Message.getFieldWithDefault(msg, 4, ""),
    pb_private: jspb.Message.getBooleanFieldWithDefault(msg, 5, false),
    createdById: jspb.Message.getFieldWithDefault(msg, 20, 0),
    createdAt: jspb.Message.getFieldWithDefault(msg, 21, ""),
    updatedById: jspb.Message.getFieldWithDefault(msg, 22, 0),
    updatedAt: jspb.Message.getFieldWithDefault(msg, 23, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.blog.Post}
 */
proto.blog.Post.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.blog.Post;
  return proto.blog.Post.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.blog.Post} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.blog.Post}
 */
proto.blog.Post.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setTitle(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setContent(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setPrivate(value);
      break;
    case 20:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCreatedById(value);
      break;
    case 21:
      var value = /** @type {string} */ (reader.readString());
      msg.setCreatedAt(value);
      break;
    case 22:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setUpdatedById(value);
      break;
    case 23:
      var value = /** @type {string} */ (reader.readString());
      msg.setUpdatedAt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.blog.Post.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.blog.Post.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.blog.Post} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blog.Post.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getTitle();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getContent();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getPrivate();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = message.getCreatedById();
  if (f !== 0) {
    writer.writeInt32(
      20,
      f
    );
  }
  f = message.getCreatedAt();
  if (f.length > 0) {
    writer.writeString(
      21,
      f
    );
  }
  f = message.getUpdatedById();
  if (f !== 0) {
    writer.writeInt32(
      22,
      f
    );
  }
  f = message.getUpdatedAt();
  if (f.length > 0) {
    writer.writeString(
      23,
      f
    );
  }
};


/**
 * optional int32 id = 1;
 * @return {number}
 */
proto.blog.Post.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.blog.Post.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string title = 3;
 * @return {string}
 */
proto.blog.Post.prototype.getTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.blog.Post.prototype.setTitle = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string content = 4;
 * @return {string}
 */
proto.blog.Post.prototype.getContent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.blog.Post.prototype.setContent = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional bool private = 5;
 * @return {boolean}
 */
proto.blog.Post.prototype.getPrivate = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 5, false));
};


/** @param {boolean} value */
proto.blog.Post.prototype.setPrivate = function(value) {
  jspb.Message.setProto3BooleanField(this, 5, value);
};


/**
 * optional int32 created_by_id = 20;
 * @return {number}
 */
proto.blog.Post.prototype.getCreatedById = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 20, 0));
};


/** @param {number} value */
proto.blog.Post.prototype.setCreatedById = function(value) {
  jspb.Message.setProto3IntField(this, 20, value);
};


/**
 * optional string created_at = 21;
 * @return {string}
 */
proto.blog.Post.prototype.getCreatedAt = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 21, ""));
};


/** @param {string} value */
proto.blog.Post.prototype.setCreatedAt = function(value) {
  jspb.Message.setProto3StringField(this, 21, value);
};


/**
 * optional int32 updated_by_id = 22;
 * @return {number}
 */
proto.blog.Post.prototype.getUpdatedById = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 22, 0));
};


/** @param {number} value */
proto.blog.Post.prototype.setUpdatedById = function(value) {
  jspb.Message.setProto3IntField(this, 22, value);
};


/**
 * optional string updated_at = 23;
 * @return {string}
 */
proto.blog.Post.prototype.getUpdatedAt = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 23, ""));
};


/** @param {string} value */
proto.blog.Post.prototype.setUpdatedAt = function(value) {
  jspb.Message.setProto3StringField(this, 23, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.blog.PostRequest.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.blog.PostRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.blog.PostRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.blog.PostRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blog.PostRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    post: (f = msg.getPost()) && proto.blog.Post.toObject(includeInstance, f),
    idsList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    userId: jspb.Message.getFieldWithDefault(msg, 3, 0),
    limit: jspb.Message.getFieldWithDefault(msg, 10, 0),
    offset: jspb.Message.getFieldWithDefault(msg, 11, 0),
    total: jspb.Message.getFieldWithDefault(msg, 12, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.blog.PostRequest}
 */
proto.blog.PostRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.blog.PostRequest;
  return proto.blog.PostRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.blog.PostRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.blog.PostRequest}
 */
proto.blog.PostRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.blog.Post;
      reader.readMessage(value,proto.blog.Post.deserializeBinaryFromReader);
      msg.setPost(value);
      break;
    case 2:
      var value = /** @type {!Array<number>} */ (reader.readPackedInt32());
      msg.setIdsList(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setUserId(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLimit(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setOffset(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTotal(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.blog.PostRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.blog.PostRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.blog.PostRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blog.PostRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPost();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.blog.Post.serializeBinaryToWriter
    );
  }
  f = message.getIdsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      2,
      f
    );
  }
  f = message.getUserId();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt32(
      10,
      f
    );
  }
  f = message.getOffset();
  if (f !== 0) {
    writer.writeInt32(
      11,
      f
    );
  }
  f = message.getTotal();
  if (f !== 0) {
    writer.writeInt32(
      12,
      f
    );
  }
};


/**
 * optional Post post = 1;
 * @return {?proto.blog.Post}
 */
proto.blog.PostRequest.prototype.getPost = function() {
  return /** @type{?proto.blog.Post} */ (
    jspb.Message.getWrapperField(this, proto.blog.Post, 1));
};


/** @param {?proto.blog.Post|undefined} value */
proto.blog.PostRequest.prototype.setPost = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.blog.PostRequest.prototype.clearPost = function() {
  this.setPost(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.blog.PostRequest.prototype.hasPost = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated int32 ids = 2;
 * @return {!Array<number>}
 */
proto.blog.PostRequest.prototype.getIdsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/** @param {!Array<number>} value */
proto.blog.PostRequest.prototype.setIdsList = function(value) {
  jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 */
proto.blog.PostRequest.prototype.addIds = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.blog.PostRequest.prototype.clearIdsList = function() {
  this.setIdsList([]);
};


/**
 * optional int32 user_id = 3;
 * @return {number}
 */
proto.blog.PostRequest.prototype.getUserId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.blog.PostRequest.prototype.setUserId = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 limit = 10;
 * @return {number}
 */
proto.blog.PostRequest.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.blog.PostRequest.prototype.setLimit = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int32 offset = 11;
 * @return {number}
 */
proto.blog.PostRequest.prototype.getOffset = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.blog.PostRequest.prototype.setOffset = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional int32 total = 12;
 * @return {number}
 */
proto.blog.PostRequest.prototype.getTotal = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/** @param {number} value */
proto.blog.PostRequest.prototype.setTotal = function(value) {
  jspb.Message.setProto3IntField(this, 12, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.blog.PostResponse.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.blog.PostResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.blog.PostResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.blog.PostResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blog.PostResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    post: (f = msg.getPost()) && proto.blog.Post.toObject(includeInstance, f),
    postsList: jspb.Message.toObjectList(msg.getPostsList(),
    proto.blog.Post.toObject, includeInstance),
    limit: jspb.Message.getFieldWithDefault(msg, 10, 0),
    offset: jspb.Message.getFieldWithDefault(msg, 11, 0),
    total: jspb.Message.getFieldWithDefault(msg, 12, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.blog.PostResponse}
 */
proto.blog.PostResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.blog.PostResponse;
  return proto.blog.PostResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.blog.PostResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.blog.PostResponse}
 */
proto.blog.PostResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.blog.Post;
      reader.readMessage(value,proto.blog.Post.deserializeBinaryFromReader);
      msg.setPost(value);
      break;
    case 2:
      var value = new proto.blog.Post;
      reader.readMessage(value,proto.blog.Post.deserializeBinaryFromReader);
      msg.addPosts(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLimit(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setOffset(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTotal(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.blog.PostResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.blog.PostResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.blog.PostResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blog.PostResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPost();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.blog.Post.serializeBinaryToWriter
    );
  }
  f = message.getPostsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.blog.Post.serializeBinaryToWriter
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt32(
      10,
      f
    );
  }
  f = message.getOffset();
  if (f !== 0) {
    writer.writeInt32(
      11,
      f
    );
  }
  f = message.getTotal();
  if (f !== 0) {
    writer.writeInt32(
      12,
      f
    );
  }
};


/**
 * optional Post post = 1;
 * @return {?proto.blog.Post}
 */
proto.blog.PostResponse.prototype.getPost = function() {
  return /** @type{?proto.blog.Post} */ (
    jspb.Message.getWrapperField(this, proto.blog.Post, 1));
};


/** @param {?proto.blog.Post|undefined} value */
proto.blog.PostResponse.prototype.setPost = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.blog.PostResponse.prototype.clearPost = function() {
  this.setPost(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.blog.PostResponse.prototype.hasPost = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated Post posts = 2;
 * @return {!Array<!proto.blog.Post>}
 */
proto.blog.PostResponse.prototype.getPostsList = function() {
  return /** @type{!Array<!proto.blog.Post>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.blog.Post, 2));
};


/** @param {!Array<!proto.blog.Post>} value */
proto.blog.PostResponse.prototype.setPostsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.blog.Post=} opt_value
 * @param {number=} opt_index
 * @return {!proto.blog.Post}
 */
proto.blog.PostResponse.prototype.addPosts = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.blog.Post, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.blog.PostResponse.prototype.clearPostsList = function() {
  this.setPostsList([]);
};


/**
 * optional int32 limit = 10;
 * @return {number}
 */
proto.blog.PostResponse.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.blog.PostResponse.prototype.setLimit = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int32 offset = 11;
 * @return {number}
 */
proto.blog.PostResponse.prototype.getOffset = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.blog.PostResponse.prototype.setOffset = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional int32 total = 12;
 * @return {number}
 */
proto.blog.PostResponse.prototype.getTotal = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/** @param {number} value */
proto.blog.PostResponse.prototype.setTotal = function(value) {
  jspb.Message.setProto3IntField(this, 12, value);
};


goog.object.extend(exports, proto.blog);
