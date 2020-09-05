package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)


type Member struct {
	Id int	`db:"id"`
	Name string	`db:"name"`
	Gender string	`db:"gender"`
	Email string	`db:"email"`
	Phone string	`db:"phone"`
	Password string	`db:"password"`
	Integral float32	`db:"integral"`
	Birthday string	`db:"birthday"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
}

func notTransaction() {
	//定义数据库对象
	var db *sqlx.DB

	//定义mysql数据源，配置数据库地址，帐号以及密码， dsn格式下面会解释
	dsn := "root:root1243@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"

	//根据数据源dsn和mysql驱动, 创建数据库对象
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

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
	fmt.Println(member)

	// 增
	//db.MustExec("INSERT INTO member(name, gender, email, phone, password, integral, birthday, create_time, update_time) " +
	//	"VALUES (?,?,?,?,?,?,?,?,?)",
	//	"yaomaoze", "M", "yaomaoze@gmail.com", "15918711111",  "admin", 100,  "1991-05-11", time.Now(),time.Now())

	//db.NamedExec("INSERT INTO member(name, gender, email, phone, password, integral, birthday, create_time, update_time) " +
	//	" VALUES (:name, :gender, :email, :phone, :password, :integral, :birthday, :create_time, :update_time)",
	//	&member)

	// 删
	db.MustExec("delete from member where id = ?", 7)

	var members []Member
	db.Select(&members, "SELECT name, gender FROM member ORDER BY name ASC")
	fmt.Println(len(members))
	for _, m := range members {
		fmt.Println(m.Name, m.Gender)
	}

	rows, err := db.Query("SELECT name FROM member ORDER BY name ASC")

	// iterate over each row
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		fmt.Println(name)
	}

	// 改
	db.MustExec("update member set name = ? where id = ? and name = ? ", "lailai", 1, "yaomaoze")

}

func transaction() {

	//定义数据库对象
	var db *sqlx.DB

	//定义mysql数据源，配置数据库地址，帐号以及密码， dsn格式下面会解释
	dsn := "root:root1243@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"

	//根据数据源dsn和mysql驱动, 创建数据库对象
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	// 删
	if _, err := tx.Exec("delete from member where id = ?", 2); err != nil {
		tx.Rollback()
		return
	}

	//a := 10/0
	//fmt.Println(a)

	// 改
	if _, err := tx.Exec("update member set name = ? where id = ? and name = ? ", "quanquan", 1, "yaomaoze"); err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
}

func main() {

	transaction()
}
