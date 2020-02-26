package services

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	grpcservices "github.com/pip-services3-go/pip-services3-grpc-go/services"
)

type BeaconsGrpcServiceV1 struct {
	*grpcservices.CommandableGrpcService
}

func NewBeaconsGrpcServiceV1() *BeaconsGrpcServiceV1 {
	bchs := BeaconsGrpcServiceV1{
		CommandableGrpcService: grpcservices.NewCommandableGrpcService("v1.beacons"),
	}
	bchs.DependencyResolver.Put("controller", cref.NewDescriptor("beacons", "controller", "*", "*", "1.0"))
	return &bchs
}
