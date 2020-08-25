package persistence

import (
	"encoding/json"
	"reflect"
	"strings"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	persist "github.com/pip-services3-go/pip-services3-couchbase-go/persistence"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	"gopkg.in/couchbase/gocb.v1"
)

type BeaconsCouchbasePersistence struct {
	persist.IdentifiableCouchbasePersistence
}

func NewBeaconsCouchbasePersistence() *BeaconsCouchbasePersistence {
	proto := reflect.TypeOf(&data1.BeaconV1{})
	c := BeaconsCouchbasePersistence{}
	c.IdentifiableCouchbasePersistence = *persist.NewIdentifiableCouchbasePersistence(proto, "beaconBucket", "beacons")
	return &c
}

func (c *BeaconsCouchbasePersistence) composeFilter(filter *cdata.FilterParams) string {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	criteria := make([]string, 0, 0)

	id := filter.GetAsString("id")
	if id != "" {
		criteria = append(criteria, " id='"+id+"'")
	}

	siteId := filter.GetAsString("site_id")
	if siteId != "" {
		criteria = append(criteria, " site_id='"+siteId+"'")
	}
	label := filter.GetAsString("label")
	if label != "" {
		criteria = append(criteria, " label='"+label+"'")
	}
	udi := filter.GetAsString("udi")
	if udi != "" {
		criteria = append(criteria, " udi='"+udi+"'")
	}

	udis := filter.GetAsString("udis")
	var arrUdis []string = make([]string, 0, 0)
	if udis != "" {
		arrUdis = strings.Split(udis, ",")
		if len(arrUdis) > 1 {
			criteria = append(criteria, " udi IN ['"+strings.Join(arrUdis, "','")+"'] ")
		}
	}
	if len(criteria) > 1 {
		return strings.Join(criteria, " AND ")
	}
	if len(criteria) == 1 {
		return criteria[0]
	}
	return ""
}

func (c *BeaconsCouchbasePersistence) Create(correlationId string, item *data1.BeaconV1) (*data1.BeaconV1, error) {
	value, err := c.IdentifiableCouchbasePersistence.Create(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	result, _ := value.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsCouchbasePersistence) GetListByIds(correlationId string, ids []string) ([]*data1.BeaconV1, error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}

	result, err := c.IdentifiableCouchbasePersistence.GetListByIds(correlationId, convIds)

	if result == nil || err != nil {
		return nil, err
	}

	items := make([]*data1.BeaconV1, len(result))
	for i, v := range result {
		val, _ := v.(*data1.BeaconV1)
		items[i] = val
	}
	return items, nil
}

func (c *BeaconsCouchbasePersistence) GetOneById(correlationId string, id string) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableCouchbasePersistence.GetOneById(correlationId, id)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, nil
}

func (c *BeaconsCouchbasePersistence) Update(correlationId string, item *data1.BeaconV1) (*data1.BeaconV1, error) {
	value, err := c.IdentifiableCouchbasePersistence.Update(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	result, _ := value.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsCouchbasePersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableCouchbasePersistence.UpdatePartially(correlationId, id, data)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, err
}

func (c *BeaconsCouchbasePersistence) DeleteById(correlationId string, id string) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableCouchbasePersistence.DeleteById(correlationId, id)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, err
}

func (c *BeaconsCouchbasePersistence) DeleteByIds(correlationId string, ids []string) error {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}

	return c.IdentifiableCouchbasePersistence.DeleteByIds(correlationId, convIds)
}

func (c *BeaconsCouchbasePersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (*data1.BeaconV1DataPage, error) {
	tempPage, err := c.IdentifiableCouchbasePersistence.GetPageByFilter(correlationId, c.composeFilter(filter), paging, "", "")

	if tempPage == nil || err != nil {
		return nil, err
	}

	// Convert to BeaconV1DataPage
	dataLen := int64(len(tempPage.Data)) // For full release tempPage and delete this by GC
	beaconData := make([]*data1.BeaconV1, dataLen)
	for i, v := range tempPage.Data {
		beaconData[i] = v.(*data1.BeaconV1)
	}
	page := data1.NewBeaconV1DataPage(tempPage.Total, beaconData)
	return page, nil
}

func (c *BeaconsCouchbasePersistence) GetOneByUdi(correlationId string, udi string) (*data1.BeaconV1, error) {
	if udi == "" {
		return nil, nil
	}

	statement := "SELECT * FROM `" + c.BucketName + "`" + " WHERE udi = '" + udi + "' LIMIT 1"

	query := gocb.NewN1qlQuery(statement)
	queryRes, queryErr := c.Bucket.ExecuteN1qlQuery(query, nil)
	if queryErr != nil {
		return nil, queryErr
	}
	buf := make(map[string]interface{})
	queryRes.Next(&buf)
	docPointer := c.GetProtoPtr()
	// select *
	jsonBuf, _ := json.Marshal(buf[c.BucketName])

	json.Unmarshal(jsonBuf, docPointer.Interface())
	item := c.GetConvResult(docPointer)

	if item == nil {
		return nil, nil //??
	}

	//--------------------------------------------
	result, _ := item.(*data1.BeaconV1)
	return result, nil
}
