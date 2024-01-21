// sub-package router
package router

import "github.com/gin-gonic/gin"

// I maisculo indica que a funcao Initialize é exportada automaticamente
func Initialize() {
	// Inicializa o Router utilizando as configuracoes default do gin
	router := gin.Default()
	// Define a rota /ping e o contexto da mesma
	router.GET("/ping", func(c *gin.Context) {
		// Retorna um JSON com a mensagem pong
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Inicia o servidor HTTP e escuta as requisicoes na porta 8080
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
