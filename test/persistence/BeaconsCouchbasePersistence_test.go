package test_persistence

import (
	"os"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	persist "github.com/pip-templates/pip-templates-microservice-go/persistence"
)

// Warning!
// Bucket - "beaconBucket" must be exists!!!
// Autocreate backet method not work in this release.
func TestBeaconsCouchbasePersistence(t *testing.T) {
	var persistence *persist.BeaconsCouchbasePersistence
	var fixture *BeaconsPersistenceFixture

	couchbaseUri := os.Getenv("COUCHBASE_URI")
	couchbaseHost := os.Getenv("COUCHBASE_HOST")
	if couchbaseHost == "" {
		couchbaseHost = "localhost"
	}
	couchbasePort := os.Getenv("COUCHBASE_PORT")
	if couchbasePort == "" {
		couchbasePort = "8091"
	}
	couchbaseUser := os.Getenv("COUCHBASE_USER")
	if couchbaseUser == "" {
		couchbaseUser = "Administrator"
	}
	couchbasePass := os.Getenv("COUCHBASE_PASS")
	if couchbasePass == "" {
		couchbasePass = "password"
	}

	if couchbaseUri == "" && couchbaseHost == "" {
		return
	}

	persistence = persist.NewBeaconsCouchbasePersistence()
	dbConfig := cconf.NewConfigParamsFromTuples(
		"options.auto_create", false, // true
		"options.auto_index", true,
		"connection.uri", couchbaseUri,
		"connection.host", couchbaseHost,
		"connection.port", couchbasePort,
		"connection.operation_timeout", 2,
		"connection.detailed_errcodes", 1,
		"credential.username", couchbaseUser,
		"credential.password", couchbasePass,
	)
	persistence.Configure(dbConfig)

	fixture = NewBeaconsPersistenceFixture(persistence)

	opnErr := persistence.Open("")
	if opnErr == nil {
		persistence.Clear("")
	}

	defer persistence.Close("")

	t.Run("BeaconsCouchbasePersistence:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("BeaconsCouchbasePersistence:Get with Filters", fixture.TestGetWithFilters)
}
