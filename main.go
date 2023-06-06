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
	port := os.Getenv("PORT")
	e := routes.Init()
	e.Start(":" + port)
}
