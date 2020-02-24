package clients

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	bdata "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

type BeaconsNullClientV1 struct {
}

func NewBeaconsNullClientV1() *BeaconsNullClientV1 {
	return &BeaconsNullClientV1{}
}

func (c *BeaconsNullClientV1) getBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error) {
	return bdata.NewEmptyBeaconV1DataPage(), nil
}

func (c *BeaconsNullClientV1) getBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) getBeaconByUdi(correlationId string, udi string) (beacon *bdata.BeaconV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) calculatePosition(correlationId string, siteId string, udis []string) (position *bdata.GeoPointV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) createBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) updateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {
	return nil, nil
}

func (c *BeaconsNullClientV1) deleteBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {
	return nil, nil
}
