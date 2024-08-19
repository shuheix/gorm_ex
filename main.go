package main

import (
	"github.com/davecgh/go-spew/spew"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	//db.AutoMigrate(&Product{})

	//db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)

	// {{1 2024-08-19 20:18:52.661687111 +0900 +0900 2024-08-19 20:18:52.661687111 +0900 +0900 {0001-01-01 00:00:00 +0000 UTC false}} D42 100}
	spew.Dump(product)
}
