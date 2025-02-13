package main

import (
	"log"

	"github.com/savanyv/Golang-Findest/internal/app"
	"github.com/savanyv/Golang-Findest/internal/config"
)

func main() {
	config := config.LoadConfig()

	server := app.NewServer(config)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}