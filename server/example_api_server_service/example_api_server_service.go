package example_api_server_service

import (
	"context"
	"github.com/kurtosis-tech/example-api-server/api/golang/example_api_server_rpc_api_bindings"
	"github.com/kurtosis-tech/example-datastore-server/api/golang/datastore_rpc_api_bindings"
	"github.com/palantir/stacktrace"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

const (
	booksReadValueToIncrement uint64 = 1
	booksReadValueBase int = 10
	booksReadValueBitSize int = 32
)

type ExampleAPIServerService struct {
	// This embedding is required by gRPC
	example_api_server_rpc_api_bindings.UnimplementedExampleAPIServerServiceServer
	datastoreClient datastore_rpc_api_bindings.DatastoreServiceClient
}

func NewExampleAPIServerService(datastoreClient datastore_rpc_api_bindings.DatastoreServiceClient) *ExampleAPIServerService {
	return &ExampleAPIServerService{datastoreClient: datastoreClient}
}

func (exampleAPIServerService ExampleAPIServerService) IsAvailable(ctx context.Context, args *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (exampleAPIServerService ExampleAPIServerService) AddPerson(ctx context.Context, args *example_api_server_rpc_api_bindings.AddPersonArgs) (*emptypb.Empty, error) {

	personId := args.GetPersonId()

	// Make sure the person doesn't already exist
	exists, err := exampleAPIServerService.existsPersonInDatastore(ctx, personId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred checking if person with ID '%v' exists", personId)
	}
	if exists {
		return nil, stacktrace.NewError("A person with ID '%v' already exists", personId)
	}

	initialBooksReadValue := "0"

	upsertArgs := &datastore_rpc_api_bindings.UpsertArgs{
		Key:   personId,
		Value: initialBooksReadValue,
	}
	if _, err := exampleAPIServerService.datastoreClient.Upsert(ctx, upsertArgs); err != nil {
		return nil, stacktrace.NewError("An error occurred adding person with ID '%v' ", personId)
	}

	return &emptypb.Empty{}, nil
}

func (exampleAPIServerService ExampleAPIServerService) GetPerson(ctx context.Context, args *example_api_server_rpc_api_bindings.GetPersonArgs) (*example_api_server_rpc_api_bindings.GetPersonResponse, error) {
	personId := args.GetPersonId()

	// If the person doesn't already exist, return an error
	exists, err := exampleAPIServerService.existsPersonInDatastore(ctx, personId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred checking if person with ID '%v' exists", personId)
	}
	if !exists {
		return nil, stacktrace.NewError("No person with ID '%v' exists", personId)
	}

	getArgs := &datastore_rpc_api_bindings.GetArgs{
		Key: personId,
	}
	getResponse, err := exampleAPIServerService.datastoreClient.Get(ctx, getArgs)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred getting data for person with ID '%v'", personId)
	}

	getPersonResponse := &example_api_server_rpc_api_bindings.GetPersonResponse{
		BooksRead: getResponse.GetValue(),
	}

	return getPersonResponse, nil
}

func (exampleAPIServerService ExampleAPIServerService) IncrementBooksRead(ctx context.Context, args *example_api_server_rpc_api_bindings.IncrementBooksReadArgs) (*emptypb.Empty, error) {
	personId := args.GetPersonId()

	// Make sure the person exists
	exists, err := exampleAPIServerService.existsPersonInDatastore(ctx, personId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred checking if person with ID '%v' exists", personId)
	}
	if !exists {
		return nil, stacktrace.NewError("No person with ID '%v' exists", personId)
	}

	getArgs := &datastore_rpc_api_bindings.GetArgs{
		Key: personId,
	}
	getResponse, err := exampleAPIServerService.datastoreClient.Get(ctx, getArgs)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred getting data for person with ID '%v'", personId)
	}

	currentBooksReadValue := getResponse.GetValue()

	currentBooksReadValueU64, err := strconv.ParseUint(currentBooksReadValue, booksReadValueBase, booksReadValueBitSize)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred parsing string value '%v' with base '%v' and bit size '%v' to uint", currentBooksReadValue, booksReadValueBase, booksReadValueBitSize)
	}

	booksReadValueIncremented := currentBooksReadValueU64 + booksReadValueToIncrement
	booksReadValueIncrementedStr := strconv.FormatUint(booksReadValueIncremented, booksReadValueBase)

	upsertArgs := &datastore_rpc_api_bindings.UpsertArgs{
		Key:   personId,
		Value: booksReadValueIncrementedStr,
	}
	if _, err := exampleAPIServerService.datastoreClient.Upsert(ctx, upsertArgs); err != nil {
		return nil, stacktrace.NewError("An error occurred adding person with ID '%v' ", personId)
	}

	return &emptypb.Empty{}, nil
}

// ====================================================================================================
// 									   Private helper methods
// ====================================================================================================
func (exampleAPIServerService ExampleAPIServerService) existsPersonInDatastore(ctx context.Context, personId string) (bool, error) {
	existsArgs := &datastore_rpc_api_bindings.ExistsArgs{
		Key: personId,
	}
	existsResponse, err := exampleAPIServerService.datastoreClient.Exists(ctx, existsArgs)
	if err != nil {
		return false, stacktrace.Propagate(err, "An error occurred checking if person with ID '%v' already exists", personId)
	}

	return existsResponse.GetExists(), nil
}
