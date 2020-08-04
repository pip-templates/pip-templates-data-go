package test_clients

import (
	"testing"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	clients1 "github.com/pip-templates/pip-templates-microservice-go/clients/version1"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	"github.com/stretchr/testify/assert"
)

type BeaconsClientV1Fixture struct {
	Beacon1 data1.BeaconV1
	Beacon2 data1.BeaconV1
	Beacon3 data1.BeaconV1
	client  clients1.IBeaconsClientV1
}

func NewBeaconsClientV1Fixture(client clients1.IBeaconsClientV1) *BeaconsClientV1Fixture {
	c := BeaconsClientV1Fixture{}
	c.Beacon1 = data1.BeaconV1{
		Id:      "1",
		Udi:     "00001",
		Type:    data1.AltBeacon,
		Site_id: "1",
		Label:   "TestBeacon1",
		Center:  data1.GeoPointV1{Type: "Point", Coordinates: [][]float32{{0.0, 0.0}}},
		Radius:  50,
	}
	c.Beacon2 = data1.BeaconV1{
		Id:      "2",
		Udi:     "00002",
		Type:    data1.IBeacon,
		Site_id: "1",
		Label:   "TestBeacon2",
		Center:  data1.GeoPointV1{Type: "Point", Coordinates: [][]float32{{2.0, 2.0}}},
		Radius:  70,
	}
	c.client = client
	return &c
}

func (c *BeaconsClientV1Fixture) testCreateBeacons(t *testing.T) {
	// Create the first beacon
	beacon, err := c.client.CreateBeacon("", &c.Beacon1)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, c.Beacon1.Udi, beacon.Udi)
	assert.Equal(t, c.Beacon1.Site_id, beacon.Site_id)
	assert.Equal(t, c.Beacon1.Type, beacon.Type)
	assert.Equal(t, c.Beacon1.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Create the second beacon
	beacon, err = c.client.CreateBeacon("", &c.Beacon2)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, c.Beacon2.Udi, beacon.Udi)
	assert.Equal(t, c.Beacon2.Site_id, beacon.Site_id)
	assert.Equal(t, c.Beacon2.Type, beacon.Type)
	assert.Equal(t, c.Beacon2.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)
}

func (c *BeaconsClientV1Fixture) TestCrudOperations(t *testing.T) {
	var beacon1 data1.BeaconV1

	// Create items
	c.testCreateBeacons(t)

	// Get all beacons
	page, err := c.client.GetBeacons("", cdata.NewEmptyFilterParams(), cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)
	beacon1 = *page.Data[0]

	// Update the beacon
	beacon1.Label = "ABC"
	beacon, err := c.client.UpdateBeacon("", &beacon1)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)
	assert.Equal(t, "ABC", beacon.Label)

	// Get beacon by udi
	beacon, err = c.client.GetBeaconByUdi("", beacon1.Udi)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)

	// Delete the beacon
	beacon, err = c.client.DeleteBeaconById("", beacon1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)

	// Try to get deleted beacon
	beacon, err = c.client.GetBeaconById("", beacon1.Id)
	assert.Nil(t, err)
	assert.Nil(t, beacon)
}

func (c *BeaconsClientV1Fixture) TestCalculatePosition(t *testing.T) {
	// Create items
	c.testCreateBeacons(t)

	// Calculate position for one beacon
	position, err := c.client.CalculatePosition("", "1", []string{"00001"})
	assert.Nil(t, err)
	assert.NotNil(t, position)
	assert.Equal(t, "Point", position.Type)
	assert.Equal(t, (float32)(0), position.Coordinates[0][0])
	assert.Equal(t, (float32)(0), position.Coordinates[0][1])
}
