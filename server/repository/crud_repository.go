package repository

import (
	"fmt"
	"github.com/leoyakubov/go-crud-rest-api/server/response"

	"github.com/jinzhu/gorm"
)

var (
	IS_DELETED  = "is_deleted"
	NOT_DELETED = map[string]interface{}{IS_DELETED: false}
)

type Repository struct {
	DB *gorm.DB
}

func NewRepositiry(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Add(add interface{}) error {
	db := r.DB.Create(add)

	if db.Error != nil {
		return fmt.Errorf("En arror occured while adding new task %v", db.Error)
	}

	return nil
}

func (r *Repository) Update(upd interface{}, res interface{}, id int) error {
	db := r.DB.Where(NOT_DELETED).First(res, id).Update(upd)

	if db.RecordNotFound() {
		return response.ErrTaskNotFound
	}

	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (r *Repository) FindOneById(res interface{}, id int) (err error) {
	db := r.DB.Where(NOT_DELETED).First(res, id)

	if db.RecordNotFound() {
		return response.ErrTaskNotFound
	}

	return db.Error
}

func (r *Repository) FindAll(v interface{}) (err error) {
	if err = r.DB.Where(NOT_DELETED).Find(v).Error; err != nil {
		return err
	}

	return err
}

func (r *Repository) Delete(res interface{}, id int) error {
	db := r.DB.Where(NOT_DELETED).First(res, id).Update(IS_DELETED, true)

	if db.RecordNotFound() {
		return response.ErrTaskNotFound
	}

	if db.Error != nil {
		return db.Error
	}

	return nil
}
