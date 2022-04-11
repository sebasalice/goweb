//Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
//Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.

// 1) Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
// 2) Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.

//Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
//Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
//Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
//Genera un handler para el endpoint llamado “GetAll”.
//Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.

package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transacciones struct {
	Id                 int     `json:"id"`
	Codigo_transaccion string  `json:"codigo_transaccion"`
	Moneda             string  `json:"moneda"`
	Monto              float64 `json:"monto"`
	Emisor             string  `json:"emisor"`
	Receptor           string  `json:"receptor"`
	Fecha              string  `json:"fecha"`
}

var trans []transacciones
var lastID int

func GetAll(c *gin.Context) {
	var trans []transacciones
	jsonData, err := os.ReadFile("./transacciones.json")
	if err != nil {
		log.Fatal(err)
	}
	if err1 := json.Unmarshal([]byte(jsonData), &trans); err1 != nil {
		log.Fatal(err1)
	}

	c.JSON(200, trans)

}
func transaccionesById(c *gin.Context) {
	for _, value := range trans {
		idString := c.Param("id")         //captar el param del context y meterlo en una variable
		id, err := strconv.Atoi(idString) //convertirlo a string para que json pueda leerlo
		if err != nil {                   //errores
			log.Fatal("error al convertir id")
		}
		if value.Id == id { //trans es una estructura asi que hay que acceder a sus valores con .VALOR
			c.JSON(200, value)
		}

	}
}
func QueryParam(c *gin.Context) {
	jsonData, err := os.ReadFile("./transacciones.json")
	if err != nil {
		log.Fatal(err)
	}
	if err1 := json.Unmarshal([]byte(jsonData), &trans); err1 != nil {
		log.Fatal(err1)
	}

	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal("error fatal")
	}
	for _, value := range trans {
		if value.Id == id {
			c.JSON(200, value)
		}

	}

}
func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req transacciones
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++
		req.Id = lastID
		if req.Codigo_transaccion == "" {
			c.JSON(400, gin.H{
				"error": "Codigo_transaccion es un campo requerido",
			})
			return
		}
		if req.Moneda == "" {
			c.JSON(400, gin.H{
				"error": "Moneda es un campo requerido",
			})
			return
		}
		if req.Monto == 0 {
			c.JSON(400, gin.H{
				"error": "Monto es un campo requerido",
			})
			return
		}
		if req.Emisor == "" {
			c.JSON(400, gin.H{
				"error": "Emisor es un campo requerido",
			})
			return
		}
		if req.Receptor == "" {
			c.JSON(400, gin.H{
				"error": "Receptor es un campo requerido",
			})
			return
		}
		if req.Fecha == "" {
			c.JSON(400, gin.H{
				"error": "Fecha es un campo requerido",
			})
			return
		}
		c.JSON(200, req)

	}
}

func main() {
	router := gin.Default()
	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, Sebastian",
		})
	})
	router.GET("/transacciones", GetAll)
	router.GET("/transacciones/:id", transaccionesById)
	router.GET("/transaccionesQuery", QueryParam)
	pr := router.Group("/productos")
	pr.POST("/", Guardar())
	router.Run()

}

/*1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).*/
