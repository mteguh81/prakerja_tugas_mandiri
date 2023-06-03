package main

import (
	"prakerja/eventmanagement/configs"
	"prakerja/eventmanagement/routes"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDatabase()
}

func main() {
	e := routes.Init()
	e.Start(":8000")
}
