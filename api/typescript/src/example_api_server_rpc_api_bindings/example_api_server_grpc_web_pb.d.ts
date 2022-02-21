import * as grpcWeb from 'grpc-web';

import * as example_api_server_pb from './example_api_server_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class ExampleAPIServerServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  isAvailable(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  addPerson(
    request: example_api_server_pb.AddPersonArgs,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  getPerson(
    request: example_api_server_pb.GetPersonArgs,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: example_api_server_pb.GetPersonResponse) => void
  ): grpcWeb.ClientReadableStream<example_api_server_pb.GetPersonResponse>;

  incrementBooksRead(
    request: example_api_server_pb.IncrementBooksReadArgs,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

}

export class ExampleAPIServerServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  isAvailable(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  addPerson(
    request: example_api_server_pb.AddPersonArgs,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  getPerson(
    request: example_api_server_pb.GetPersonArgs,
    metadata?: grpcWeb.Metadata
  ): Promise<example_api_server_pb.GetPersonResponse>;

  incrementBooksRead(
    request: example_api_server_pb.IncrementBooksReadArgs,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

}

