package main

import (
	//导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
	//导入sqlx包
	"github.com/jmoiron/sqlx"
)

func main() {

	//定义数据库对象
	var db *sqlx.DB

	//定义mysql数据源，配置数据库地址，帐号以及密码， dsn格式下面会解释
	dsn := "root:root1243@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"

	//根据数据源dsn和mysql驱动, 创建数据库对象
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

}
