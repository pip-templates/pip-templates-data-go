package test_persistence

import (
	"testing"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	bdata "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	bpersist "github.com/pip-templates/pip-templates-microservice-go/persistence"
	"github.com/stretchr/testify/assert"
)

type BeaconsPersistenceFixture struct {
	Beacon1     bdata.BeaconV1
	Beacon2     bdata.BeaconV1
	Beacon3     bdata.BeaconV1
	persistence bpersist.IBeaconsPersistence
}

func NewBeaconsPersistenceFixture(persistence bpersist.IBeaconsPersistence) *BeaconsPersistenceFixture {
	bpf := BeaconsPersistenceFixture{}
	bpf.Beacon1 = bdata.BeaconV1{
		Id:      "1",
		Udi:     "00001",
		Type:    bdata.BeaconTypeV1.AltBeacon,
		Site_id: "1",
		Label:   "TestBeacon1",
		Center:  bdata.GeoPointV1{Type: "Point", Lat: 0, Lng: 0},
		Radius:  50,
	}
	bpf.Beacon2 = bdata.BeaconV1{
		Id:      "2",
		Udi:     "00002",
		Type:    bdata.BeaconTypeV1.IBeacon,
		Site_id: "1",
		Label:   "TestBeacon2",
		Center:  bdata.GeoPointV1{Type: "Point", Lat: 2, Lng: 2},
		Radius:  70,
	}
	bpf.Beacon3 = bdata.BeaconV1{
		Id:      "3",
		Udi:     "00003",
		Type:    bdata.BeaconTypeV1.AltBeacon,
		Site_id: "2",
		Label:   "TestBeacon3",
		Center:  bdata.GeoPointV1{Type: "Point", Lat: 10, Lng: 10},
		Radius:  50,
	}
	bpf.persistence = persistence
	return &bpf
}

func (c *BeaconsPersistenceFixture) testCreateBeacons(t *testing.T) {

	// Create the first beacon
	beacon, err := c.persistence.Create("", c.Beacon1)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, c.Beacon1.Udi, beacon.Udi)
	assert.Equal(t, c.Beacon1.Site_id, beacon.Site_id)
	assert.Equal(t, c.Beacon1.Type, beacon.Type)
	assert.Equal(t, c.Beacon1.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Create the second beacon
	beacon, err = c.persistence.Create("", c.Beacon2)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, c.Beacon2.Udi, beacon.Udi)
	assert.Equal(t, c.Beacon2.Site_id, beacon.Site_id)
	assert.Equal(t, c.Beacon2.Type, beacon.Type)
	assert.Equal(t, c.Beacon2.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Create the third beacon
	beacon, err = c.persistence.Create("", c.Beacon3)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, c.Beacon3.Udi, beacon.Udi)
	assert.Equal(t, c.Beacon3.Site_id, beacon.Site_id)
	assert.Equal(t, c.Beacon3.Type, beacon.Type)
	assert.Equal(t, c.Beacon3.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)
}

func (c *BeaconsPersistenceFixture) TestCrudOperations(t *testing.T) {
	var beacon1 bdata.BeaconV1

	// Create items
	c.testCreateBeacons(t)

	// Get all beacons
	page, err := c.persistence.GetPageByFilter("", cdata.NewEmptyFilterParams(), cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 3)
	beacon1 = *page.Data[0]

	// Update the beacon
	beacon1.Label = "ABC"
	beacon, err := c.persistence.Update("", beacon1)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)
	assert.Equal(t, "ABC", beacon.Label)

	// Get beacon by udi
	beacon, err = c.persistence.GetOneByUdi("", beacon1.Udi)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)

	// Delete the beacon
	beacon, err = c.persistence.DeleteById("", beacon1.Id)
	assert.Nil(t, err)
	assert.NotNil(t, beacon)
	assert.Equal(t, beacon1.Id, beacon.Id)

	// Try to get deleted beacon
	beacon, err = c.persistence.GetOneById("", beacon1.Id)
	assert.Nil(t, err)
	assert.Nil(t, beacon)

}

func (c *BeaconsPersistenceFixture) TestGetWithFilters(t *testing.T) {

	// Create items
	c.testCreateBeacons(t)

	// Filter by id
	page, err := c.persistence.GetPageByFilter("",
		cdata.NewFilterParamsFromTuples(
			"id", "1",
		),
		cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.Len(t, page.Data, 1)

	// Filter by udi
	page, err = c.persistence.GetPageByFilter(
		"",
		cdata.NewFilterParamsFromTuples(
			"udi", "00002",
		),
		cdata.NewEmptyPagingParams())
	assert.Nil(t, err)
	assert.Len(t, page.Data, 1)

	// Filter by udis
	page, err = c.persistence.GetPageByFilter(
		"",
		cdata.NewFilterParamsFromTuples(
			"udis", "00001,00003",
		),
		cdata.NewEmptyPagingParams())

	assert.Nil(t, err)
	assert.Len(t, page.Data, 2)

	// Filter by site_id
	page, err = c.persistence.GetPageByFilter(
		"",
		cdata.NewFilterParamsFromTuples(
			"site_id", "1",
		),
		cdata.NewEmptyPagingParams())

	assert.Nil(t, err)
	assert.Len(t, page.Data, 2)
}
