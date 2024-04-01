package main

import (
	"fmt"
	"github.com/cromerc/gormexperiments/internal/adapter/entity"
	"github.com/cromerc/gormexperiments/internal/adapter/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		panic(err)
	}

	repo := repository.New(db)

	newProduct, err := repo.CreateProduct(entity.Product{Code: "D41", Price: 10})
	fmt.Printf("%-v\n", newProduct)

	err = repo.Begin()
	if err != nil {
		panic(err)
	}
	newProduct, err = repo.CreateProduct(entity.Product{Code: "D42", Price: 100})
	err = repo.Rollback()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%-v\n", newProduct)

	err = repo.Begin()
	if err != nil {
		panic(err)
	}
	newProduct, err = repo.CreateProduct(entity.Product{Code: "D55", Price: 1000})
	err = repo.Commit()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%-v\n", newProduct)
}
