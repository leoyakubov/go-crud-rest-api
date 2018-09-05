package controller

import (
	"go-crud-rest-api/server/config"
)

type BaseController struct {
	Config *config.Config
}

func (bc *BaseController) SetConfig(conf *config.Config) {
	bc.Config = conf
}
