package main

import (
	"github.com/hooneun/if-commerce/config"
	"github.com/hooneun/if-commerce/server"
	"log"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.RunAPI(":1323"))
}
