package test_clients

import (
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	clients1 "github.com/pip-templates/pip-templates-microservice-go/clients/version1"
	logic "github.com/pip-templates/pip-templates-microservice-go/logic"
	persist "github.com/pip-templates/pip-templates-microservice-go/persistence"
)

func TestBeaconsDirectClientV1(t *testing.T) {
	var persistence *persist.BeaconsMemoryPersistence
	var controller *logic.BeaconsController
	var client *clients1.BeaconsDirectClientV1
	var fixture *BeaconsClientV1Fixture

	persistence = persist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = logic.NewBeaconsController()
	controller.Configure(cconf.NewEmptyConfigParams())
	client = clients1.NewBeaconsDirectClientV1()

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("beacons", "client", "direct", "default", "1.0"), client,
	)

	controller.SetReferences(references)
	client.SetReferences(references)
	fixture = NewBeaconsClientV1Fixture(client)

	persistence.Open("")
	defer persistence.Close("")

	t.Run("TestBeaconsDirectClientV1:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("TestBeaconsDirectClientV:1Calculate Positions", fixture.TestCalculatePosition)
}
