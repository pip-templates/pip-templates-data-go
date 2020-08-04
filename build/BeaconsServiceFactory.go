package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	logic "github.com/pip-templates/pip-templates-microservice-go/logic"
	persist "github.com/pip-templates/pip-templates-microservice-go/persistence"
	services1 "github.com/pip-templates/pip-templates-microservice-go/services/version1"
)

type BeaconsServiceFactory struct {
	cbuild.Factory
}

func NewBeaconsServiceFactory() *BeaconsServiceFactory {
	c := BeaconsServiceFactory{}
	c.Factory = *cbuild.NewFactory()

	memoryPersistenceDescriptor := cref.NewDescriptor("beacons", "persistence", "memory", "*", "1.0")
	filePersistenceDescriptor := cref.NewDescriptor("beacons", "persistence", "file", "*", "1.0")
	mongoDbPersistenceDescriptor := cref.NewDescriptor("beacons", "persistence", "mongodb", "*", "1.0")
	couchbasePersistenceDescriptor := cref.NewDescriptor("beacons", "persistence", "couchbase", "*", "1.0")
	controllerDescriptor := cref.NewDescriptor("beacons", "controller", "default", "*", "1.0")
	httpServiceV1Descriptor := cref.NewDescriptor("beacons", "service", "http", "*", "1.0")
	grpcServiceV1Descriptor := cref.NewDescriptor("beacons", "service", "grpc", "*", "1.0")

	c.RegisterType(memoryPersistenceDescriptor, persist.NewBeaconsMemoryPersistence)
	c.RegisterType(filePersistenceDescriptor, persist.NewBeaconsFilePersistence)
	c.RegisterType(mongoDbPersistenceDescriptor, persist.NewBeaconsMongoDbPersistence)
	c.RegisterType(couchbasePersistenceDescriptor, persist.NewBeaconsCouchbasePersistence)
	c.RegisterType(controllerDescriptor, logic.NewBeaconsController)
	c.RegisterType(httpServiceV1Descriptor, services1.NewBeaconsHttpServiceV1)
	c.RegisterType(grpcServiceV1Descriptor, services1.NewBeaconsGrpcServiceV1)
	return &c
}
