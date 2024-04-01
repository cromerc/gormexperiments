package domain

import (
	"github.com/cromerc/gormexperiments/internal/adapter/entity"
	"github.com/google/uuid"
)

type IRepository interface {
	Begin() error
	Rollback() error
	Commit() error
	CreateProduct(product entity.Product) (*entity.Product, error)
	FindProduct(id uuid.UUID) (*entity.Product, error)
	FindProductByCode(product entity.Product) (*entity.Product, error)
	UpdateProduct(product entity.Product) (*entity.Product, error)
	DeleteProduct(id uuid.UUID) error
}
