package main

import (
	"w8s/configs"
	"w8s/database"
	"w8s/routers"
)

func main() {
	config := configs.LoadEnv()
	database.StartDB(&config)
	routers.StartServer().Run(config.Port)
}
