package clients

import (
	"encoding/json"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	rpcclient "github.com/pip-services3-go/pip-services3-rpc-go/clients"
	bdata "github.com/pip-templates/pip-templates-microservice-go/src/data/version1"
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
	calValue, calErr := c.CallCommand("get_beacons", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	var data bdata.BeaconV1DataPage
	convErr := json.Unmarshal(calValue.([]byte), &data)
	if convErr != nil {
		return nil, convErr
	}
	return &data, nil
}

func (c *BeaconsHttpClientV1) GetBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {
	params := cdata.NewEmptyStringValueMap()
	params.Put("beacon_id", beaconId)
	calValue, calErr := c.CallCommand("get_beacon_by_id", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	var data bdata.BeaconV1
	convErr := json.Unmarshal(calValue.([]byte), &data)
	if convErr != nil {
		return nil, convErr
	}
	return &data, nil
}

func (c *BeaconsHttpClientV1) GetBeaconByUdi(correlationId string, udi string) (beacon *bdata.BeaconV1, err error) {
	params := cdata.NewEmptyStringValueMap()
	params.Put("udi", udi)
	calValue, calErr := c.CallCommand("get_beacon_by_udi", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	var data bdata.BeaconV1
	convErr := json.Unmarshal(calValue.([]byte), &data)
	if convErr != nil {
		return nil, convErr
	}
	return &data, nil
}

func (c *BeaconsHttpClientV1) CalculatePosition(correlationId string, siteId string, udis []string) (position *bdata.GeoPointV1, err error) {
	params := cdata.NewEmptyStringValueMap()
	params.Put("site_id", siteId)
	params.Put("udis", udis)
	calValue, calErr := c.CallCommand("calculate_position", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	var data bdata.GeoPointV1
	convErr := json.Unmarshal(calValue.([]byte), &data)
	if convErr != nil {
		return nil, convErr
	}
	return &data, nil
}

func (c *BeaconsHttpClientV1) CreateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {
	params := cdata.NewEmptyStringValueMap()
	params.Put("beacon", beacon)
	calValue, calErr := c.CallCommand("create_beacon", correlationId, nil, params)
	if calErr != nil {
		return nil, calErr
	}
	var data bdata.BeaconV1
	convErr := json.Unmarshal(calValue.([]byte), &data)
	if convErr != nil {
		return nil, convErr
	}
	return &data, nil
}

func (c *BeaconsHttpClientV1) UpdateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error) {
	params := cdata.NewEmptyStringValueMap()
	params.Put("beacon", beacon)

	calValue, calErr := c.CallCommand("update_beacon", correlationId, nil, params)
	if calErr != nil {
		return nil, calErr
	}
	var data bdata.BeaconV1
	convErr := json.Unmarshal(calValue.([]byte), &data)
	if convErr != nil {
		return nil, convErr
	}
	return &data, nil
}

func (c *BeaconsHttpClientV1) DeleteBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error) {
	params := cdata.NewEmptyStringValueMap()
	params.Put("beacon_id", beaconId)
	calValue, calErr := c.CallCommand("delete_beacon_by_id", correlationId, params, nil)
	if calErr != nil {
		return nil, calErr
	}
	var data bdata.BeaconV1
	convErr := json.Unmarshal(calValue.([]byte), &data)
	if convErr != nil {
		return nil, convErr
	}
	return &data, nil
}
