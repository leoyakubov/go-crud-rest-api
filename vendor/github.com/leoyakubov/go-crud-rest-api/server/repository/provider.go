package repository

import (
	"fmt"
	"github.com/leoyakubov/go-crud-rest-api/server/config"

	"github.com/jinzhu/gorm"
)

type IPersistenceProvider interface {
	TaskRepo() ITaskRepository
}

type PersistenceProvider struct {
	task       ITaskRepository
	repository *Repository
	DB         *gorm.DB
}

func NewPersistenceProvider(db *gorm.DB) IPersistenceProvider {
	return &PersistenceProvider{
		DB: db,
	}
}

func (pp *PersistenceProvider) TaskRepo() ITaskRepository {
	if pp.repository == nil {
		pp.repository = NewRepositiry(pp.DB)
	}

	if pp.task == nil {
		pp.task = NewTaskRepository(pp.repository)
	}

	return pp.task
}

func NewDb(conf *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(conf.DBDialect, ConnString(conf))

	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.SingularTable(conf.DBGormSingularTable)
	db.LogMode(conf.DBGormLogMode)

	return db, nil
}

func ConnString(conf *config.Config) string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s%s",
		conf.DBUsername, conf.DBPassword, conf.DBProtocol,
		conf.DBHostname, conf.DBPort, conf.DBName, conf.DBParameters)
}
