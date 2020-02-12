package persistence

import (
	data "github.com/pip-services-samples/pip-samples-beacons-node/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/src/data/version1"
)

type IBeaconsPersistence interface {
	GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *data.BeaconV1DataPage, err error)

	GetOneById(correlationId string, id string) (res *data.BeaconV1, err error)

	GetOneByUdi(correlationId string, udi string) (res *data.BeaconV1, err error)

	Create(correlationId string, item data.BeaconV1) (res *data.BeaconV1, err error)

	Update(correlationId string, item data.BeaconV1) (res *data.BeaconV1, err error)

	DeleteById(correlationId string, id string) (res *data.BeaconV1, err error)
}
