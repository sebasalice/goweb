/*
Ejercicio 1 - Filtremos nuestro endpoint
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint,
el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Usuarios struct {
	Id             int     `form:"id" json:"id"`
	Edad           int     `form:"edad" json:"edad"`
	Nombre         string  `form:"nombre" json:"nombre"`
	Apellido       string  `form:"apellido" json:"apellido"`
	Email          string  `form:"email" json:"email"`
	Fecha_creacion string  `form:"fecha_creacion" json:"fecha_creacion"`
	Altura         float64 `form:"altura" json:"altura"`
	Activo         bool    `form:"activo" json:"activo"`
}

func GetAll(c *gin.Context) {
	data, err1 := os.ReadFile("./usuarios.json")
	var pe Usuarios
	if err1 == nil {
		if err := json.Unmarshal([]byte(data), &pe); err != nil {
			log.Fatal(err)
		}
		c.JSON(200, pe)
	}
}
func BuscarEmpleado(ctxt *gin.Context) {
	fmt.Println(ctxt.Query("id"))
	var pa Usuarios
	user, ok := pa[ctxt.Query("id")]
	if ok {
		ctxt.String(200, "Información del empleado %s, nombre: %s", ctxt.Query("id"), user)
	} else {
		ctxt.String(404, "Información del empleado ¡No existe!")
	}
}

func main() {
	server := gin.Default()
	server.GET("/usuarios", GetAll)
	server.GET("/usuarioso", BuscarEmpleado)
	server.Run()
	//server.GET("/", PaginaPrincipal)
}

/*
Ejercicio 2 - Get one endpoint
Generar un nuevo endpoint que nos permita traer un solo resultado del array
de la temática. Utilizando path parameters el endpoint debería ser /temática/:id
(recuerda que siempre tiene que ser en plural la temática).
Una vez recibido el id devuelve la posición correspondiente.
Genera una nueva ruta.
Genera un handler para la ruta creada.
Dentro del handler busca el item que necesitas.
Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.
*/
