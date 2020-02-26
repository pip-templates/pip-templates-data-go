package test_clients

import (
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	bclients "github.com/pip-templates/pip-templates-microservice-go/clients/version1"
	blogic "github.com/pip-templates/pip-templates-microservice-go/logic"
	bpersist "github.com/pip-templates/pip-templates-microservice-go/persistence"
	bservices "github.com/pip-templates/pip-templates-microservice-go/services/version1"
)

func TestBeaconsGrpcClientV1(t *testing.T) {

	var persistence *bpersist.BeaconsMemoryPersistence
	var controller *blogic.BeaconsController
	var service *bservices.BeaconsGrpcServiceV1
	var client *bclients.BeaconsGrpcClientV1
	var fixture *BeaconsClientV1Fixture

	persistence = bpersist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = blogic.NewBeaconsController()
	controller.Configure(cconf.NewEmptyConfigParams())

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3002",
		"connection.host", "localhost",
	)

	service = bservices.NewBeaconsGrpcServiceV1()
	service.Configure(httpConfig)

	client = bclients.NewBeaconsGrpcClientV1()
	client.Configure(httpConfig)

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("beacons", "service", "grpc", "default", "1.0"), service,
		cref.NewDescriptor("beacons", "client", "grpc", "default", "1.0"), client,
	)
	controller.SetReferences(references)
	service.SetReferences(references)
	client.SetReferences(references)

	fixture = NewBeaconsClientV1Fixture(client)

	opnErr := persistence.Open("")
	if opnErr != nil {
		panic("TestBeaconsGrpcClientV1:Error open persistence!")
	}

	opnErr = service.Open("")
	if opnErr != nil {
		panic("TestBeaconsGrpcClientV1:Error open service!")
	}

	client.Open("")

	defer client.Close("")
	defer service.Close("")
	defer persistence.Close("")

	t.Run("BeaconsGrpcClientV1:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("BeaconsGrpcClientV1:1Calculate Positions", fixture.TestCalculatePosition)

}
