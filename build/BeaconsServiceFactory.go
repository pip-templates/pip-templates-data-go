package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	blogic "github.com/pip-templates/pip-templates-microservice-go/logic"
	bpersist "github.com/pip-templates/pip-templates-microservice-go/persistence"
	bservices "github.com/pip-templates/pip-templates-microservice-go/services/version1"
)

type BeaconsServiceFactory struct {
	cbuild.Factory
	MemoryPersistenceDescriptor  *cref.Descriptor
	FilePersistenceDescriptor    *cref.Descriptor
	MongoDbPersistenceDescriptor *cref.Descriptor
	ControllerDescriptor         *cref.Descriptor
	HttpServiceV1Descriptor      *cref.Descriptor
	GrpcServiceV1Descriptor      *cref.Descriptor
}

func NewBeaconsServiceFactory() *BeaconsServiceFactory {

	bsf := BeaconsServiceFactory{}
	bsf.Factory = *cbuild.NewFactory()

	bsf.MemoryPersistenceDescriptor = cref.NewDescriptor("beacons", "persistence", "memory", "*", "1.0")
	bsf.FilePersistenceDescriptor = cref.NewDescriptor("beacons", "persistence", "file", "*", "1.0")
	bsf.MongoDbPersistenceDescriptor = cref.NewDescriptor("beacons", "persistence", "mongodb", "*", "1.0")
	bsf.ControllerDescriptor = cref.NewDescriptor("beacons", "controller", "default", "*", "1.0")
	bsf.HttpServiceV1Descriptor = cref.NewDescriptor("beacons", "service", "http", "*", "1.0")
	bsf.GrpcServiceV1Descriptor = cref.NewDescriptor("beacons", "service", "grpc", "*", "1.0")

	bsf.RegisterType(bsf.MemoryPersistenceDescriptor, bpersist.NewBeaconsMemoryPersistence)
	bsf.RegisterType(bsf.FilePersistenceDescriptor, bpersist.NewBeaconsFilePersistence)
	bsf.RegisterType(bsf.MongoDbPersistenceDescriptor, bpersist.NewBeaconsMongoDbPersistence)
	bsf.RegisterType(bsf.ControllerDescriptor, blogic.NewBeaconsController)
	bsf.RegisterType(bsf.HttpServiceV1Descriptor, bservices.NewBeaconsHttpServiceV1)
	bsf.RegisterType(bsf.GrpcServiceV1Descriptor, bservices.NewBeaconsGrpcServiceV1)
	return &bsf
}
