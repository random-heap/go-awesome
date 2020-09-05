package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func notTransaction() {
	db, err := gorm.Open("mysql", "root:root1243@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 全局禁用表名复数
	db.SingularTable(true)

	//初始化
	var member = Member{
		Name: "yaomaoze",
		Gender: "M",
		Email: "yaomaoze@gmail.com",
		Phone: "15918711111",
		Password: "admin",
		Integral: 100,
		Birthday: "1991-05-11",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	// 增
	db.Model(member).Create(&member)

	// 删
	db.Where("id = 1").Delete(&Member{})

	// 查
	var m Member
	//db.First(&m)
	db.Where("name = ?", "yaomaoze").First(&m)

	fmt.Println(m.Name, m.CreateTime)

	// 改
	db.Model(&m).Where("name = ?", "yaomaoze").Update("name", "hong")

	// 原生SQL
	var m2 Member
	db.Raw("SELECT name, phone FROM member WHERE name = ?", "hong").Scan(&m2)
	fmt.Println("m2:", m2.Name, m2.Phone)
}

func transaction() {
	db, err := gorm.Open("mysql", "root:root1243@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 全局禁用表名复数
	db.SingularTable(true)

	tx := db.Begin()

	//初始化
	var member = Member{
		Name: "yaomaoze",
		Gender: "M",
		Email: "yaomaoze@gmail.com",
		Phone: "15918711111",
		Password: "admin",
		Integral: 100,
		Birthday: "1991-05-11",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	// 增
	if err := tx.Model(member).Create(&member).Error; err != nil {
		tx.Rollback()
		return
	}

	// 删
	if err := tx.Where("id = 10").Delete(&Member{}).Error; err != nil {
		tx.Rollback()
		return
	}

	a := 10 / 0
	fmt.Println(a)

	tx.Commit()
}

func main() {

	transaction()
}
