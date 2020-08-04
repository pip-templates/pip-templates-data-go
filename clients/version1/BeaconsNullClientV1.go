package clients

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

type BeaconsNullClientV1 struct {
}

func NewBeaconsNullClientV1() *BeaconsNullClientV1 {
	return &BeaconsNullClientV1{}
}

func (c *BeaconsNullClientV1) getBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *data1.BeaconV1DataPage, err error) {
	return data1.NewEmptyBeaconV1DataPage(), nil
}

func (c *BeaconsNullClientV1) getBeaconById(correlationId string, beaconId string) (beacon *data1.BeaconV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) getBeaconByUdi(correlationId string, udi string) (beacon *data1.BeaconV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) calculatePosition(correlationId string, siteId string, udis []string) (position *data1.GeoPointV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) createBeacon(correlationId string, beacon *data1.BeaconV1) (res *data1.BeaconV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) updateBeacon(correlationId string, beacon *data1.BeaconV1) (res *data1.BeaconV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) deleteBeaconById(correlationId string, beaconId string) (beacon *data1.BeaconV1, err error) {
	return nil, nil
}
