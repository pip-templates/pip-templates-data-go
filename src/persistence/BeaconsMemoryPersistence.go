package persistence

import (
	"reflect"
	"strings"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cmperist "github.com/pip-services3-go/pip-services3-data-go/persistence"
	bdata "github.com/pip-templates/pip-templates-microservice-go/src/data/version1"
)

type BeaconsMemoryPersistence struct {
	cmperist.IdentifiableMemoryPersistence
}

func NewBeaconsMemoryPersistence() *BeaconsMemoryPersistence {
	proto := reflect.TypeOf(&bdata.BeaconV1{})
	bmp := BeaconsMemoryPersistence{}
	bmp.IdentifiableMemoryPersistence = *cmperist.NewIdentifiableMemoryPersistence(proto)
	//bmp.MaxPageSize = 1000
	return &bmp
}

func (c *BeaconsMemoryPersistence) composeFilter(filter *cdata.FilterParams) func(beacon interface{}) bool {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	id := filter.GetAsNullableString("id")
	siteId := filter.GetAsNullableString("site_id")
	label := filter.GetAsNullableString("label")
	udi := filter.GetAsNullableString("udi")
	udis := filter.GetAsNullableString("udis")

	var arrUdis []string = make([]string, 0, 0)
	if udis != nil {
		arrUdis = strings.Split(*udis, ",")
	}

	return func(beacon interface{}) bool {
		item, ok := beacon.(bdata.BeaconV1)
		if !ok {
			return false
		}
		if id != nil && item.Id != *id {
			return false
		}
		if siteId != nil && item.Site_id != *siteId {
			return false
		}
		if label != nil && item.Label != *label {
			return false
		}
		if udi != nil && item.Udi != *udi {
			return false
		}
		if len(arrUdis) > 0 && strings.Index(*udis, item.Udi) < 0 {
			return false
		}
		return true
	}
}

func (c *BeaconsMemoryPersistence) Create(correlationId string, item bdata.BeaconV1) (result *bdata.BeaconV1, err error) {
	value, err := c.IdentifiableMemoryPersistence.Create(correlationId, item)
	if value != nil {
		val, _ := value.(*bdata.BeaconV1)
		result = val
	}
	return result, err
}

func (c *BeaconsMemoryPersistence) GetListByIds(correlationId string, ids []string) (items []bdata.BeaconV1, err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	result, err := c.IdentifiableMemoryPersistence.GetListByIds(correlationId, convIds)
	items = make([]bdata.BeaconV1, len(result))
	for i, v := range result {
		val, _ := v.(bdata.BeaconV1)
		items[i] = val
	}
	return items, err
}

func (c *BeaconsMemoryPersistence) GetOneById(correlationId string, id string) (item *bdata.BeaconV1, err error) {
	result, err := c.IdentifiableMemoryPersistence.GetOneById(correlationId, id)
	if result != nil {
		val, _ := result.(*bdata.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsMemoryPersistence) Update(correlationId string, item bdata.BeaconV1) (result *bdata.BeaconV1, err error) {
	value, err := c.IdentifiableMemoryPersistence.Update(correlationId, item)
	if value != nil {
		val, _ := value.(*bdata.BeaconV1)
		result = val
	}
	return result, err
}

func (c *BeaconsMemoryPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (item *bdata.BeaconV1, err error) {
	result, err := c.IdentifiableMemoryPersistence.UpdatePartially(correlationId, id, data)

	if result != nil {
		val, _ := result.(*bdata.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsMemoryPersistence) DeleteById(correlationId string, id string) (item *bdata.BeaconV1, err error) {
	result, err := c.IdentifiableMemoryPersistence.DeleteById(correlationId, id)
	if result != nil {
		val, _ := result.(*bdata.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsMemoryPersistence) DeleteByIds(correlationId string, ids []string) (err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	return c.IdentifiableMemoryPersistence.DeleteByIds(correlationId, convIds)
}

func (c *BeaconsMemoryPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error) {
	tempPage, resErr := c.IdentifiableMemoryPersistence.GetPageByFilter(correlationId, c.composeFilter(filter), paging, nil, nil)
	if resErr != nil {
		return nil, resErr
	}
	// Convert to BeaconV1DataPage
	dataLen := int64(len(tempPage.Data)) // For full release tempPage and delete this by GC
	beaconData := make([]bdata.BeaconV1, dataLen)
	for i, v := range tempPage.Data {
		beaconData[i] = v.(bdata.BeaconV1)
	}
	page = bdata.NewBeaconV1DataPage(&dataLen, beaconData)
	return page, nil
}

func (c *BeaconsMemoryPersistence) GetOneByUdi(correlationId string, udi string) (res *bdata.BeaconV1, err error) {

	var item *bdata.BeaconV1
	for _, v := range c.Items {
		if buf, ok := v.(bdata.BeaconV1); ok {
			if buf.Udi == udi {
				item = &buf
				break
			}
		}
	}

	if item != nil {
		c.Logger.Trace(correlationId, "Found beacon by %s", udi)
	} else {
		c.Logger.Trace(correlationId, "Cannot find beacon by %s", udi)
	}

	return item, nil
}
