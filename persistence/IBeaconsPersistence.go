package persistence

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

type IBeaconsPersistence interface {
	GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *data1.BeaconV1DataPage, err error)

	GetOneById(correlationId string, id string) (res *data1.BeaconV1, err error)

	GetOneByUdi(correlationId string, udi string) (res *data1.BeaconV1, err error)

	Create(correlationId string, item *data1.BeaconV1) (res *data1.BeaconV1, err error)

	Update(correlationId string, item *data1.BeaconV1) (res *data1.BeaconV1, err error)

	DeleteById(correlationId string, id string) (res *data1.BeaconV1, err error)
}
