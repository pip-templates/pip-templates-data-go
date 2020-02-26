package logic

import (
	ccomand "github.com/pip-services3-go/pip-services3-commons-go/commands"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	bdata "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	bpersist "github.com/pip-templates/pip-templates-microservice-go/persistence"
)

type BeaconsController struct {
	persistence bpersist.IBeaconsPersistence
	commandSet  *BeaconsCommandSet
}

func NewBeaconsController() *BeaconsController {
	bc := BeaconsController{}
	return &bc
}

func (c *BeaconsController) Configure(config *cconf.ConfigParams) {

}

func (c *BeaconsController) SetReferences(references cref.IReferences) {
	ref, err := references.GetOneRequired(cref.NewDescriptor("beacons", "persistence", "*", "*", "1.0"))
	if ref != nil && err == nil {
		c.persistence = ref.(bpersist.IBeaconsPersistence)
	}
}

func (c *BeaconsController) GetCommandSet() *ccomand.CommandSet {
	if c.commandSet == nil {
		c.commandSet = NewBeaconsCommandSet(c)
	}
	return &c.commandSet.CommandSet
}

func (c *BeaconsController) GetBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error) {
	return c.persistence.GetPageByFilter(correlationId, filter, paging)
}

func (c *BeaconsController) GetBeaconById(correlationId string, beaconId string) (page *bdata.BeaconV1, err error) {
	return c.persistence.GetOneById(correlationId, beaconId)
}

func (c *BeaconsController) GetBeaconByUdi(correlationId string, beaconId string) (page *bdata.BeaconV1, err error) {
	return c.persistence.GetOneByUdi(correlationId, beaconId)
}

func (c *BeaconsController) CalculatePosition(correlationId string, siteId string, udis []string) (position *bdata.GeoPointV1, err error) {
	beacons := make([]bdata.BeaconV1, 0, 0)
	pos := bdata.GeoPointV1{
		Type:        "Point",
		Coordinates: make([][]float32, 1, 1),
	}

	pos.Coordinates[0] = make([]float32, 2, 2)

	if udis == nil || len(udis) == 0 {

		return nil, nil
	}

	page, getErr := c.persistence.GetPageByFilter(
		correlationId, cdata.NewFilterParamsFromTuples(
			"site_id", siteId,
			"udis", udis),
		cdata.NewEmptyPagingParams())

	if getErr != nil || page == nil {
		return nil, getErr
	}
	for _, v := range page.Data {
		beacons = append(beacons, *v)
	}

	var lat float32 = 0
	var lng float32 = 0
	var count = 0

	for _, beacon := range beacons {
		if beacon.Center.Type == "Point" {
			lng += beacon.Center.Coordinates[0][0]
			lat += beacon.Center.Coordinates[0][1]
			count += 1
		}
	}

	if count > 0 {
		pos.Type = "Point"
		pos.Coordinates[0][0] = lng / (float32)(count)
		pos.Coordinates[0][1] = lat / (float32)(count)
	}

	return &pos, nil
}

func (c *BeaconsController) CreateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {

	if beacon.Id == "" {
		beacon.Id = cdata.IdGenerator.NextLong()
	}

	if beacon.Type == "" {
		beacon.Type = bdata.BeaconTypeV1.Unknown
	}

	return c.persistence.Create(correlationId, beacon)
}

func (c *BeaconsController) UpdateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {

	if beacon.Type == "" {
		beacon.Type = bdata.BeaconTypeV1.Unknown
	}

	return c.persistence.Update(correlationId, beacon)
}

func (c *BeaconsController) DeleteBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {
	return c.persistence.DeleteById(correlationId, beaconId)
}
