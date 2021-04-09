package main

import (
	"hospital/model"
	"hospital/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
