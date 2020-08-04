package main

import (
	"os"

	cont "github.com/pip-templates/pip-templates-microservice-go/container"
)

func main() {
	proc := cont.NewBeaconsProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(os.Args)
}
