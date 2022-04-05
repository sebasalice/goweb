package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type usuarios struct {
	Id             int     `json:"id"`
	Edad           int     `json:"edad"`
	Nombre         string  `json:"nombre"`
	Apellido       string  `json:"apellido"`
	Email          string  `json:"email"`
	Fecha_creacion string  `json:"fecha_creacion"`
	Altura         float64 `json:"altura"`
	Activo         bool    `json:"activo"`
}

func GetAll(c *gin.Context) {
	data, err1 := os.ReadFile("./usuarios.json")
	var pe usuarios
	if err1 == nil {
		if err := json.Unmarshal([]byte(data), &pe); err != nil {
			log.Fatal(err)
		}
		c.JSON(200, pe)
	}
}

func main() {
	router := gin.Default()
	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Sebastian",
		})
	})

	router.GET("/usuarios", GetAll)
	router.Run()
}

//Ejercicio 2 - Hola {nombre}

//Crea dentro de la carpeta go-web un archivo llamado main.go
//Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
//Pegale al endpoint para corroborar que la respuesta sea la correcta.

//Ejercicio 3 - Listar Entidad

//ya  habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
//Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
//Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
//Genera un handler para el endpoint llamado “GetAll”.
//Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint
