package test_persistence

import (
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	persist "github.com/pip-templates/pip-templates-microservice-go/persistence"
)

func TestBeaconsMemoryPersistence(t *testing.T) {
	var persistence *persist.BeaconsMemoryPersistence
	var fixture *BeaconsPersistenceFixture

	persistence = persist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())
	fixture = NewBeaconsPersistenceFixture(persistence)

	persistence.Open("")

	defer persistence.Close("")

	t.Run("BeaconsMemoryPersistence:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("BeaconsMemoryPersistence:Get with Filters", fixture.TestGetWithFilters)
}
