package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	STRING_VACIO           = ""
	INT_ZERO               = 0
	TRANSACCIONES_FILENAME = "./transacciones.json"
)

type GenericResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Transaccion struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo_transaccion" binding:"required"`
	Moneda            string  `json:"moneda" binding:"required"`
	Monto             float64 `json:"monto" binding:"required"`
	Emisor            string  `json:"emisor" binding:"required"`
	Receptor          string  `json:"receptor" binding:"required"`
	FechaTransaccion  string  `json:"fecha_transaccion" binding:"required"`
}

func GetTransacciones() ([]Transaccion, error) {
	jsonData, err := os.ReadFile(TRANSACCIONES_FILENAME)
	if err != nil {
		return nil, errors.New("Archivo no encontrado")
	}
	var sl []Transaccion
	serr := json.Unmarshal((jsonData), &sl)
	if serr != nil {
		return nil, errors.New("Archivo con formato no valido")
	}
	return sl, nil
}

func GetMaxIdTransaccion(transacciones []Transaccion) int {
	var maxId int
	for _, transaccion := range transacciones {
		if maxId < transaccion.Id {
			maxId = transaccion.Id
		}
	}
	return maxId
}

func FiltrarTransacciones(transaccion, filter Transaccion) bool {
	return (filter.Id == INT_ZERO || transaccion.Id == filter.Id) &&
		(filter.CodigoTransaccion == STRING_VACIO || transaccion.CodigoTransaccion == filter.CodigoTransaccion) &&
		(filter.Moneda == STRING_VACIO || transaccion.Moneda == filter.Moneda) &&
		(filter.Monto == INT_ZERO || transaccion.Monto == filter.Monto) &&
		(filter.Emisor == STRING_VACIO || transaccion.Emisor == filter.Emisor) &&
		(filter.Receptor == STRING_VACIO || transaccion.Receptor == filter.Receptor) &&
		(filter.FechaTransaccion == STRING_VACIO || transaccion.FechaTransaccion == filter.FechaTransaccion)
}

func GetAll(c *gin.Context) {
	sl, err := GetTransacciones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, GenericResponse{Message: "Respuesta exitosa", Data: sl})
}

func GetFilterTransacciones(c *gin.Context) {
	var filter Transaccion
	filter.Id, _ = strconv.Atoi(c.Query("id"))
	filter.CodigoTransaccion = c.Query("codigo_transaccion")
	filter.Moneda = c.Query("moneda")
	filter.Monto, _ = strconv.ParseFloat(c.Query("monto"), 64)
	filter.Emisor = c.Query("emisor")
	filter.Receptor = c.Query("receptor")
	filter.FechaTransaccion = c.Query("fecha_transaccion")

	transacciones, err := GetTransacciones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Message: err.Error()})
		return
	}

	var transaccionesFiltradas []Transaccion

	for _, transaccion := range transacciones {
		if FiltrarTransacciones(transaccion, filter) {
			transaccionesFiltradas = append(transaccionesFiltradas, transaccion)
		}
	}

	if len(transaccionesFiltradas) == INT_ZERO {
		c.JSON(http.StatusNotFound, GenericResponse{Message: "Ninguna transaccion cumple el criterio"})
		return
	}

	c.JSON(http.StatusOK, GenericResponse{Message: "Respuesta filtrada con exito", Data: transaccionesFiltradas})
}

func GetTransaccion(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("Id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Message: err.Error()})
		return
	}

	transacciones, err := GetTransacciones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Message: err.Error()})
		return
	}

	for _, transaccion := range transacciones {
		if transaccion.Id == idParam {
			c.JSON(http.StatusOK, GenericResponse{Message: "Respuesta filtrada con exito", Data: transaccion})
			return
		}
	}

	c.JSON(http.StatusNotFound, GenericResponse{Message: "Transaccion no encontrada"})
}

func GetTheme() gin.HandlerFunc {
	now := time.Now().Local()
	limit := time.Date(2022, 4, 4, 14, 8, 0, 0, time.Local)
	if limit.After(now) {
		return func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"theme": "dark",
			})
		}
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"theme": "light",
		})
	}
}

func SaveTransaccion(c *gin.Context) {
	var transaccion Transaccion
	if err := c.ShouldBindJSON(&transaccion); err != nil {
		c.JSON(http.StatusBadRequest, GenericResponse{Message: err.Error()})
		return
	}
	lastId := GetMaxIdTransaccion(transaccionesList) + 1
	transaccion.Id = lastId
	transaccionesList = append(transaccionesList, transaccion)
	content, err := json.Marshal(transaccionesList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericResponse{Message: "Error al almacenar la nueva transaccion"})
		return
	}
	os.WriteFile(TRANSACCIONES_FILENAME, content, 0644)
	c.JSON(http.StatusOK, GenericResponse{Message: "Transaccion almacenada con exito", Data: transaccion})
}

var transaccionesList []Transaccion

func main() {
	router := gin.Default()

	transaccionesList, _ = GetTransacciones()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hola brandon"})
	})

	transaccionesGroup := router.Group("/transacciones")
	transaccionesGroup.GET("", GetAll)
	transaccionesGroup.GET("/", GetFilterTransacciones)
	transaccionesGroup.GET("/:Id", GetTransaccion)
	transaccionesGroup.POST("", SaveTransaccion)

	router.GET("/theme", GetTheme())

	router.Run()
}
