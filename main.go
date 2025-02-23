package main

import (
	"github.com/KSKigor/api_rest_go/config"
	"github.com/KSKigor/api_rest_go/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
