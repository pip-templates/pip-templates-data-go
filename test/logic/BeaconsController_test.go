package test_logic

import (
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	bdata "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	blogic "github.com/pip-templates/pip-templates-microservice-go/logic"
	bpersist "github.com/pip-templates/pip-templates-microservice-go/persistence"
	"github.com/stretchr/testify/assert"
)

var Beacon1 bdata.BeaconV1 = bdata.BeaconV1{
	Id:      "1",
	Udi:     "00001",
	Type:    bdata.BeaconTypeV1.AltBeacon,
	Site_id: "1",
	Label:   "TestBeacon1",
	Center:  bdata.GeoPointV1{Type: "Point", Lat: 0, Lng: 0},
	Radius:  50,
}

var Beacon2 bdata.BeaconV1 = bdata.BeaconV1{
	Id:      "2",
	Udi:     "00002",
	Type:    bdata.BeaconTypeV1.IBeacon,
	Site_id: "1",
	Label:   "TestBeacon2",
	Center:  bdata.GeoPointV1{Type: "Point", Lat: 2, Lng: 2},
	Radius:  70,
}

var persistence *bpersist.BeaconsMemoryPersistence
var controller *blogic.BeaconsController

func TestBeaconsController(t *testing.T) {

	persistence = bpersist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = blogic.NewBeaconsController()
	controller.Configure(cconf.NewEmptyConfigParams())

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
	)

	controller.SetReferences(references)

	persistence.Open("")

	defer persistence.Close("")

	t.Run("BeaconsController:CRUD Operations", CrudOperations)
	persistence.Clear("")
	t.Run("BeaconsController:Calculate Positions", CalculatePositions)
}

func CrudOperations(t *testing.T) {
	var beacon1 bdata.BeaconV1

	// Create the first beacon
	beacon, err := controller.CreateBeacon("", Beacon1)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon1.Udi, beacon.Udi)
	assert.Equal(t, Beacon1.Site_id, beacon.Site_id)
	assert.Equal(t, Beacon1.Type, beacon.Type)
	assert.Equal(t, Beacon1.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Create the second beacon
	beacon, err = controller.CreateBeacon("", Beacon2)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon2.Udi, beacon.Udi)
	assert.Equal(t, Beacon2.Site_id, beacon.Site_id)
	assert.Equal(t, Beacon2.Type, beacon.Type)
	assert.Equal(t, Beacon2.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Get all beacons
	page, err := controller.GetBeacons("", cdata.NewEmptyFilterParams(), cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)
	beacon1 = *page.Data[0]

	// Update the beacon
	beacon1.Label = "ABC"
	beacon, err = controller.UpdateBeacon("", beacon1)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)
	assert.Equal(t, "ABC", beacon.Label)

	// Get beacon by udi
	beacon, err = controller.GetBeaconByUdi("", beacon1.Udi)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)

	// Delete the beacon
	beacon, err = controller.DeleteBeaconById("", beacon1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)

	// Try to get deleted beacon
	beacon, err = controller.GetBeaconById("", beacon1.Id)
	assert.Nil(t, err)
	assert.Nil(t, beacon)
}

func CalculatePositions(t *testing.T) {

	// Create the first beacon
	beacon, err := controller.CreateBeacon("", Beacon1)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon1.Udi, beacon.Udi)
	assert.Equal(t, Beacon1.Site_id, beacon.Site_id)
	assert.Equal(t, Beacon1.Type, beacon.Type)
	assert.Equal(t, Beacon1.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Create the second beacon
	beacon, err = controller.CreateBeacon("", Beacon2)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon2.Udi, beacon.Udi)
	assert.Equal(t, Beacon2.Site_id, beacon.Site_id)
	assert.Equal(t, Beacon2.Type, beacon.Type)
	assert.Equal(t, Beacon2.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Calculate position for one beacon
	position, err := controller.CalculatePosition("", "1", []string{"00001"})
	assert.Nil(t, err)
	assert.NotNil(t, position)
	assert.Equal(t, "Point", position.Type)
	assert.Equal(t, (float32)(0.0), position.Lat)
	assert.Equal(t, (float32)(0.0), position.Lng)

	// Calculate position for two beacons
	position, err = controller.CalculatePosition("", "1", []string{"00001", "00002"})
	assert.Nil(t, err)
	assert.NotNil(t, position)
	assert.Equal(t, "Point", position.Type)
	assert.Equal(t, (float32)(1.0), position.Lat)
	assert.Equal(t, (float32)(1.0), position.Lng)
}
