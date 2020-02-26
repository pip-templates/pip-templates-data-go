package test_services

import (
	"context"
	"encoding/json"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cmdproto "github.com/pip-services3-go/pip-services3-grpc-go/protos"
	bdata "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	blogic "github.com/pip-templates/pip-templates-microservice-go/logic"
	bpersist "github.com/pip-templates/pip-templates-microservice-go/persistence"
	bservices "github.com/pip-templates/pip-templates-microservice-go/services/version1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func TestBeaconsGrpcServiceV1(t *testing.T) {

	var persistence *bpersist.BeaconsMemoryPersistence
	var controller *blogic.BeaconsController
	var service *bservices.BeaconsGrpcServiceV1
	var client cmdproto.CommandableClient

	persistence = bpersist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = blogic.NewBeaconsController()
	controller.Configure(cconf.NewEmptyConfigParams())
	service = bservices.NewBeaconsGrpcServiceV1()
	service.Configure(cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3002",
		"connection.host", "localhost",
	))

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("beacons", "service", "grpc", "default", "1.0"), service,
	)

	controller.SetReferences(references)
	service.SetReferences(references)

	opnErr := persistence.Open("")
	if opnErr != nil {
		panic("Can't open persistence")
	}
	srvOpnErr := service.Open("")
	assert.Nil(t, srvOpnErr)

	defer service.Close("")
	defer persistence.Close("")

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("localhost:3002", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client = cmdproto.NewCommandableClient(conn)

	var beacon1 bdata.BeaconV1
	// Create the first beacon
	requestParams := make(map[string]interface{})
	requestParams["beacon"] = Beacon1
	jsonBuf, _ := json.Marshal(requestParams)

	request := cmdproto.InvokeRequest{}
	request.Method = "v1.beacons.create_beacon"
	request.ArgsEmpty = false
	request.ArgsJson = string(jsonBuf)
	response, err := client.Invoke(context.TODO(), &request)
	assert.Nil(t, err)
	var beacon bdata.BeaconV1
	jsonErr := json.Unmarshal([]byte(response.ResultJson), &beacon)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon1.Udi, beacon.Udi)
	assert.Equal(t, Beacon1.Site_id, beacon.Site_id)
	assert.Equal(t, Beacon1.Type, beacon.Type)
	assert.Equal(t, Beacon1.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Create the second beacon
	requestParams = make(map[string]interface{})
	requestParams["beacon"] = Beacon2
	jsonBuf, _ = json.Marshal(requestParams)

	request.Method = "v1.beacons.create_beacon"
	request.ArgsEmpty = false
	request.ArgsJson = string(jsonBuf)
	response, err = client.Invoke(context.TODO(), &request)
	assert.Nil(t, err)

	jsonErr = json.Unmarshal([]byte(response.ResultJson), &beacon)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon2.Udi, beacon.Udi)
	assert.Equal(t, Beacon2.Site_id, beacon.Site_id)
	assert.Equal(t, Beacon2.Type, beacon.Type)
	assert.Equal(t, Beacon2.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Get all beacons
	request.Method = "v1.beacons.get_beacons"
	request.ArgsEmpty = false
	request.ArgsJson = "{}"
	response, err = client.Invoke(context.TODO(), &request)
	assert.Nil(t, err)

	var page bdata.BeaconV1DataPage
	jsonErr = json.Unmarshal([]byte(response.ResultJson), &page)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)
	beacon1 = *page.Data[0]

	// Update the beacon
	beacon1.Label = "ABC"
	requestParams = make(map[string]interface{})
	requestParams["beacon"] = beacon1
	jsonBuf, _ = json.Marshal(requestParams)

	request.Method = "v1.beacons.update_beacon"
	request.ArgsEmpty = false
	request.ArgsJson = string(jsonBuf)
	response, err = client.Invoke(context.TODO(), &request)
	assert.Nil(t, err)

	jsonErr = json.Unmarshal([]byte(response.ResultJson), &beacon)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon1.Id, beacon.Id)
	assert.Equal(t, "ABC", beacon.Label)

	// Get beacon by udi
	requestParams = make(map[string]interface{})
	requestParams["udi"] = beacon1.Udi
	jsonBuf, _ = json.Marshal(requestParams)

	request.Method = "v1.beacons.get_beacon_by_udi"
	request.ArgsEmpty = false
	request.ArgsJson = string(jsonBuf)
	response, err = client.Invoke(context.TODO(), &request)
	assert.Nil(t, err)

	jsonErr = json.Unmarshal([]byte(response.ResultJson), &beacon)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon1.Id, beacon.Id)

	// Calculate position for one beacon
	requestParams = make(map[string]interface{})
	requestParams["site_id"] = "1"
	requestParams["udis"] = []string{"00001"}
	jsonBuf, _ = json.Marshal(requestParams)

	request.Method = "v1.beacons.calculate_position"
	request.ArgsEmpty = false
	request.ArgsJson = string(jsonBuf)
	response, err = client.Invoke(context.TODO(), &request)
	assert.Nil(t, err)

	var position bdata.GeoPointV1
	jsonErr = json.Unmarshal([]byte(response.ResultJson), &position)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)

	assert.NotNil(t, position)
	assert.Equal(t, "Point", position.Type)
	assert.Equal(t, (float32)(0.0), position.Lat)
	assert.Equal(t, (float32)(0.0), position.Lng)

	// Delete the beacon
	requestParams = make(map[string]interface{})
	requestParams["beacon_id"] = beacon1.Id
	jsonBuf, _ = json.Marshal(requestParams)

	request.Method = "v1.beacons.delete_beacon_by_id"
	request.ArgsEmpty = false
	request.ArgsJson = string(jsonBuf)
	response, err = client.Invoke(context.TODO(), &request)
	assert.Nil(t, err)

	jsonErr = json.Unmarshal([]byte(response.ResultJson), &beacon)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon1.Id, beacon.Id)

	// Try to get deleted beacon
	requestParams = make(map[string]interface{})
	requestParams["beacon_id"] = beacon1.Id
	jsonBuf, _ = json.Marshal(requestParams)

	request.Method = "v1.beacons.get_beacon_by_id"
	request.ArgsEmpty = false
	request.ArgsJson = string(jsonBuf)
	response, err = client.Invoke(context.TODO(), &request)
	assert.Nil(t, err)

	beacon = bdata.BeaconV1{}

	jsonErr = json.Unmarshal([]byte(response.ResultJson), &beacon)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Empty(t, beacon)
}
