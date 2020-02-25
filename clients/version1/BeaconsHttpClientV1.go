package clients

import (
	"reflect"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	rpcclient "github.com/pip-services3-go/pip-services3-rpc-go/clients"
	bdata "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

var (
	beaconV1DataPageType = reflect.TypeOf(&bdata.BeaconV1DataPage{})
	beaconV1Type         = reflect.TypeOf(&bdata.BeaconV1{})
	geoPointV1Type       = reflect.TypeOf(&bdata.GeoPointV1{})
)

type BeaconsHttpClientV1 struct {
	rpcclient.CommandableHttpClient
}

func NewBeaconsHttpClientV1() *BeaconsHttpClientV1 {
	bhc := BeaconsHttpClientV1{}
	bhc.CommandableHttpClient = *rpcclient.NewCommandableHttpClient("v1/beacons")
	return &bhc
}

func (c *BeaconsHttpClientV1) GetBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error) {
	params := cdata.NewEmptyStringValueMap()
	c.AddFilterParams(params, filter)
	c.AddPagingParams(params, paging)
	calValue, calErr := c.CallCommand(beaconV1DataPageType, "get_beacons", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	page, _ = calValue.(*bdata.BeaconV1DataPage)
	return page, err
}

func (c *BeaconsHttpClientV1) GetBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {
	params := cdata.NewEmptyStringValueMap()
	params.Put("beacon_id", beaconId)
	calValue, calErr := c.CallCommand(beaconV1Type, "get_beacon_by_id", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*bdata.BeaconV1)
	return beacon, err
}

func (c *BeaconsHttpClientV1) GetBeaconByUdi(correlationId string, udi string) (beacon *bdata.BeaconV1, err error) {
	params := cdata.NewEmptyStringValueMap()
	params.Put("udi", udi)
	calValue, calErr := c.CallCommand(beaconV1Type, "get_beacon_by_udi", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*bdata.BeaconV1)
	return beacon, err
}

func (c *BeaconsHttpClientV1) CalculatePosition(correlationId string, siteId string, udis []string) (position *bdata.GeoPointV1, err error) {

	params := make(map[string]interface{})
	params["site_id"] = siteId
	params["udis"] = udis

	calValue, calErr := c.CallCommand(geoPointV1Type, "calculate_position", correlationId, nil, params)
	if calErr != nil {
		return nil, calErr
	}
	position, _ = calValue.(*bdata.GeoPointV1)
	return position, err
}

func (c *BeaconsHttpClientV1) CreateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {

	params := make(map[string]interface{})
	params["beacon"] = beacon

	calValue, calErr := c.CallCommand(beaconV1Type, "create_beacon", correlationId, nil, params)
	if calErr != nil {
		return nil, calErr
	}
	res, _ = calValue.(*bdata.BeaconV1)
	return res, err
}

func (c *BeaconsHttpClientV1) UpdateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {

	params := make(map[string]interface{})
	params["beacon"] = beacon

	calValue, calErr := c.CallCommand(beaconV1Type, "update_beacon", correlationId, nil, params)
	if calErr != nil {
		return nil, calErr
	}
	res, _ = calValue.(*bdata.BeaconV1)
	return res, err
}

func (c *BeaconsHttpClientV1) DeleteBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {

	params := cdata.NewEmptyStringValueMap()
	params.Put("beacon_id", beaconId)
	calValue, calErr := c.CallCommand(beaconV1Type, "delete_beacon_by_id", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	beacon, _ = calValue.(*bdata.BeaconV1)
	return beacon, err
}
