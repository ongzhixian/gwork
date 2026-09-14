package main

import (
	"log"
	"readyperfectly/application"
)

func main() {
	appConfig, err := application.GetApplicationConfiguration()
	if err != nil {
		log.Fatal("Error getting application configuration on start.", err)
	}

	log.Printf("All done. %s", appConfig)
}
