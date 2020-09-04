package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {

	db, err := gorm.Open("mysql", "root:root1243@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

	fmt.Println(product.Code, product.Price)

	db.Model(&product).Update("Price", 200)

	db.First(&product, "code = ?", "D42")

	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)


	fmt.Println(product.Code, product.Price)


	defer db.Close()
}
