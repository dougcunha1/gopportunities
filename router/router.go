// sub-package router
package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Change debug mode to release mode
	gin.SetMode(gin.ReleaseMode)
	// Initialize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the erver
	router.Run(":8080")
}
