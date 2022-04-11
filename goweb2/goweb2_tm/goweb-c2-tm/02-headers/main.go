package main

import "github.com/gin-gonic/gin"

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

func Saludar() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Hola!",
		})
	}
}

func main() {
	r := gin.Default()
	pr := r.Group("/hola")
	pr.POST("/", Saludar())
	r.Run()
}
