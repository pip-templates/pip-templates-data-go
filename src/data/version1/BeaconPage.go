package data

type BeaconPage struct {
	Total *int64   `json:"total"`
	Data  []Beacon `json:"data"`
}

func NewEmptyBeaconPage() *BeaconPage {
	return &BeaconPage{}
}

func NewBeaconPage(total *int64, data []Dummy) *BeaconPage {
	return &BeaconPage{Total: total, Data: data}
}
