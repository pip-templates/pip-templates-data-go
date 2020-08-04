package clients

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	clients "github.com/pip-services3-go/pip-services3-rpc-go/clients"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

type BeaconsDirectClientV1 struct {
	clients.DirectClient
	controller IBeaconsClientV1
}

func NewBeaconsDirectClientV1() *BeaconsDirectClientV1 {
	c := BeaconsDirectClientV1{}
	c.DirectClient = *clients.NewDirectClient()
	c.DependencyResolver.Put("controller", cref.NewDescriptor("beacons", "controller", "*", "*", "1.0"))
	return &c
}

func (c *BeaconsDirectClientV1) SetReferences(references cref.IReferences) {
	c.DirectClient.SetReferences(references)

	controller, ok := c.Controller.(IBeaconsClientV1)
	if !ok {
		panic("BeaconsDirectClientV1: Cant't resolv dependency 'controller' to IBeaconsClientV1")
	}
	c.controller = controller
}

func (c *BeaconsDirectClientV1) GetBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *data1.BeaconV1DataPage, err error) {
	timing := c.Instrument(correlationId, "beacons.get_beacons")
	res, err := c.controller.GetBeacons(correlationId, filter, paging)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) GetBeaconById(correlationId string, beaconId string) (beacon *data1.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.get_beacon_by_id")
	res, err := c.controller.GetBeaconById(correlationId, beaconId)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) GetBeaconByUdi(correlationId string, udi string) (beacon *data1.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.get_beacon_by_udi")
	res, err := c.controller.GetBeaconByUdi(correlationId, udi)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) CalculatePosition(correlationId string, siteId string, udis []string) (position *data1.GeoPointV1, err error) {
	timing := c.Instrument(correlationId, "beacons.calculate_position")
	res, err := c.controller.CalculatePosition(correlationId, siteId, udis)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) CreateBeacon(correlationId string, beacon *data1.BeaconV1) (res *data1.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.create_beacon")
	res, err = c.controller.CreateBeacon(correlationId, beacon)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) UpdateBeacon(correlationId string, beacon *data1.BeaconV1) (res *data1.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.update_beacon")
	res, err = c.controller.UpdateBeacon(correlationId, beacon)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) DeleteBeaconById(correlationId string, beaconId string) (beacon *data1.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.delete_beacon_by_id")
	beacon, err = c.controller.DeleteBeaconById(correlationId, beaconId)
	timing.EndTiming()
	return beacon, err
}
