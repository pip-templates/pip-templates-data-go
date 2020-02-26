package container

import (
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	rpcbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
	bfactory "github.com/pip-templates/pip-templates-microservice-go/build"
)

type BeaconsProcess struct {
	cproc.ProcessContainer
}

func NewBeaconsProcess() *BeaconsProcess {

	bp := BeaconsProcess{}
	bp.ProcessContainer = *cproc.NewEmptyProcessContainer()
	bp.AddFactory(bfactory.NewBeaconsServiceFactory())
	bp.AddFactory(rpcbuild.NewDefaultRpcFactory())
	return &bp
}
