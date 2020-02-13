package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	bclients "github.com/pip-templates/pip-templates-microservice-go/src/clients/version1"
)

type BeaconsClientFactory struct {
	cbuild.Factory
	NullClientDescriptor   *cref.Descriptor
	DirectClientDescriptor *cref.Descriptor
	HttpClientDescriptor   *cref.Descriptor
}

func NewBeaconsClientFactory() *BeaconsClientFactory {

	bcf := BeaconsClientFactory{}
	bcf.Factory = *cbuild.NewFactory()

	bcf.NullClientDescriptor = cref.NewDescriptor("beacons", "client", "null", "*", "1.0")
	bcf.DirectClientDescriptor = cref.NewDescriptor("beacons", "client", "direct", "*", "1.0")
	bcf.HttpClientDescriptor = cref.NewDescriptor("beacons", "client", "http", "*", "1.0")

	bcf.RegisterType(bcf.NullClientDescriptor, bclients.NewBeaconsNullClientV1)
	bcf.RegisterType(bcf.DirectClientDescriptor, bclients.NewBeaconsDirectClientV1)
	bcf.RegisterType(bcf.HttpClientDescriptor, bclients.NewBeaconsHttpClientV1)
	return &bcf
}
