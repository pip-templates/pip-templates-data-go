package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	clients1 "github.com/pip-templates/pip-templates-microservice-go/clients/version1"
)

type BeaconsClientFactory struct {
	cbuild.Factory
}

func NewBeaconsClientFactory() *BeaconsClientFactory {
	c := BeaconsClientFactory{}
	c.Factory = *cbuild.NewFactory()

	nullClientDescriptor := cref.NewDescriptor("beacons", "client", "null", "*", "1.0")
	directClientDescriptor := cref.NewDescriptor("beacons", "client", "direct", "*", "1.0")
	httpClientDescriptor := cref.NewDescriptor("beacons", "client", "http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("beacons", "client", "grpc", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewBeaconsNullClientV1)
	c.RegisterType(directClientDescriptor, clients1.NewBeaconsDirectClientV1)
	c.RegisterType(httpClientDescriptor, clients1.NewBeaconsHttpClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewBeaconsGrpcClientV1)
	return &c
}
