package logic

import (
	"encoding/json"
	"strings"

	ccomand "github.com/pip-services3-go/pip-services3-commons-go/commands"
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	crun "github.com/pip-services3-go/pip-services3-commons-go/run"
	cvalid "github.com/pip-services3-go/pip-services3-commons-go/validate"
	bdata "github.com/pip-templates/pip-templates-microservice-go/src/data/version1"
)

type BeaconsCommandSet struct {
	ccomand.CommandSet
	controller IBeaconsController
}

func NewBeaconsCommandSet(controller IBeaconsController) *BeaconsCommandSet {

	bcs := BeaconsCommandSet{}
	bcs.CommandSet = *ccomand.NewCommandSet()
	bcs.controller = controller
	bcs.AddCommand(bcs.makeGetBeaconsCommand())
	bcs.AddCommand(bcs.makeGetBeaconByIdCommand())
	bcs.AddCommand(bcs.makeGetBeaconByUdiCommand())
	bcs.AddCommand(bcs.makeCalculatePositionCommand())
	bcs.AddCommand(bcs.makeCreateBeaconCommand())
	bcs.AddCommand(bcs.makeUpdateBeaconCommand())
	bcs.AddCommand(bcs.makeDeleteBeaconByIdCommand())
	return &bcs
}

func (c *BeaconsCommandSet) makeGetBeaconsCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"get_beacons",
		cvalid.NewObjectSchema().
			WithOptionalProperty("filter", cvalid.NewFilterParamsSchema()).
			WithOptionalProperty("paging", cvalid.NewPagingParamsSchema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			filter := cdata.NewFilterParamsFromValue(args.Get("filter"))
			paging := cdata.NewPagingParamsFromValue(args.Get("paging"))
			return c.controller.GetBeacons(correlationId, filter, paging)
		})
}

func (c *BeaconsCommandSet) makeGetBeaconByIdCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"get_beacon_by_id",
		cvalid.NewObjectSchema().
			WithRequiredProperty("beacon_id", cconv.String),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			beaconId := args.GetAsString("beacon_id")
			return c.controller.GetBeaconById(correlationId, beaconId)
		})
}

func (c *BeaconsCommandSet) makeGetBeaconByUdiCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"get_beacon_by_udi",
		cvalid.NewObjectSchema().
			WithRequiredProperty("udi", cconv.String),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			udi := args.GetAsString("udi")
			return c.controller.GetBeaconByUdi(correlationId, udi)
		})
}

func (c *BeaconsCommandSet) makeCalculatePositionCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"calculate_position",
		cvalid.NewObjectSchema().
			WithRequiredProperty("site_id", cconv.String).
			WithRequiredProperty("udis", cvalid.NewArraySchema(cconv.String)), //cconv.String),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			siteId := args.GetAsString("site_id")
			udis := args.GetAsString("udis")
			arrUdis := make([]string, 0, 0)
			arrUdis = strings.Split(udis, ",")
			return c.controller.CalculatePosition(correlationId, siteId, arrUdis)
		})
}

func (c *BeaconsCommandSet) makeCreateBeaconCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"create_beacon",
		cvalid.NewObjectSchema().
			WithRequiredProperty("beacon", bdata.NewBeaconV1Schema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			val, _ := json.Marshal(args.Get("beacon"))
			var beacon bdata.BeaconV1
			json.Unmarshal(val, &beacon)

			return c.controller.CreateBeacon(correlationId, beacon)
		})
}

func (c *BeaconsCommandSet) makeUpdateBeaconCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"update_beacon",
		cvalid.NewObjectSchema().
			WithRequiredProperty("beacon", bdata.NewBeaconV1Schema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			val, _ := json.Marshal(args.Get("beacon"))
			var beacon bdata.BeaconV1
			json.Unmarshal(val, &beacon)
			return c.controller.UpdateBeacon(correlationId, beacon)
		})
}

func (c *BeaconsCommandSet) makeDeleteBeaconByIdCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"delete_beacon_by_id",
		cvalid.NewObjectSchema().
			WithRequiredProperty("beacon_id", cconv.String),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			beaconId := args.GetAsString("beacon_id")
			return c.controller.DeleteBeaconById(correlationId, beaconId)
		})
}
