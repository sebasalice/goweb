package main

import "github.com/gin-gonic/gin"

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cant"`
	Precio   float64 `json:"precio"`
}

func main() {
	r := gin.Default()
	r.POST("/productos", func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		req.ID = 4
		c.JSON(200, req)
	})

	r.Run()
}
