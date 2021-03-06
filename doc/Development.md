# Development and Testing Guide <br/> Beacons Microservice

This document provides high-level instructions on how to build and test the microservice.

* [Environment Setup](#setup)
* [Installing](#install)
* [Building](#build)
* [Testing](#test)

## <a name="setup"></a> Environment Setup

This is a Golang project, so you'll have to install Golang tools to work with it. 
You can download them from the official website: https://golang.org/dl/

You can check that Go has been successfully installed by running the following command:
```bash
go version
```

To work with the GitHub code repository you need to install Git from: https://git-scm.com/downloads

If you plan on developing and testing with persistent storages other than in-memory/flat file persistences,
you may need to install a database server:
- Download and install MongoDB from: https://www.mongodb.org/downloads
- Download and install Couchbase from: https://www.couchbase.com/downloads?family=server

## <a name="install"></a> Installing

Once your environment is ready, you can check out the microservice's source code from the GitHub repository:
```bash
git clone https://github.com/pip-templates/pip-templates-microservice-go.git
```

Then go to the project folder and install dependent modules:
```bash
# Install dependencies
go install ./bin/run.go
```

If you've worked with the microservice before, you can check out the latest changes and update the dependencies:
```bash
# Update source code updates from GitHub
git pull
```

## <a name="build"></a> Building

This microservice is written in the language Go, so, if you make changes to the source code, you need to compile 
it before running or committing to GitHub. The process will output an executable Go file named run.go in the ./bin folder:

```bash
go build ./bin/run.go
```

## <a name="test"></a> Testing

The command to run unit tests is as follows:
```bash
go test -v ./test/...
```