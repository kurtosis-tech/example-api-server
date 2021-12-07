// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var example_api_server_pb = require('./example_api_server_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_example_api_server_api_AddPersonArgs(arg) {
  if (!(arg instanceof example_api_server_pb.AddPersonArgs)) {
    throw new Error('Expected argument of type example_api_server_api.AddPersonArgs');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_example_api_server_api_AddPersonArgs(buffer_arg) {
  return example_api_server_pb.AddPersonArgs.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_example_api_server_api_GetPersonArgs(arg) {
  if (!(arg instanceof example_api_server_pb.GetPersonArgs)) {
    throw new Error('Expected argument of type example_api_server_api.GetPersonArgs');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_example_api_server_api_GetPersonArgs(buffer_arg) {
  return example_api_server_pb.GetPersonArgs.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_example_api_server_api_GetPersonResponse(arg) {
  if (!(arg instanceof example_api_server_pb.GetPersonResponse)) {
    throw new Error('Expected argument of type example_api_server_api.GetPersonResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_example_api_server_api_GetPersonResponse(buffer_arg) {
  return example_api_server_pb.GetPersonResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_example_api_server_api_IncrementBooksReadArgs(arg) {
  if (!(arg instanceof example_api_server_pb.IncrementBooksReadArgs)) {
    throw new Error('Expected argument of type example_api_server_api.IncrementBooksReadArgs');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_example_api_server_api_IncrementBooksReadArgs(buffer_arg) {
  return example_api_server_pb.IncrementBooksReadArgs.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}


var ExampleAPIServerServiceService = exports.ExampleAPIServerServiceService = {
  // Used to check service availability
isAvailable: {
    path: '/example_api_server_api.ExampleAPIServerService/IsAvailable',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Add a person ID into the book reads store
addPerson: {
    path: '/example_api_server_api.ExampleAPIServerService/AddPerson',
    requestStream: false,
    responseStream: false,
    requestType: example_api_server_pb.AddPersonArgs,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_example_api_server_api_AddPersonArgs,
    requestDeserialize: deserialize_example_api_server_api_AddPersonArgs,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Get a person books read value by person ID
getPerson: {
    path: '/example_api_server_api.ExampleAPIServerService/GetPerson',
    requestStream: false,
    responseStream: false,
    requestType: example_api_server_pb.GetPersonArgs,
    responseType: example_api_server_pb.GetPersonResponse,
    requestSerialize: serialize_example_api_server_api_GetPersonArgs,
    requestDeserialize: deserialize_example_api_server_api_GetPersonArgs,
    responseSerialize: serialize_example_api_server_api_GetPersonResponse,
    responseDeserialize: deserialize_example_api_server_api_GetPersonResponse,
  },
  // Increment in 1 the amount of person books read by person ID
incrementBooksRead: {
    path: '/example_api_server_api.ExampleAPIServerService/IncrementBooksRead',
    requestStream: false,
    responseStream: false,
    requestType: example_api_server_pb.IncrementBooksReadArgs,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_example_api_server_api_IncrementBooksReadArgs,
    requestDeserialize: deserialize_example_api_server_api_IncrementBooksReadArgs,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.ExampleAPIServerServiceClient = grpc.makeGenericClientConstructor(ExampleAPIServerServiceService);
