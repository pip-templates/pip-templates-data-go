package clients

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	grpcclient "github.com/pip-services3-go/pip-services3-grpc-go/clients"
	bdata "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

type BeaconsGrpcClientV1 struct {
	*grpcclient.CommandableGrpcClient
}

func NewBeaconsGrpcClientV1() *BeaconsGrpcClientV1 {
	bhc := BeaconsGrpcClientV1{}
	bhc.CommandableGrpcClient = grpcclient.NewCommandableGrpcClient("v1.beacons")
	return &bhc
}

func (c *BeaconsGrpcClientV1) GetBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error) {

	params := cdata.NewEmptyStringValueMap()
	c.AddFilterParams(params, filter)
	c.AddPagingParams(params, paging)

	calValue, calErr := c.CallCommand(beaconV1DataPageType, "get_beacons", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	page, _ = calValue.(*bdata.BeaconV1DataPage)
	return page, err
}

func (c *BeaconsGrpcClientV1) GetBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {

	params := cdata.NewEmptyStringValueMap()
	params.Put("beacon_id", beaconId)

	calValue, calErr := c.CallCommand(beaconV1Type, "get_beacon_by_id", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*bdata.BeaconV1)
	return beacon, err
}

func (c *BeaconsGrpcClientV1) GetBeaconByUdi(correlationId string, udi string) (beacon *bdata.BeaconV1, err error) {

	params := cdata.NewEmptyStringValueMap()
	params.Put("udi", udi)

	calValue, calErr := c.CallCommand(beaconV1Type, "get_beacon_by_udi", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*bdata.BeaconV1)
	return beacon, err
}

func (c *BeaconsGrpcClientV1) CalculatePosition(correlationId string, siteId string, udis []string) (position *bdata.GeoPointV1, err error) {

	params := make(map[string]interface{})
	params["site_id"] = siteId
	params["udis"] = udis

	calValue, calErr := c.CallCommand(geoPointV1Type, "calculate_position", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	position, _ = calValue.(*bdata.GeoPointV1)
	return position, err
}

func (c *BeaconsGrpcClientV1) CreateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {

	params := make(map[string]interface{})
	params["beacon"] = beacon

	calValue, calErr := c.CallCommand(beaconV1Type, "create_beacon", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	res, _ = calValue.(*bdata.BeaconV1)
	return res, err
}

func (c *BeaconsGrpcClientV1) UpdateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {

	params := make(map[string]interface{})
	params["beacon"] = beacon

	calValue, calErr := c.CallCommand(beaconV1Type, "update_beacon", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	res, _ = calValue.(*bdata.BeaconV1)
	return res, err
}

func (c *BeaconsGrpcClientV1) DeleteBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {

	params := cdata.NewEmptyStringValueMap()
	params.Put("beacon_id", beaconId)

	calValue, calErr := c.CallCommand(beaconV1Type, "delete_beacon_by_id", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*bdata.BeaconV1)
	return beacon, err
}
