package services

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	rpcservices "github.com/pip-services3-go/pip-services3-rpc-go/services"
)

type BeaconsHttpServiceV1 struct {
	*rpcservices.CommandableHttpService
}

func NewBeaconsHttpServiceV1() *BeaconsHttpServiceV1 {
	bchs := BeaconsHttpServiceV1{
		CommandableHttpService: rpcservices.NewCommandableHttpService("v1/beacons"),
	}
	bchs.DependencyResolver.Put("controller", cref.NewDescriptor("beacons", "controller", "*", "*", "1.0"))
	return &bchs
}
