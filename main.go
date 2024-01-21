package main

// /router refere-se ao sub-pacote router
import (
	"github.com/dougcunha1/gopportunities/config"
	"github.com/dougcunha1/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()
}
