package main

import (
	"os"
	"prakerja/eventmanagement/configs"
	"prakerja/eventmanagement/routes"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDatabase()
}

func main() {
	port := os.Getenv("WEB_PORT")
	e := routes.Init()
	e.Start(":" + port)
}
