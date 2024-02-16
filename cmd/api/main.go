package main

import (
	"log"
	config "sample/pkg/config"

	"sample/pkg/di"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("error loading the env file")
	// }

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("can't start server ", err)
	} else {
		server.Start()
	}
}
