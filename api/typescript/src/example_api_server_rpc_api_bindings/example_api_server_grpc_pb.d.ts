// GENERATED CODE -- DO NOT EDIT!

// package: example_api_server_api
// file: example_api_server.proto

import * as example_api_server_pb from "./example_api_server_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface IExampleAPIServerServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  isAvailable: grpc.MethodDefinition<google_protobuf_empty_pb.Empty, google_protobuf_empty_pb.Empty>;
  addPerson: grpc.MethodDefinition<example_api_server_pb.AddPersonArgs, google_protobuf_empty_pb.Empty>;
  getPerson: grpc.MethodDefinition<example_api_server_pb.GetPersonArgs, example_api_server_pb.GetPersonResponse>;
  incrementBooksRead: grpc.MethodDefinition<example_api_server_pb.IncrementBooksReadArgs, google_protobuf_empty_pb.Empty>;
}

export const ExampleAPIServerServiceService: IExampleAPIServerServiceService;

export interface IExampleAPIServerServiceServer extends grpc.UntypedServiceImplementation {
  isAvailable: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, google_protobuf_empty_pb.Empty>;
  addPerson: grpc.handleUnaryCall<example_api_server_pb.AddPersonArgs, google_protobuf_empty_pb.Empty>;
  getPerson: grpc.handleUnaryCall<example_api_server_pb.GetPersonArgs, example_api_server_pb.GetPersonResponse>;
  incrementBooksRead: grpc.handleUnaryCall<example_api_server_pb.IncrementBooksReadArgs, google_protobuf_empty_pb.Empty>;
}

export class ExampleAPIServerServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  isAvailable(argument: google_protobuf_empty_pb.Empty, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  isAvailable(argument: google_protobuf_empty_pb.Empty, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  isAvailable(argument: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  addPerson(argument: example_api_server_pb.AddPersonArgs, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  addPerson(argument: example_api_server_pb.AddPersonArgs, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  addPerson(argument: example_api_server_pb.AddPersonArgs, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  getPerson(argument: example_api_server_pb.GetPersonArgs, callback: grpc.requestCallback<example_api_server_pb.GetPersonResponse>): grpc.ClientUnaryCall;
  getPerson(argument: example_api_server_pb.GetPersonArgs, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<example_api_server_pb.GetPersonResponse>): grpc.ClientUnaryCall;
  getPerson(argument: example_api_server_pb.GetPersonArgs, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<example_api_server_pb.GetPersonResponse>): grpc.ClientUnaryCall;
  incrementBooksRead(argument: example_api_server_pb.IncrementBooksReadArgs, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  incrementBooksRead(argument: example_api_server_pb.IncrementBooksReadArgs, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  incrementBooksRead(argument: example_api_server_pb.IncrementBooksReadArgs, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
}
