package services1

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	services1 "github.com/pip-services3-go/pip-services3-grpc-go/services"
)

type BeaconsGrpcServiceV1 struct {
	*services1.CommandableGrpcService
}

func NewBeaconsGrpcServiceV1() *BeaconsGrpcServiceV1 {
	c := BeaconsGrpcServiceV1{
		CommandableGrpcService: services1.NewCommandableGrpcService("v1.beacons"),
	}
	c.DependencyResolver.Put("controller", cref.NewDescriptor("beacons", "controller", "*", "*", "1.0"))
	return &c
}
