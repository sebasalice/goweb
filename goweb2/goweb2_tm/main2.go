package main

import "github.com/gin-gonic/gin"

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

var products []request
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
		c.JSON(200, req)
		products = append(products, req)

	}
}

func main() {
	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/", Guardar())
	r.Run()
}
