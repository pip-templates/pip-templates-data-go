package test_persistence

import (
	"os"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	persist "github.com/pip-templates/pip-templates-microservice-go/persistence"
)

func TestBeaconsMongoDbPersistence(t *testing.T) {

	var persistence *persist.BeaconsMongoDbPersistence
	var fixture *BeaconsPersistenceFixture

	mongoUri := os.Getenv("MONGO_SERVICE_URI")
	mongoHost := os.Getenv("MONGO_SERVICE_HOST")

	if mongoHost == "" {
		mongoHost = "localhost"
	}
	mongoPort := os.Getenv("MONGO_SERVICE_PORT")
	if mongoPort == "" {
		mongoPort = "27017"
	}

	mongoDatabase := os.Getenv("MONGO_SERVICE_DB")
	if mongoDatabase == "" {
		mongoDatabase = "test"
	}

	// Exit if mongo connection is not set
	if mongoUri == "" && mongoHost == "" {
		return
	}

	persistence = persist.NewBeaconsMongoDbPersistence()
	persistence.Configure(cconf.NewConfigParamsFromTuples(
		"connection.uri", mongoUri,
		"connection.host", mongoHost,
		"connection.port", mongoPort,
		"connection.database", mongoDatabase,
	))

	fixture = NewBeaconsPersistenceFixture(persistence)

	opnErr := persistence.Open("")
	if opnErr == nil {
		persistence.Clear("")
	}

	defer persistence.Close("")

	t.Run("BeaconsMongoDbPersistence:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("BeaconsMongoDbPersistence:Get with Filters", fixture.TestGetWithFilters)

}
