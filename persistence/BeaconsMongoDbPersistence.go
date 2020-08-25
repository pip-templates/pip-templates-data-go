package persistence

import (
	"reflect"
	"strings"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	persist "github.com/pip-services3-go/pip-services3-mongodb-go/persistence"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BeaconsMongoDbPersistence struct {
	persist.IdentifiableMongoDbPersistence
}

func NewBeaconsMongoDbPersistence() *BeaconsMongoDbPersistence {
	proto := reflect.TypeOf(&data1.BeaconV1{})
	c := BeaconsMongoDbPersistence{}
	c.IdentifiableMongoDbPersistence = *persist.NewIdentifiableMongoDbPersistence(proto, "beacons")
	return &c
}

func (c *BeaconsMongoDbPersistence) composeFilter(filter *cdata.FilterParams) interface{} {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	criteria := make([]bson.M, 0, 0)

	id := filter.GetAsString("id")
	if id != "" {
		criteria = append(criteria, bson.M{"_id": id})
	}

	siteId := filter.GetAsString("site_id")
	if siteId != "" {
		criteria = append(criteria, bson.M{"site_id": siteId})
	}
	label := filter.GetAsString("label")
	if label != "" {
		criteria = append(criteria, bson.M{"label": label})
	}
	udi := filter.GetAsString("udi")
	if udi != "" {
		criteria = append(criteria, bson.M{"udi": udi})
	}

	udis := filter.GetAsString("udis")
	var arrUdis []string = make([]string, 0, 0)
	if udis != "" {
		arrUdis = strings.Split(udis, ",")
		if len(arrUdis) > 1 {
			criteria = append(criteria, bson.M{"udi": bson.D{{"$in", arrUdis}}})
		}
	}
	if len(criteria) > 0 {
		return bson.D{{"$and", criteria}}
	}
	return bson.M{}
}

func (c *BeaconsMongoDbPersistence) Create(correlationId string, item *data1.BeaconV1) (*data1.BeaconV1, error) {
	value, err := c.IdentifiableMongoDbPersistence.Create(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	result, _ := value.(*data1.BeaconV1)
	return result, nil
}

func (c *BeaconsMongoDbPersistence) GetListByIds(correlationId string, ids []string) ([]*data1.BeaconV1, error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}

	result, err := c.IdentifiableMongoDbPersistence.GetListByIds(correlationId, convIds)

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

func (c *BeaconsMongoDbPersistence) GetOneById(correlationId string, id string) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableMongoDbPersistence.GetOneById(correlationId, id)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, nil
}

func (c *BeaconsMongoDbPersistence) Update(correlationId string, item *data1.BeaconV1) (*data1.BeaconV1, error) {
	value, err := c.IdentifiableMongoDbPersistence.Update(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	result, _ := value.(*data1.BeaconV1)
	return result, err
}

func (c *BeaconsMongoDbPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableMongoDbPersistence.UpdatePartially(correlationId, id, data)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, err
}

func (c *BeaconsMongoDbPersistence) DeleteById(correlationId string, id string) (*data1.BeaconV1, error) {
	result, err := c.IdentifiableMongoDbPersistence.DeleteById(correlationId, id)

	if result == nil || err != nil {
		return nil, err
	}

	item, _ := result.(*data1.BeaconV1)
	return item, err
}

func (c *BeaconsMongoDbPersistence) DeleteByIds(correlationId string, ids []string) error {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}

	return c.IdentifiableMongoDbPersistence.DeleteByIds(correlationId, convIds)
}

func (c *BeaconsMongoDbPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (*data1.BeaconV1DataPage, error) {
	tempPage, err := c.IdentifiableMongoDbPersistence.GetPageByFilter(correlationId, c.composeFilter(filter), paging, nil, nil)
	// Todo: Here we shall receive a reference instead of object
	if /*tempPage == nil ||*/ err != nil {
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

func (c *BeaconsMongoDbPersistence) GetOneByUdi(correlationId string, udi string) (*data1.BeaconV1, error) {
	filter := bson.M{"udi": udi}
	p := c.NewObjectByPrototype()
	r := c.Collection.FindOne(c.Connection.Ctx, filter)
	err := r.Decode(p.Interface())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	item := c.ConvertResultToPublic(p, c.Prototype)

	if item == nil {
		return nil, nil //??
	}

	result, _ := item.(*data1.BeaconV1)
	return result, nil
}
