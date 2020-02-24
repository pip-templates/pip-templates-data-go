package test_persistence

import (
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	bpersist "github.com/pip-templates/pip-templates-microservice-go/persistence"
)

func TestBeaconsFilePersistence(t *testing.T) {
	var persistence *bpersist.BeaconsFilePersistence
	var fixture *BeaconsPersistenceFixture

	persistence = bpersist.NewBeaconsFilePersistence("../../persistence_data/beacons.test.json")
	persistence.Configure(cconf.NewEmptyConfigParams())
	fixture = NewBeaconsPersistenceFixture(persistence)

	opnErr := persistence.Open("")
	if opnErr == nil {
		persistence.Clear("")
	}

	defer persistence.Close("")

	t.Run("BeaconsFilePersistence:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("BeaconsFilePersistence:Get with Filters", fixture.TestGetWithFilters)
}
