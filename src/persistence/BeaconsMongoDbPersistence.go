package persistence

import (
	"reflect"
	"strings"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	mngpersist "github.com/pip-services3-go/pip-services3-mongodb-go/persistence"
	bdata "github.com/pip-templates/pip-templates-microservice-go/src/data/version1"
	data "github.com/pip-templates/pip-templates-microservice-go/src/data/version1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BeaconsMongoDbPersistence struct {
	mngpersist.IdentifiableMongoDbPersistence
}

func NewBeaconsMongoDbPersistence() *BeaconsMongoDbPersistence {
	proto := reflect.TypeOf(&data.BeaconV1{})
	bmp := BeaconsMongoDbPersistence{}
	bmp.IdentifiableMongoDbPersistence = *mngpersist.NewIdentifiableMongoDbPersistence(proto, "beacons")
	//bmp.MaxPageSize = 1000;
	return &bmp
}

func (c *BeaconsMongoDbPersistence) composeFilter(filter *cdata.FilterParams) interface{} {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	criteria := make([]bson.M, 0, 0)

	id := filter.GetAsString("id")
	if id != "" {
		//criteria["_id"] = id
		criteria = append(criteria, bson.M{"_id": id})
	}

	siteId := filter.GetAsString("site_id")
	if siteId != "" {
		//criteria["site_id"] = siteId
		criteria = append(criteria, bson.M{"site_id": siteId})
	}
	label := filter.GetAsString("label")
	if label != "" {
		//criteria["label"] = label
		criteria = append(criteria, bson.M{"label": label})
	}
	udi := filter.GetAsString("udi")
	if udi != "" {
		//criteria["udi"] = udi
		criteria = append(criteria, bson.M{"udi": udi})
	}

	udis := filter.GetAsString("udis")
	var arrUdis []string = make([]string, 0, 0)
	if udis != "" {
		arrUdis = strings.Split(udis, ",")
		if len(arrUdis) > 1 {
			//criteria["udi"] = bson.D{{"$in", arrUdis}}
			criteria = append(criteria, bson.M{"udi": bson.D{{"$in", arrUdis}}})
		}
	}
	if len(criteria) > 0 {
		return bson.D{{"$and", criteria}}
	}
	return bson.M{}
}

func (c *BeaconsMongoDbPersistence) Create(correlationId string, item bdata.BeaconV1) (result *bdata.BeaconV1, err error) {
	value, err := c.IdentifiableMongoDbPersistence.Create(correlationId, item)

	if value != nil {
		val, _ := value.(*bdata.BeaconV1)
		result = val
	}
	return result, err
}

func (c *BeaconsMongoDbPersistence) GetListByIds(correlationId string, ids []string) (items []*bdata.BeaconV1, err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	result, err := c.IdentifiableMongoDbPersistence.GetListByIds(correlationId, convIds)
	items = make([]*bdata.BeaconV1, len(result))
	for i, v := range result {
		val, _ := v.(*bdata.BeaconV1)
		items[i] = val
	}
	return items, err
}

func (c *BeaconsMongoDbPersistence) GetOneById(correlationId string, id string) (item *bdata.BeaconV1, err error) {
	result, err := c.IdentifiableMongoDbPersistence.GetOneById(correlationId, id)
	if result != nil {
		val, _ := result.(*bdata.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsMongoDbPersistence) Update(correlationId string, item bdata.BeaconV1) (result *bdata.BeaconV1, err error) {
	value, err := c.IdentifiableMongoDbPersistence.Update(correlationId, item)
	if value != nil {
		val, _ := value.(*bdata.BeaconV1)
		result = val
	}
	return result, err
}

func (c *BeaconsMongoDbPersistence) UpdatePartially(correlationId string, id string, data *cdata.AnyValueMap) (item *bdata.BeaconV1, err error) {
	result, err := c.IdentifiableMongoDbPersistence.UpdatePartially(correlationId, id, data)

	if result != nil {
		val, _ := result.(*bdata.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsMongoDbPersistence) DeleteById(correlationId string, id string) (item *bdata.BeaconV1, err error) {
	result, err := c.IdentifiableMongoDbPersistence.DeleteById(correlationId, id)
	if result != nil {
		val, _ := result.(*bdata.BeaconV1)
		item = val
	}
	return item, err
}

func (c *BeaconsMongoDbPersistence) DeleteByIds(correlationId string, ids []string) (err error) {
	convIds := make([]interface{}, len(ids))
	for i, v := range ids {
		convIds[i] = v
	}
	return c.IdentifiableMongoDbPersistence.DeleteByIds(correlationId, convIds)
}

func (c *BeaconsMongoDbPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error) {
	tempPage, resErr := c.IdentifiableMongoDbPersistence.GetPageByFilter(correlationId, c.composeFilter(filter), paging, nil, nil)
	if resErr != nil {
		return nil, resErr
	}
	// Convert to BeaconV1DataPage
	dataLen := int64(len(tempPage.Data)) // For full release tempPage and delete this by GC
	beaconData := make([]*bdata.BeaconV1, dataLen)
	for i, v := range tempPage.Data {
		beaconData[i] = v.(*bdata.BeaconV1)
	}
	page = bdata.NewBeaconV1DataPage(&dataLen, beaconData)
	return page, nil
}

func (c *BeaconsMongoDbPersistence) GetOneByUdi(correlationId string, udi string) (result *bdata.BeaconV1, err error) {

	filter := bson.M{"udi": udi}
	docPointer := c.GetProtoPtr()
	foRes := c.Collection.FindOne(c.Connection.Ctx, filter)
	ferr := foRes.Decode(docPointer.Interface())
	if ferr != nil {
		if ferr == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, ferr
	}
	item := c.GetConvResult(docPointer, c.Prototype)

	if item != nil {
		val, _ := item.(*bdata.BeaconV1)
		result = val
	}

	return result, nil

}
