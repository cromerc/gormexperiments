package repository

import (
	"errors"
	"github.com/cromerc/gormexperiments/internal/adapter/entity"
	"github.com/cromerc/gormexperiments/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRepository interface {
	domain.IRepository
	getDB() *gorm.DB
	closeTransaction() error
}

type Repository struct {
	db []*gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: []*gorm.DB{db}}
}

func (r *Repository) getDB() *gorm.DB {
	return r.db[len(r.db)-1]
}

func (r *Repository) closeTransaction() error {
	if len(r.db) <= 1 {
		return errors.New("no transaction to close")
	}
	index := len(r.db) - 1
	r.db = r.db[:index]
	return nil
}

func (r *Repository) Begin() error {
	r.db = append(r.db, r.db[len(r.db)-1].Begin())
	return nil
}

func (r *Repository) Rollback() error {
	if len(r.db) <= 1 {
		return errors.New("no transaction to rollback")
	}
	r.getDB().Rollback()
	return r.closeTransaction()
}

func (r *Repository) Commit() error {
	if len(r.db) <= 1 {
		return errors.New("no transaction to commit")
	}
	r.getDB().Commit()
	return r.closeTransaction()
}

func (r *Repository) CreateProduct(product entity.Product) (*entity.Product, error) {
	r.getDB().Create(&product)
	return &product, nil
}

func (r *Repository) FindProduct(id uuid.UUID) (*entity.Product, error) {
	var product entity.Product
	r.getDB().First(&product, id)
	return &product, nil
}

func (r *Repository) FindProductByCode(product entity.Product) (*entity.Product, error) {
	r.getDB().First(&product, "code = ?", product.Code)
	return &product, nil
}

func (r *Repository) UpdateProduct(product entity.Product) (*entity.Product, error) {
	r.getDB().Model(&product).Updates(product)
	return &product, nil
}

func (r *Repository) DeleteProduct(id uuid.UUID) error {
	r.getDB().Delete(&entity.Product{}, id)
	return nil
}
