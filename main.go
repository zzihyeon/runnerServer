package main

import (
	"runnerserver/config"
	"runnerserver/controller"
	"runnerserver/infrastructure"
	"runnerserver/service"
)

func main() {
	initialize()
	Run()
	for {
	}
}

func initialize() {
	config.Initialize()
	controller.Initialize()
	service.Initialize()
	infrastructure.Initialize()
}

func Run() {
	go controller.Run()
	go service.Run()
	go infrastructure.Run()
}
