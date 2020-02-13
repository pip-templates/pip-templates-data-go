package data

//implements IStringIdentifiable
type BeaconV1 struct {
	Id      string     `json:"id" bson:"id"`
	Site_id string     `json:"site_id" bson:"site_id"`
	Type    string     `json:"type" bson:"type"`
	Udi     string     `json:"udi" bson:"udi"`
	Label   string     `json:"Label" bson:"Label"`
	Center  GeoPointV1 `json:"center" bson:"center"` // GeoJson
	Radius  float32    `json:"radius" bson:"radius"`
}
