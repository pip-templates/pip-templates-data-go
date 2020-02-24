package data

type GeoPointV1 struct {
	Type string  `json:"type" bson:"type"`
	Lng  float32 `json:"lng" bson:"lng"`
	Lat  float32 `json:"lat" bson:"lat"`
}
