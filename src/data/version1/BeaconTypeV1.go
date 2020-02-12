package data

var BeaconTypeV1 TBeaconTypeV1 = *NewTBeaconTypeV1()

type TBeaconTypeV1 struct {
	Unknown      string
	AltBeacon    string
	iBeacon      string
	EddyStoneUdi string
}

func NewTBeaconTypeV1() *TBeaconTypeV1 {
	bt := TBeaconTypeV1{
		Unknown:      "unknown",
		AltBeacon:    "altbeacon",
		iBeacon:      "ibeacon",
		EddyStoneUdi: "eddystone-udi",
	}
	return &bt
}
