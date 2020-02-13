package clients

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	bdata "github.com/pip-templates/pip-templates-microservice-go/src/data/version1"
)

type IBeaconsClientV1 interface {
	GetBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error)

	GetBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error)

	GetBeaconByUdi(correlationId string, udi string) (beacon *bdata.BeaconV1, err error)

	CalculatePosition(correlationId string, siteId string, udis []string) (position *bdata.GeoPointV1, err error)

	CreateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error)

	UpdateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error)

	DeleteBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error)
}
