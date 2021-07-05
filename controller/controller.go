package controller

import "runnerserver/controller/router"

func Initialize() {
	router.Initialize()
}

func Run() {
	router.Run()
}
