package main

import "time"

type Member struct {
	Id int
	Name string
	Gender string
	Email string
	Phone string
	Password string
	Integral float32
	Birthday string
	CreateTime time.Time
	UpdateTime time.Time
}


