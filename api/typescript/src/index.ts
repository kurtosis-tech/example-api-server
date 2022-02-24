export { EXAMPLE_API_SERVER_VERSION } from "./example_api_server_version/example_api_server_version";
export { LISTEN_PROTOCOL, LISTEN_PORT} from "./example_api_server_rpc_api_consts/example_api_server_rpc_api_consts"

// TODO: DELETE THIS LINE AFTER gRPC WEB IS IMPLEMENTED
export type { ExampleAPIServerServiceClient } from "./example_api_server_rpc_api_bindings/example_api_server_grpc_pb";

//Example API server RPC API Bindings
export type { ExampleAPIServerServiceClient as ExampleAPIServerServiceClientNode} from "./example_api_server_rpc_api_bindings/example_api_server_grpc_pb";
export type { ExampleAPIServerServiceClient as ExampleAPIServerServiceClientWeb} from "./example_api_server_rpc_api_bindings/example_api_server_grpc_web_pb";
export { AddPersonArgs, GetPersonArgs, GetPersonResponse, IncrementBooksReadArgs} from "./example_api_server_rpc_api_bindings/example_api_server_pb";

