package persistence

import (
	"context"
	"reflect"
	"strings"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	ppersist "github.com/pip-services3-go/pip-services3-postgres-go/persistence"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
)

type BeaconsJsonPostgresPersistence struct {
	ppersist.IdentifiableJsonPostgresPersistence
}

func NewBeaconsJsonPostgresPersistence() *BeaconsJsonPostgresPersistence {
	proto := reflect.TypeOf(&data1.BeaconV1{})
	c := &BeaconsJsonPostgresPersistence{
		IdentifiableJsonPostgresPersistence: *ppersist.NewIdentifiableJsonPostgresPersistence(proto, "beacons_json"),
	}

	c.EnsureTable("", "")
	c.EnsureIndex("beacons_json_key", map[string]string{"(data->'id')": "1"}, map[string]string{"unique": "true"})
	return c
}

func (c *BeaconsJsonPostgresPersistence) composeFilter(filter *cdata.FilterParams) interface{} {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	criteria := cdata.NewEmptyStringValueMap()
	result := strings.Builder{}

	id := filter.GetAsString("id")
	if id != "" {
		criteria.SetAsObject("\"id\"=", id)
	}

	siteId := filter.GetAsString("site_id")
	if siteId != "" {
		criteria.SetAsObject("data->>'site_id'=", siteId)
	}
	label := filter.GetAsString("label")
	if label != "" {
		criteria.SetAsObject("data->>'label'=", label)
	}
	udi := filter.GetAsString("udi")
	if udi != "" {
		criteria.SetAsObject("data->>'udi'=", udi)
	}

	if criteria.Len() > 0 {

		for item, val := range criteria.Value() {
			if result.String() != "" {
				result.WriteString(" AND ")
			}
			result.WriteString(item + "'" + val + "'")
		}
	}

	udis := filter.GetAsString("udis")
	var arrUdis []string = make([]string, 0, 0)
	if udis != "" {
		arrUdis = strings.Split(udis, ",")
		if len(arrUdis) > 1 {
			values := strings.Join(arrUdis, "','")
			values = "('" + values + "')"
			if result.String() != "" {
				result.WriteString(" AND ")
			}
			result.WriteString("data->>'udi' in " + values)
		}
	}
	return result.String()
}

func (c *BeaconsJsonPostgresPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *data1.BeaconV1DataPage, err error) {

	tempPage, err := c.IdentifiablePostgresPersistence.GetPageByFilter(correlationId,
		c.composeFilter(filter), paging, nil, nil)
	// Convert to BeaconsPage
	dataLen := int64(len(tempPage.Data)) // For full release tempPage and delete this by GC
	data := make([]*data1.BeaconV1, dataLen)
	for i, v := range tempPage.Data {
		data[i] = v.(*data1.BeaconV1)
	}
	page = data1.NewBeaconV1DataPage(&dataLen, data)
	return page, err
}

func (c *BeaconsJsonPostgresPersistence) Create(correlationId string, item *data1.BeaconV1) (result *data1.BeaconV1, err error) {
	value, err := c.IdentifiablePostgresPersistence.Create(correlationId, item)

	if value != nil {
		val, _ := value.(*data1.BeaconV1)
		result = val
	}
	return result, err
}

func (c *BeaconsJsonPostgresPersistence) GetListByIds(correlationId string, ids []string) (items []*data1.BeaconV1, err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	result, err := c.IdentifiablePostgresPersistence.GetListByIds(correlationId, convIds)
	items = make([]*data1.BeaconV1, len(result))
	for i, v := range result {
		val, _ := v.(*data1.BeaconV1)
		items[i] = val
	}
	return items, err
}

func (c *BeaconsJsonPostgresPersistence) GetOneById(correlationId string, id string) (item *data1.BeaconV1, err error) {
	result, err := c.IdentifiablePostgresPersistence.GetOneById(correlationId, id)
	if result != nil {
		val, _ := result.(*data1.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsJsonPostgresPersistence) Update(correlationId string, item *data1.BeaconV1) (result *data1.BeaconV1, err error) {
	value, err := c.IdentifiablePostgresPersistence.Update(correlationId, item)
	if value != nil {
		val, _ := value.(*data1.BeaconV1)
		result = val
	}
	return result, err
}

func (c *BeaconsJsonPostgresPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (item *data1.BeaconV1, err error) {
	// In json persistence this method must call from IdentifiableJsonPostgresPersistence
	result, err := c.IdentifiableJsonPostgresPersistence.UpdatePartially(correlationId, id, data)

	if result != nil {
		val, _ := result.(*data1.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsJsonPostgresPersistence) DeleteById(correlationId string, id string) (item *data1.BeaconV1, err error) {
	result, err := c.IdentifiablePostgresPersistence.DeleteById(correlationId, id)
	if result != nil {
		val, _ := result.(*data1.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsJsonPostgresPersistence) DeleteByIds(correlationId string, ids []string) (err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	return c.IdentifiablePostgresPersistence.DeleteByIds(correlationId, convIds)
}

func (c *BeaconsJsonPostgresPersistence) GetOneByUdi(correlationId string, udi string) (*data1.BeaconV1, error) {

	query := "SELECT * FROM " + c.QuoteIdentifier(c.TableName) + " WHERE data->>'udi'=$1 LIMIT 1"

	qResult, qErr := c.Client.Query(context.TODO(), query, udi)
	if qErr != nil {
		return nil, qErr
	}
	defer qResult.Close()
	if qResult.Next() {
		rows, vErr := qResult.Values()
		if vErr == nil && len(rows) > 0 {
			result := c.ConvertFromRows(qResult.FieldDescriptions(), rows)
			if result == nil {
				c.Logger.Trace(correlationId, "Nothing found from %s with udi = %s", c.TableName, udi)
			} else {
				c.Logger.Trace(correlationId, "Retrieved from %s with udi = %s", c.TableName, udi)
				val, _ := result.(*data1.BeaconV1)
				return val, nil
			}
		}
		return nil, vErr
	}

	c.Logger.Trace(correlationId, "Nothing found from %s with id = %s", c.TableName, udi)
	return nil, qResult.Err()
}
