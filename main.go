package main

import (
	"flag"

	"github.com/leoyakubov/go-crud-rest-api/server/core"

	"github.com/Sirupsen/logrus"
)

func main() {
	path := flag.String("config", "", "App config file: -config=</path/to/config/file>")
	flag.Parse()

	logger := logrus.New()
	logger.Infoln("Config path: ", *path)

	if path == nil || *path == "" {
		logger.Fatal("Config file is reguired to start up the server!")
	}

	server := core.BuildServer()

	if err := server.Init(*path, logger); err != nil {
		logger.Fatal(err)
	}

	if err := server.Run(); err != nil {
		logger.Fatal(err)
	}
}
