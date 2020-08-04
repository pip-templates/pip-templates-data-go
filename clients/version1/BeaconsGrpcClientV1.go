package clients

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	clients "github.com/pip-services3-go/pip-services3-grpc-go/clients"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

type BeaconsGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewBeaconsGrpcClientV1() *BeaconsGrpcClientV1 {
	c := BeaconsGrpcClientV1{}
	c.CommandableGrpcClient = clients.NewCommandableGrpcClient("v1.beacons")
	return &c
}

func (c *BeaconsGrpcClientV1) GetBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *data1.BeaconV1DataPage, err error) {
	params := cdata.NewEmptyStringValueMap()
	c.AddFilterParams(params, filter)
	c.AddPagingParams(params, paging)

	calValue, calErr := c.CallCommand(beaconV1DataPageType, "get_beacons", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	page, _ = calValue.(*data1.BeaconV1DataPage)
	return page, err
}

func (c *BeaconsGrpcClientV1) GetBeaconById(correlationId string, beaconId string) (beacon *data1.BeaconV1, err error) {
	params := cdata.NewStringValueMapFromTuples(
		"beacon_id", beaconId,
	)

	calValue, calErr := c.CallCommand(beaconV1Type, "get_beacon_by_id", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*data1.BeaconV1)
	return beacon, nil
}

func (c *BeaconsGrpcClientV1) GetBeaconByUdi(correlationId string, udi string) (beacon *data1.BeaconV1, err error) {
	params := cdata.NewStringValueMapFromTuples(
		"udi", udi,
	)

	calValue, calErr := c.CallCommand(beaconV1Type, "get_beacon_by_udi", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*data1.BeaconV1)
	return beacon, nil
}

func (c *BeaconsGrpcClientV1) CalculatePosition(correlationId string, siteId string, udis []string) (position *data1.GeoPointV1, err error) {
	params := cdata.NewStringValueMapFromTuples(
		"site_id", siteId,
		"udis", udis,
	)

	calValue, calErr := c.CallCommand(geoPointV1Type, "calculate_position", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	position, _ = calValue.(*data1.GeoPointV1)
	return position, nil
}

func (c *BeaconsGrpcClientV1) CreateBeacon(correlationId string, beacon *data1.BeaconV1) (res *data1.BeaconV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"beacon", *beacon,
	)

	calValue, calErr := c.CallCommand(beaconV1Type, "create_beacon", correlationId, params.Value())
	if calErr != nil {
		return nil, calErr
	}
	res, _ = calValue.(*data1.BeaconV1)
	return res, nil
}

func (c *BeaconsGrpcClientV1) UpdateBeacon(correlationId string, beacon *data1.BeaconV1) (res *data1.BeaconV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"beacon", *beacon,
	)

	calValue, calErr := c.CallCommand(beaconV1Type, "update_beacon", correlationId, params.Value())
	if calErr != nil {
		return nil, calErr
	}
	res, _ = calValue.(*data1.BeaconV1)
	return res, nil
}

func (c *BeaconsGrpcClientV1) DeleteBeaconById(correlationId string, beaconId string) (beacon *data1.BeaconV1, err error) {
	params := cdata.NewStringValueMapFromTuples(
		"beacon_id", beaconId,
	)

	calValue, calErr := c.CallCommand(beaconV1Type, "delete_beacon_by_id", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*data1.BeaconV1)
	return beacon, nil
}
