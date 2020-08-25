package services

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	services1 "github.com/pip-services3-go/pip-services3-rpc-go/services"
)

type BeaconsHttpServiceV1 struct {
	*services1.CommandableHttpService
}

func NewBeaconsHttpServiceV1() *BeaconsHttpServiceV1 {
	c := BeaconsHttpServiceV1{
		CommandableHttpService: services1.NewCommandableHttpService("v1/beacons"),
	}
	c.DependencyResolver.Put("controller", cref.NewDescriptor("beacons", "controller", "*", "*", "1.0"))
	return &c
}
