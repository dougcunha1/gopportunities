// sub-package router
package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the erver
	router.Run(":8080")
}
