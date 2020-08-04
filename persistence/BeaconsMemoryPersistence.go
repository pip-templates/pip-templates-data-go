package persistence

import (
	"reflect"
	"strings"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cpersist "github.com/pip-services3-go/pip-services3-data-go/persistence"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

type BeaconsMemoryPersistence struct {
	cpersist.IdentifiableMemoryPersistence
}

func NewBeaconsMemoryPersistence() *BeaconsMemoryPersistence {
	proto := reflect.TypeOf(&data1.BeaconV1{})
	c := BeaconsMemoryPersistence{}
	c.IdentifiableMemoryPersistence = *cpersist.NewIdentifiableMemoryPersistence(proto)
	//c.MaxPageSize = 1000
	return &c
}

func (c *BeaconsMemoryPersistence) composeFilter(filter *cdata.FilterParams) func(beacon interface{}) bool {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	id := filter.GetAsString("id")
	siteId := filter.GetAsString("site_id")
	label := filter.GetAsString("label")
	udi := filter.GetAsString("udi")
	udis := filter.GetAsString("udis")

	var arrUdis []string = make([]string, 0, 0)
	if udis != "" {
		arrUdis = strings.Split(udis, ",")
	}

	return func(beacon interface{}) bool {
		item, ok := beacon.(data1.BeaconV1)
		if !ok {
			return false
		}
		if id != "" && item.Id != id {
			return false
		}
		if siteId != "" && item.Site_id != siteId {
			return false
		}
		if label != "" && item.Label != label {
			return false
		}
		if udi != "" && item.Udi != udi {
			return false
		}
		if len(arrUdis) > 0 && strings.Index(udis, item.Udi) < 0 {
			return false
		}
		return true
	}
}

func (c *BeaconsMemoryPersistence) Create(correlationId string, item *data1.BeaconV1) (*data1.BeaconV1, error) {
	value, err := c.IdentifiableMemoryPersistence.Create(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	result, _ := value.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsMemoryPersistence) GetListByIds(correlationId string, ids []string) ([]data1.BeaconV1, error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}

	result, err := c.IdentifiableMemoryPersistence.GetListByIds(correlationId, convIds)

	if result == nil || err != nil {
		return nil, err
	}

	items := make([]data1.BeaconV1, len(result))
	for i, v := range result {
		val, _ := v.(data1.BeaconV1)
		items[i] = val
	}
	return items, err
}

func (c *BeaconsMemoryPersistence) GetOneById(correlationId string, id string) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableMemoryPersistence.GetOneById(correlationId, id)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, err
}

func (c *BeaconsMemoryPersistence) Update(correlationId string, item *data1.BeaconV1) (*data1.BeaconV1, error) {
	value, err := c.IdentifiableMemoryPersistence.Update(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	result, _ := value.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsMemoryPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableMemoryPersistence.UpdatePartially(correlationId, id, data)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, nil
}

func (c *BeaconsMemoryPersistence) DeleteById(correlationId string, id string) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableMemoryPersistence.DeleteById(correlationId, id)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, nil
}

func (c *BeaconsMemoryPersistence) DeleteByIds(correlationId string, ids []string) error {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}

	return c.IdentifiableMemoryPersistence.DeleteByIds(correlationId, convIds)
}

func (c *BeaconsMemoryPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (*data1.BeaconV1DataPage, error) {
	tempPage, err := c.IdentifiableMemoryPersistence.GetPageByFilter(correlationId, c.composeFilter(filter), paging, nil, nil)

	if tempPage == nil || err != nil {
		return nil, err
	}

	// Convert to BeaconV1DataPage
	dataLen := len(tempPage.Data) // For full release tempPage and delete this by GC
	beaconData := make([]*data1.BeaconV1, dataLen)
	for i, v := range tempPage.Data {
		beaconData[i] = v.(*data1.BeaconV1)
	}

	page := data1.NewBeaconV1DataPage(tempPage.Total, beaconData)
	return page, nil
}

func (c *BeaconsMemoryPersistence) GetOneByUdi(correlationId string, udi string) (*data1.BeaconV1, error) {
	var item *data1.BeaconV1
	for _, v := range c.Items {
		if buf, ok := v.(data1.BeaconV1); ok {
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
