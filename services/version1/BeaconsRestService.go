package services1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	crefer "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cvalid "github.com/pip-services3-go/pip-services3-commons-go/validate"
	"github.com/pip-services3-go/pip-services3-rpc-go/services"
	data1 "github.com/pip-templates/pip-templates-microservice-go/data/version1"
	logic "github.com/pip-templates/pip-templates-microservice-go/logic"
)

type BeaconsRestService struct {
	*services.RestService
	controller    logic.IBeaconsController
	numberOfCalls int
}

func NewBeaconsRestService() *BeaconsRestService {
	c := BeaconsRestService{}
	c.RestService = services.NewRestService()
	c.RestService.IRegisterable = &c
	c.numberOfCalls = 0
	c.BaseRoute = "v1/beacons"
	c.DependencyResolver.Put("controller", crefer.NewDescriptor("beacons", "controller", "default", "*", "*"))
	return &c
}

func (c *BeaconsRestService) SetReferences(references crefer.IReferences) {
	c.RestService.SetReferences(references)
	depRes, depErr := c.DependencyResolver.GetOneRequired("controller")
	if depErr == nil && depRes != nil {
		c.controller = depRes.(logic.IBeaconsController)
	}
}

func (c *BeaconsRestService) getPageByFilter(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	paginParams := make(map[string]string, 0)

	paginParams["skip"] = params.Get("skip")
	paginParams["take"] = params.Get("take")
	paginParams["total"] = params.Get("total")

	delete(params, "skip")
	delete(params, "take")
	delete(params, "total")

	result, err := c.controller.GetBeacons(
		params.Get("correlation_id"),
		cdata.NewFilterParamsFromValue(params), // W! need test
		cdata.NewPagingParamsFromTuples(paginParams),
	)
	c.SendResult(res, req, result, err)
}

func (c *BeaconsRestService) getOneById(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	vars := mux.Vars(req)

	beaconId := params.Get("beacon_id")
	if beaconId == "" {
		beaconId = vars["beacon_id"]
	}
	result, err := c.controller.GetBeaconById(
		params.Get("correlation_id"),
		beaconId)
	c.SendResult(res, req, result, err)
}

func (c *BeaconsRestService) getOneByUdi(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	vars := mux.Vars(req)

	beaconUdi := params.Get("udi")
	if beaconUdi == "" {
		beaconUdi = vars["udi"]
	}
	result, err := c.controller.GetBeaconByUdi(
		params.Get("correlation_id"),
		beaconUdi)
	c.SendResult(res, req, result, err)
}

func (c *BeaconsRestService) calculatePosition(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	vars := mux.Vars(req)

	beaconUdi := params.Get("beacon_udi")
	if beaconUdi == "" {
		beaconUdi = vars["beacon_udi"]
	}
	arrUdis := make([]string, 0, 0)
	arrUdis = strings.Split(beaconUdi, ",")

	siteId := params.Get("site_id")
	if siteId == "" {
		siteId = vars["site_id"]
	}

	result, err := c.controller.CalculatePosition(
		params.Get("correlation_id"),
		siteId,
		arrUdis)
	c.SendResult(res, req, result, err)
}

func (c *BeaconsRestService) create(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	correlationId := params.Get("correlation_id")
	var beacon data1.BeaconV1

	body, bodyErr := ioutil.ReadAll(req.Body)
	if bodyErr != nil {
		err := cerr.NewInternalError(correlationId, "JSON_CNV_ERR", "Cant convert from JSON to Beacons").WithCause(bodyErr)
		c.SendError(res, req, err)
		return
	}
	defer req.Body.Close()
	jsonErr := json.Unmarshal(body, &beacon)

	if jsonErr != nil {
		err := cerr.NewInternalError(correlationId, "JSON_CNV_ERR", "Cant convert from JSON to Beacons").WithCause(jsonErr)
		c.SendError(res, req, err)
		return
	}

	result, err := c.controller.CreateBeacon(
		correlationId,
		&beacon,
	)
	c.SendCreatedResult(res, req, result, err)
}

func (c *BeaconsRestService) update(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	correlationId := params.Get("correlation_id")

	var beacon data1.BeaconV1

	body, bodyErr := ioutil.ReadAll(req.Body)
	if bodyErr != nil {
		err := cerr.NewInternalError(correlationId, "JSON_CNV_ERR", "Cant convert from JSON to Beacons").WithCause(bodyErr)
		c.SendError(res, req, err)
		return
	}
	defer req.Body.Close()
	jsonErr := json.Unmarshal(body, &beacon)

	if jsonErr != nil {
		err := cerr.NewInternalError(correlationId, "JSON_CNV_ERR", "Cant convert from JSON to Beacons").WithCause(jsonErr)
		c.SendError(res, req, err)
		return
	}
	result, err := c.controller.UpdateBeacon(
		correlationId,
		&beacon,
	)
	c.SendResult(res, req, result, err)
}

func (c *BeaconsRestService) deleteById(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	vars := mux.Vars(req)

	beaconId := params.Get("beacon_id")
	if beaconId == "" {
		beaconId = vars["beacon_id"]
	}

	result, err := c.controller.DeleteBeaconById(
		params.Get("correlation_id"),
		beaconId,
	)
	c.SendDeletedResult(res, req, result, err)
}

func (c *BeaconsRestService) Register() {

	c.RegisterRoute(
		"get", "/beacons",
		&cvalid.NewObjectSchema().WithOptionalProperty("skip", cconv.String).
			WithOptionalProperty("take", cconv.String).
			WithOptionalProperty("total", cconv.String).
			WithOptionalProperty("body", cvalid.NewFilterParamsSchema()).Schema,
		c.getPageByFilter,
	)

	c.RegisterRoute(
		"get", "/beacons/{beacon_id}",
		&cvalid.NewObjectSchema().
			WithRequiredProperty("beacon_id", cconv.String).Schema,
		c.getOneById,
	)

	c.RegisterRoute(
		"get", "/beacons/udi/{beacon_udi}",
		&cvalid.NewObjectSchema().
			WithRequiredProperty("beacon_udi", cconv.String).Schema,
		c.getOneByUdi,
	)

	c.RegisterRoute(
		"get", "/beacons/calculate_position/{site_id}/{beacon_udi}",
		&cvalid.NewObjectSchema().
			WithRequiredProperty("site_id", cconv.String).
			WithRequiredProperty("beacon_udi", cconv.String).Schema,
		c.calculatePosition,
	)

	c.RegisterRoute(
		"post", "/beacons",
		&cvalid.NewObjectSchema().
			WithRequiredProperty("body", data1.NewBeaconV1Schema()).Schema,
		c.create,
	)

	c.RegisterRoute(
		"put", "/beacons",
		&cvalid.NewObjectSchema().
			WithRequiredProperty("body", data1.NewBeaconV1Schema()).Schema,
		c.update,
	)

	c.RegisterRoute(
		"delete", "/beascon/{beacon_id}",
		&cvalid.NewObjectSchema().
			WithRequiredProperty("beacon_id", cconv.String).Schema,
		c.deleteById,
	)
}
