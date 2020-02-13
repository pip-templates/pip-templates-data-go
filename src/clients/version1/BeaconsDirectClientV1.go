package clients

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	rpcclient "github.com/pip-services3-go/pip-services3-rpc-go/clients"
	bdata "github.com/pip-templates/pip-templates-microservice-go/src/data/version1"
)

type BeaconsDirectClientV1 struct {
	rpcclient.DirectClient
	specificController IBeaconsClientV1
}

func NewBeaconsDirectClientV1() *BeaconsDirectClientV1 {
	bdc := BeaconsDirectClientV1{}
	bdc.DirectClient = *rpcclient.NewDirectClient()
	bdc.DependencyResolver.Put("controller", cref.NewDescriptor("beacons", "controller", "*", "*", "1.0"))
	return &bdc
}

func (c *BeaconsDirectClientV1) SetReferences(references cref.IReferences) {
	c.DirectClient.SetReferences(references)

	specificController, ok := c.Controller.(IBeaconsClientV1)
	if !ok {
		panic("BeaconsDirectClientV1: Cant't resolv dependency 'controller' to IBeaconsClientV1")
	}
	c.specificController = specificController
}

func (c *BeaconsDirectClientV1) GetBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error) {
	timing := c.Instrument(correlationId, "beacons.get_beacons")
	res, err := c.specificController.GetBeacons(correlationId, filter, paging)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) GetBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.get_beacon_by_id")
	res, err := c.specificController.GetBeaconById(correlationId, beaconId)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) GetBeaconByUdi(correlationId string, udi string) (beacon *bdata.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.get_beacon_by_udi")
	res, err := c.specificController.GetBeaconByUdi(correlationId, udi)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) CalculatePosition(correlationId string, siteId string, udis []string) (position *bdata.GeoPointV1, err error) {
	timing := c.Instrument(correlationId, "beacons.calculate_position")
	res, err := c.specificController.CalculatePosition(correlationId, siteId, udis)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) CreateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.create_beacon")
	res, err = c.specificController.CreateBeacon(correlationId, beacon)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) UpdateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.update_beacon")
	res, err = c.specificController.UpdateBeacon(correlationId, beacon)
	timing.EndTiming()
	return res, err
}

func (c *BeaconsDirectClientV1) DeleteBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {
	timing := c.Instrument(correlationId, "beacons.dee_beacon_by_id")
	beacon, err = c.specificController.DeleteBeaconById(correlationId, beaconId)
	timing.EndTiming()
	return beacon, err
}
