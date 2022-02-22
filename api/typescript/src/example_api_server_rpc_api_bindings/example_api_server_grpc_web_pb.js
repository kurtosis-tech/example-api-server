/**
 * @fileoverview gRPC-Web generated client stub for example_api_server_api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.example_api_server_api = require('./example_api_server_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.example_api_server_api.ExampleAPIServerServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.example_api_server_api.ExampleAPIServerServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ExampleAPIServerService_IsAvailable = new grpc.web.MethodDescriptor(
  '/example_api_server_api.ExampleAPIServerService/IsAvailable',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.example_api_server_api.ExampleAPIServerServiceClient.prototype.isAvailable =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/example_api_server_api.ExampleAPIServerService/IsAvailable',
      request,
      metadata || {},
      methodDescriptor_ExampleAPIServerService_IsAvailable,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.example_api_server_api.ExampleAPIServerServicePromiseClient.prototype.isAvailable =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/example_api_server_api.ExampleAPIServerService/IsAvailable',
      request,
      metadata || {},
      methodDescriptor_ExampleAPIServerService_IsAvailable);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.example_api_server_api.AddPersonArgs,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ExampleAPIServerService_AddPerson = new grpc.web.MethodDescriptor(
  '/example_api_server_api.ExampleAPIServerService/AddPerson',
  grpc.web.MethodType.UNARY,
  proto.example_api_server_api.AddPersonArgs,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.example_api_server_api.AddPersonArgs} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.example_api_server_api.AddPersonArgs} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.example_api_server_api.ExampleAPIServerServiceClient.prototype.addPerson =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/example_api_server_api.ExampleAPIServerService/AddPerson',
      request,
      metadata || {},
      methodDescriptor_ExampleAPIServerService_AddPerson,
      callback);
};


/**
 * @param {!proto.example_api_server_api.AddPersonArgs} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.example_api_server_api.ExampleAPIServerServicePromiseClient.prototype.addPerson =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/example_api_server_api.ExampleAPIServerService/AddPerson',
      request,
      metadata || {},
      methodDescriptor_ExampleAPIServerService_AddPerson);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.example_api_server_api.GetPersonArgs,
 *   !proto.example_api_server_api.GetPersonResponse>}
 */
const methodDescriptor_ExampleAPIServerService_GetPerson = new grpc.web.MethodDescriptor(
  '/example_api_server_api.ExampleAPIServerService/GetPerson',
  grpc.web.MethodType.UNARY,
  proto.example_api_server_api.GetPersonArgs,
  proto.example_api_server_api.GetPersonResponse,
  /**
   * @param {!proto.example_api_server_api.GetPersonArgs} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.example_api_server_api.GetPersonResponse.deserializeBinary
);


/**
 * @param {!proto.example_api_server_api.GetPersonArgs} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.example_api_server_api.GetPersonResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.example_api_server_api.GetPersonResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.example_api_server_api.ExampleAPIServerServiceClient.prototype.getPerson =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/example_api_server_api.ExampleAPIServerService/GetPerson',
      request,
      metadata || {},
      methodDescriptor_ExampleAPIServerService_GetPerson,
      callback);
};


/**
 * @param {!proto.example_api_server_api.GetPersonArgs} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.example_api_server_api.GetPersonResponse>}
 *     Promise that resolves to the response
 */
proto.example_api_server_api.ExampleAPIServerServicePromiseClient.prototype.getPerson =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/example_api_server_api.ExampleAPIServerService/GetPerson',
      request,
      metadata || {},
      methodDescriptor_ExampleAPIServerService_GetPerson);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.example_api_server_api.IncrementBooksReadArgs,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ExampleAPIServerService_IncrementBooksRead = new grpc.web.MethodDescriptor(
  '/example_api_server_api.ExampleAPIServerService/IncrementBooksRead',
  grpc.web.MethodType.UNARY,
  proto.example_api_server_api.IncrementBooksReadArgs,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.example_api_server_api.IncrementBooksReadArgs} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.example_api_server_api.IncrementBooksReadArgs} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.example_api_server_api.ExampleAPIServerServiceClient.prototype.incrementBooksRead =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/example_api_server_api.ExampleAPIServerService/IncrementBooksRead',
      request,
      metadata || {},
      methodDescriptor_ExampleAPIServerService_IncrementBooksRead,
      callback);
};


/**
 * @param {!proto.example_api_server_api.IncrementBooksReadArgs} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.example_api_server_api.ExampleAPIServerServicePromiseClient.prototype.incrementBooksRead =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/example_api_server_api.ExampleAPIServerService/IncrementBooksRead',
      request,
      metadata || {},
      methodDescriptor_ExampleAPIServerService_IncrementBooksRead);
};


module.exports = proto.example_api_server_api;

