package data

import (
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cvalid "github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type BeaconV1Schema struct {
	cvalid.ObjectSchema
}

func NewBeaconV1Schema() *BeaconV1Schema {
	bs := BeaconV1Schema{}
	bs.ObjectSchema = *cvalid.NewObjectSchema()

	bs.WithOptionalProperty("id", cconv.String)
	bs.WithRequiredProperty("site_id", cconv.String)
	bs.WithOptionalProperty("type", cconv.String)
	bs.WithRequiredProperty("udi", cconv.String)
	bs.WithOptionalProperty("label", cconv.String)
	bs.WithOptionalProperty("center", cconv.Map)
	bs.WithOptionalProperty("radius", cconv.Double)
	return &bs
}
