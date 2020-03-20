# <img src="https://github.com/pip-services/pip-services/raw/master/design/Logo.png" alt="Pip.Services Logo" style="max-width:30%"> <br/> Beacons microservice

This is the Beacons microservice from the Pip.Templates library. 

The microservice currently supports the following deployment options:
* Deployment platforms: Standalone Process
* External APIs: HTTP/REST, gRPC
* Persistence: Memory, Flat Files, MongoDB

This microservice does not depend on other microservices.

<a name="links"></a> Quick Links:

* [Download Links](doc/Downloads.md)
* [Development Guide](doc/Development.md)
* [Configuration Guide](doc/Configuration.md)
* [Deployment Guide](doc/Deployment.md)
* Communication Protocols
  - [HTTP Version 1](doc/HttpProtocolV1.md)
  <!-- Todo: gRPC  -->

## Contract

The logical contract of the microservice is presented below. 

```go
//implements IStringIdentifiable
type BeaconV1 struct {
	Id      string     `json:"id" bson:"_id"`
	Site_id string     `json:"site_id" bson:"site_id"`
	Type    string     `json:"type" bson:"type"`
	Udi     string     `json:"udi" bson:"udi"`
	Label   string     `json:"label" bson:"label"`
	Center  GeoPointV1 `json:"center" bson:"center"` // GeoJson
	Radius  float32    `json:"radius" bson:"radius"`
}

type GeoPointV1 struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates [][]float32 `json:"coordinates" bson:"coordinates"`
}

type IBeaconsClientV1 interface {
	GetBeacons(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *bdata.BeaconV1DataPage, err error)

	GetBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error)

	GetBeaconByUdi(correlationId string, udi string) (beacon *bdata.BeaconV1, err error)

	CalculatePosition(correlationId string, siteId string, udis []string) (position *bdata.GeoPointV1, err error)

	CreateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error)

	UpdateBeacon(correlationId string, beacon bdata.BeaconV1) (res *bdata.BeaconV1, err error)

	DeleteBeaconById(correlationId string, beaconId string) (beacon *bdata.BeaconV1, err error)
}
```

## Download

Right now, the only way to get the microservice is to check it out directly from the GitHub repository
```bash
git clone https://github.com/pip-templates/pip-templates-microservice-go.git
```

The Pip.Service team is working on implementing packaging, to make stable releases available as zip-downloadable archives.

## Run

Add the **config.yml** file to the config folder and set configuration parameters as needed.

Example of a microservice configuration
```yaml
---
  # Container descriptor
  - descriptor: "pip-services:context-info:default:default:1.0"
    name: "beacons"
    description: "Beacons microservice"
  
  # Console logger
  - descriptor: "pip-services:logger:console:default:1.0"
    level: "trace"
  
  # Perfomance counter that post values to log
  - descriptor: "pip-services:counters:log:default:1.0"
  
  {{^if MONGO_ENABLED}}{{^if FILE_ENABLED}}
  # Memory persistence
  - descriptor: "beacons:persistence:memory:default:1.0"
  {{/if}}{{/if}}
  
  {{#if FILE_ENABLED}}
  # File persistence
  - descriptor: "beacons:persistence:file:default:1.0"
    path: {{FILE_PATH}}{{^if FILE_PATH}}"./data/beacons.json"{{/if}}
  {{/if}}
  
  {{#if MONGO_ENABLED}}
  # MongoDb persistence
  - descriptor: "beacons:persistence:mongodb:default:1.0"
    connection:
      uri: {{MONGO_SERVICE_URI}}
      host: {{MONGO_SERVICE_HOST}}{{^if MONGO_SERVICE_HOST}}"localhost"{{/if}}
      port: {{MONGO_SERVICE_PORT}}{{^if MONGO_SERVICE_PORT}}27017{{/if}}
      database: {{MONGO_DB}}{{^if MONGO_DB}}"test"{{/if}}
  {{/if}}
  
  {{#if COUCHBASE_ENABLED}}
    # Couchbase persistence
    - descriptor: "beacons:persistence:couchbase:default:1.0"
      connection:
        uri: {{COUCHBASE_SERVICE_URI}}
        host: {{COUCHBASE_SERVICE_HOST}}{{^if COUCHBASE_SERVICE_HOST}}"localhost"{{/if}}
        port: {{COUCHBASE_SERVICE_PORT}}{{^if COUCHBASE_SERVICE_PORT}}8091{{/if}}
        operation_timeout: 2
        detailed_errcodes: 1
      options:
        auto_create: false
        auto_index: true
      credential:
            username: {{COUCHBASE_SERVICE_USER}}{{^if COUCHBASE_SERVICE_USER}}"Administrator"{{/if}}
            password: {{COUCHBASE_SERVICE_PASSWD}}{{^if COUCHBASE_SERVICE_PASSWD}}"password"{{/if}}
    {{/if}}
    
  # Controller
  - descriptor: "beacons:controller:default:default:1.0"

  # Shared HTTP Endpoint
  - descriptor: "pip-services:endpoint:http:default:1.0"
    connection:
      protocol: http
      host: 0.0.0.0
      port: {{HTTP_PORT}}{{^if HTTP_PORT}}8080{{/if}}
  
  # HTTP Service V1
  - descriptor: "beacons:service:http:default:1.0"
  
  # Hearbeat service
  - descriptor: "pip-services:heartbeat-service:http:default:1.0"
  
  # Status service
  - descriptor: "pip-services:status-service:http:default:1.0"  

  # gRPC Service V1
  - descriptor: "beacons:service:grpc:default:1.0"
    connection:
      protocol: http
      host: 0.0.0.0
      port: {{GRPC_PORT}}{{^if GRPC_PORT}}8081{{/if}}
```
<!-- Todo -->
<!-- For more information on microservice configuration, see [The Configuration Guide](Configuration.md). -->

The microservice can be started using the command:
```bash
go bin/run.go
```

## Use

The easiest way to work with the microservice is through the client SDK. 

If you use GoLang, then get references to the required libraries:
- Pip.Services3.Commons : https://github.com/pip-services3-go/pip-services3-commons-go
- Pip.Services3.Rpc: 
https://github.com/pip-services3-go/pip-services3-rpc-go

Import needed modules from **pip-services3-commons-go** and **pip-templates-microservice-go**
```go
import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
  bclients "github.com/pip-templates/pip-templates-microservice-go/clients/version1"
  bdata "github.com/pip-templates/pip-templates-microservice-go/data/version1"
  cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)
```

Define client configuration parameters that match the configuration of the microservice's external API
```go
// Client configuration
httpConfig := cconf.NewConfigParamsFromTuples(
  "connection.protocol", "http",
  "connection.port", "3000",
  "connection.host", "localhost",
)
```

Instantiate the client and open a connection to the microservice
```go
// Create the client instance
var client := bclients.NewBeaconsHttpClientV1()

// Configure the client
client.Configure(httpConfig)

// Connect to the microservice
client.Open("")
    
// Work with the microservice
...
```

The client is now ready to perform operations
```go
// Define a beacon
var beacon := bdata.BeaconV1{
    Id:      "1",
    Udi:     "00001",
    Type:    bdata.BeaconTypeV1.AltBeacon,
    Site_id: "1",
    Label:   "TestBeacon",
    Center:  bdata.GeoPointV1{Type: "Point", Coordinates: [][]float32{{0.0, 0.0}}},
    Radius:  50,
}

// Create the beacon
result, err := client.CreateBeacon("", beacon)

// Retrieve a DataPage
page, err := client.GetBeacons("", cdata.NewFilterParamsFromTuples("label", "TestBeacon"), cdata.NewPagingParams(0, 10))

// Do something with the returned page...
// E.g. var retBeacon := *page.Data[0]
```

## Acknowledgements

This microservice was created and currently maintained by *Sergey Seroukhov*.
