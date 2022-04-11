package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

var productos []request
var lastID int

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastID++
		req.ID = lastID
		productos = append(productos, req)
		file, err := json.Marshal(productos)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		os.WriteFile("productos.json", file, 0644)
		c.JSON(200, req)
	}
}

func main() {
	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/", Guardar())
	pr.GET("/", func(c *gin.Context) {
		c.JSON(200, productos)
	})

	r.Run()
}
