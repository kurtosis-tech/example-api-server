# 0.3.4

### Fixes
* Refactored 'ExampleAPIServerServiceClient' from 'export type' to only 'export', as we need the class itself and not it`s type.

# 0.3.3

### Features
* Added gRPC web proto bindings and made small refactors


# 0.3.2
### Fixes
* Actually enable debug logging

# 0.3.1
### Changes
* Enable debug logging

# 0.3.0
### Changes
* Regenerate the API libs with `@grpc/grpc-js` rather than `grpc`

### Fixes
* `@types/google-protobuf` is now correctly declared as a runtime dependency

### Breaking Changes
* The Typescript API now uses `@grpc/grpc-js` rather than `grpc`
    * Users should switch their own TS code calling the datastore server API to `@grpc/grpc-js`

# 0.2.2
### Features
* Activated `publish_api_typescript` Circle Ci job to publish the Typescript version of the API

# 0.2.1
### Features
* Upgraded `Example API server client` in the server project to the latest version [Example API server 0.2.0](https://github.com/kurtosis-tech/example-api-server/blob/master/docs/changelog.md#020)

# 0.2.0
### Breaking Changes
* The `GetPersonResponse.books_read` field is now a `uint32`, rather than an `string`
    * Users should modify the value type where they are setting this value

### Features 
* Added `example_api_server_rpc_api_consts` which contains constants related to the example API server (e.g. ListenProtocol)

# 0.1.1
### Features
* Added `example_api_server_version.go` and `example_api_server_version_ts` which contains the server version value
* Added `update-own-version-constants.sh` script to automatically update the version value in both files
* Updated `release.sh` script to execute `update-own-version-constants.sh` during building process
* Created `ExampleAPIServerService` which implements the service on the server side
* Configured the `RPC` server that will handle the `ExampleAPIServerService` requests
* Added logic in the server to get the `datastore` config information before start it
* Added `Dockerfile` which defines `example-api-server` Docker image
* Added `.dockerignore` to improve velocity during Docker buildings
* Added `get-docker-image-tag.sh` script to automatically set Docker image tag
* Added `build.sh` script which builds the `example-api-server` binary field and builds the Docker image
* Updated Circle CI configuration adding two jobs: `check_server_code` and `push_server_artifacts` to check server code and publish the image in Docker Hub
* Upgraded `Example Datastore Server` dependency to the latest version [Example Datastore Server 0.3.0](https://github.com/kurtosis-tech/example-datastore-server/blob/master/CHANGELOG.md#030)

# 0.1.0
### Features
* Created example API server protobuf service definition in `example_api_server.proto` file
* Added `regenerate-protobuf-output.sh` shell script to regenerate rpc Golang and Typescript binding files
* Updated `index.ts` file to export Typescript binding class
* Configured `package.json` file with the repository information
