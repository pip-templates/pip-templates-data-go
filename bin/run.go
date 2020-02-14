package main

import (
	"os"

	bproc "github.com/pip-templates/pip-templates-microservice-go/src/container"
)

func main() {
	proc := bproc.NewBeaconsProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(os.Args)

}
