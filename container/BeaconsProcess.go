package container

import (
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	rpcbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
	factory "github.com/pip-templates/pip-templates-microservice-go/build"
)

type BeaconsProcess struct {
	cproc.ProcessContainer
}

func NewBeaconsProcess() *BeaconsProcess {
	c := BeaconsProcess{}
	c.ProcessContainer = *cproc.NewProcessContainer("beacons", "Beacons microservice")
	c.AddFactory(factory.NewBeaconsServiceFactory())
	c.AddFactory(rpcbuild.NewDefaultRpcFactory())
	return &c
}
