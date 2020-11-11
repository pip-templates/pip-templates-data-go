package test_services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	logic "github.com/pip-templates/pip-templates-microservice-go/logic"
	persist "github.com/pip-templates/pip-templates-microservice-go/persistence"
	services1 "github.com/pip-templates/pip-templates-microservice-go/services/version1"
	"github.com/stretchr/testify/assert"
)

func TestBeaconsRestService(t *testing.T) {
	restConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", "localhost",
		"connection.port", "3000",
	)

	var persistence *persist.BeaconsMemoryPersistence
	persistence = persist.NewBeaconsMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller := logic.NewBeaconsController()

	var service *services1.BeaconsRestService
	service = services1.NewBeaconsRestService()
	service.Configure(restConfig)

	var references *cref.References = cref.NewReferencesFromTuples(
		cref.NewDescriptor("beacons", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("beacons", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("beacons", "service", "rest", "default", "1.0"), service,
	)

	controller.SetReferences(references)
	service.SetReferences(references)

	opnErr := persistence.Open("")
	if opnErr != nil {
		panic("Can't open persistence")
	}
	service.Open("")

	defer service.Close("")
	defer persistence.Close("")

	url := "http://localhost:3000/v1/beacons"

	var beacon1 data1.BeaconV1

	// Create one beacon
	jsonBody, _ := json.Marshal(Beacon1)

	bodyReader := bytes.NewReader(jsonBody)
	postResponse, postErr := http.Post(url+"/beacons", "application/json", bodyReader)
	assert.Nil(t, postErr)
	resBody, bodyErr := ioutil.ReadAll(postResponse.Body)
	assert.Nil(t, bodyErr)

	var beacon data1.BeaconV1
	jsonErr := json.Unmarshal(resBody, &beacon)

	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon1.Udi, beacon.Udi)
	assert.Equal(t, Beacon1.SiteId, beacon.SiteId)
	assert.Equal(t, Beacon1.Type, beacon.Type)
	assert.Equal(t, Beacon1.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	beacon1 = beacon

	// Create another beacon
	jsonBody, _ = json.Marshal(Beacon2)

	bodyReader = bytes.NewReader(jsonBody)
	postResponse, postErr = http.Post(url+"/beacons", "application/json", bodyReader)
	assert.Nil(t, postErr)
	resBody, bodyErr = ioutil.ReadAll(postResponse.Body)
	assert.Nil(t, bodyErr)

	jsonErr = json.Unmarshal(resBody, &beacon)

	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon2.Udi, beacon.Udi)
	assert.Equal(t, Beacon2.SiteId, beacon.SiteId)
	assert.Equal(t, Beacon2.Type, beacon.Type)
	assert.Equal(t, Beacon2.Label, beacon.Label)
	assert.NotNil(t, beacon.Center)

	// Get all dummies
	getResponse, getErr := http.Get(url + "/beacons")
	assert.Nil(t, getErr)
	resBody, bodyErr = ioutil.ReadAll(getResponse.Body)
	assert.Nil(t, bodyErr)

	var page data1.BeaconV1DataPage
	jsonErr = json.Unmarshal(resBody, &page)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)

	// Update the beacon

	beacon1.Label = "ABC"
	jsonBody, _ = json.Marshal(beacon1)

	client := &http.Client{}
	data := bytes.NewReader(jsonBody)
	putReq, putErr := http.NewRequest(http.MethodPut, url+"/beacons", data)
	assert.Nil(t, putErr)
	putRes, putErr := client.Do(putReq)
	assert.Nil(t, putErr)
	resBody, bodyErr = ioutil.ReadAll(putRes.Body)
	jsonErr = json.Unmarshal(resBody, &beacon)
	assert.Nil(t, putErr)
	assert.NotNil(t, beacon)

	assert.Equal(t, beacon1.Id, beacon.Id)
	assert.Equal(t, "ABC", beacon.Label)

	// Get by udi
	getResponse, getErr = http.Get(url + "/beacons/udi/" + beacon1.Udi)
	assert.Nil(t, getErr)
	resBody, bodyErr = ioutil.ReadAll(getResponse.Body)
	assert.Nil(t, bodyErr)

	jsonErr = json.Unmarshal(resBody, &beacon)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)
	assert.Equal(t, Beacon1.Id, beacon.Id)

	postResponse, postErr = http.Get(url + "/v1/beacons/calculate_position/1/00001")
	assert.Nil(t, postErr)
	resBody, bodyErr = ioutil.ReadAll(postResponse.Body)
	assert.Nil(t, bodyErr)

	var position data1.GeoPointV1
	jsonErr = json.Unmarshal(resBody, &position)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, beacon)

	assert.NotNil(t, position)
	assert.Equal(t, "Point", position.Type)
	assert.Equal(t, (float32)(0.0), position.Coordinates[0][0])
	assert.Equal(t, (float32)(0.0), position.Coordinates[0][1])

	// Delete beacon
	delReq, delErr := http.NewRequest(http.MethodDelete, url+"/beacons/"+beacon.Id, nil)
	assert.Nil(t, delErr)
	_, delErr = client.Do(delReq)
	assert.Nil(t, delErr)

	// Try to get delete beacon
	page.Data = page.Data[:0]
	*page.Total = 0
	getResponse, getErr = http.Get(url + "/beacons/" + beacon1.Id)
	assert.Nil(t, getErr)
	resBody, bodyErr = ioutil.ReadAll(getResponse.Body)
	assert.Nil(t, bodyErr)
	jsonErr = json.Unmarshal(resBody, &page)
	assert.Nil(t, jsonErr)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 0)
}
