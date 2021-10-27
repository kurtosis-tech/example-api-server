/*
 * Copyright (c) 2021 - present Kurtosis Technologies Inc.
 * All Rights Reserved.
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kurtosis-tech/example-api-server/api/golang/example_api_server_rpc_api_bindings"
	"github.com/kurtosis-tech/example-api-server/server/example_api_server_service"
	"github.com/kurtosis-tech/example-datastore-server/api/golang/datastore_rpc_api_bindings"
	minimal_grpc_server "github.com/kurtosis-tech/minimal-grpc-server/golang/server"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
	"time"
)

const (
	successExitCode = 0
	failureExitCode = 1

	grpcServerStopGracePeriod = 5 * time.Second

	listenPort = 2434
	listenProtocol = "tcp"
)

var defaultKurtosisLogLevel = logrus.InfoLevel.String()

// Fields are public for JSON unmarshalling
type serverConfig struct {
	DatastoreIp string	`json:"datastoreIp"`
	DatastorePort int	`json:"datastorePort"`
}

func main() {
	// NOTE: we'll want to change the ForceColors to false if we ever want structured logging
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	err := runMain()
	if err != nil {
		logrus.Errorf("An error occurred when running the main function")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}
	os.Exit(successExitCode)
}

func runMain () error {
	logLevel, err := logrus.ParseLevel(defaultKurtosisLogLevel)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred parsing the log level string '%v':", defaultKurtosisLogLevel)
	}
	logrus.SetLevel(logLevel)

	configFilepathArg := flag.String("config", "", "Filepath to the config file")
	flag.Parse()

	configFileBytes, err := ioutil.ReadFile(*configFilepathArg)
	if err != nil {
		logrus.Errorf("An error occurred reading the config filepath: %v", err)
		return stacktrace.Propagate(err,"An error occurred reading the config filepath: %v", configFilepathArg)
	}

	var config serverConfig
	if err := json.Unmarshal(configFileBytes, &config); err != nil {
		return stacktrace.Propagate(err,"An error occurred deserializing the config file: %v", configFilepathArg)
	}
	if err := config.validate(); err != nil {
		return stacktrace.Propagate(err,"An error occurred validating the config file: %v", configFilepathArg)
	}

	datastoreURL := fmt.Sprintf(
		"%v:%v",
		config.DatastoreIp,
		config.DatastorePort,
	)

	conn, err := grpc.Dial(datastoreURL, grpc.WithInsecure())
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred dialling the datastore container via its URL")
	}
	defer conn.Close()

	datastoreServiceClient := datastore_rpc_api_bindings.NewDatastoreServiceClient(conn)

	exampleAPIServerService := example_api_server_service.NewExampleAPIServerService(datastoreServiceClient)

	exampleAPIServerServiceRegistrationFunc := func(grpcServer *grpc.Server) {
		example_api_server_rpc_api_bindings.RegisterExampleAPIServerServiceServer(grpcServer, exampleAPIServerService)
	}
	exampleAPIServer := minimal_grpc_server.NewMinimalGRPCServer(
		uint32(listenPort),
		listenProtocol,
		grpcServerStopGracePeriod,
		[]func(*grpc.Server){
			exampleAPIServerServiceRegistrationFunc,
		},
	)

	logrus.Info("Running server...")
	if err := exampleAPIServer.Run(); err != nil {
		return stacktrace.Propagate(err, "An error occurred running the server.")
	}
	return nil
}

// ====================================================================================================
// 									   Private helper methods
// ====================================================================================================
func (config serverConfig) validate() error {
	if config.DatastoreIp == "" {
		return stacktrace.NewError("No datastore IP provided")
	}
	if config.DatastorePort == 0 {
		return stacktrace.NewError("No datastore port provided")
	}
	return nil
}
