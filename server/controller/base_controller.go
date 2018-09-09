package controller

import (
	"github.com/leoyakubov/go-crud-rest-api/server/config"
	"github.com/leoyakubov/go-crud-rest-api/server/repository"
)

type BaseController struct {
	Config              *config.Config
	PersistenceProvider repository.IPersistenceProvider
}

func (bc *BaseController) SetConfig(conf *config.Config) {
	bc.Config = conf
}

func (bc *BaseController) SetPersistenceProvider(pp repository.IPersistenceProvider) {
	bc.PersistenceProvider = pp
}
