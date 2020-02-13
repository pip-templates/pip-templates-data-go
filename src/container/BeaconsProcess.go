package container

import (
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	rpcbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
	bfactory "github.com/pip-templates/pip-templates-microservice-go/src/build"
)

// import { ProcessContainer } from "pip-services3-container-node";
// import { DefaultRpcFactory } from "pip-services3-rpc-node";

// import {BeaconsServiceFactory} from "../build/BeaconsServiceFactory";

type BeaconsProcess struct {
	cproc.ProcessContainer
}

func NewBeaconsProcess() *BeaconsProcess {

	// super("beacons", "Beacons microservice");
	bp := BeaconsProcess{}
	bp.ProcessContainer = *cproc.NewEmptyProcessContainer()
	bp.AddFactory(bfactory.NewBeaconsServiceFactory())
	bp.AddFactory(rpcbuild.NewDefaultRpcFactory())
	return &bp
}
