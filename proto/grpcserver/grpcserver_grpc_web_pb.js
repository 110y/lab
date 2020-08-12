/**
 * @fileoverview gRPC-Web generated client stub for grpcserver
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.grpcserver = require('./grpcserver_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.grpcserver.InfoClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.grpcserver.InfoPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.grpcserver.ServerInfoResponse>}
 */
const methodDescriptor_Info_ServerInfo = new grpc.web.MethodDescriptor(
  '/grpcserver.Info/ServerInfo',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.grpcserver.ServerInfoResponse,
  /** @param {!proto.google.protobuf.Empty} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.grpcserver.ServerInfoResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.grpcserver.ServerInfoResponse>}
 */
const methodInfo_Info_ServerInfo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.grpcserver.ServerInfoResponse,
  /** @param {!proto.google.protobuf.Empty} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.grpcserver.ServerInfoResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.grpcserver.ServerInfoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.grpcserver.ServerInfoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.grpcserver.InfoClient.prototype.serverInfo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/grpcserver.Info/ServerInfo',
      request,
      metadata || {},
      methodDescriptor_Info_ServerInfo,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.grpcserver.ServerInfoResponse>}
 *     A native promise that resolves to the response
 */
proto.grpcserver.InfoPromiseClient.prototype.serverInfo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/grpcserver.Info/ServerInfo',
      request,
      metadata || {},
      methodDescriptor_Info_ServerInfo);
};


module.exports = proto.grpcserver;

