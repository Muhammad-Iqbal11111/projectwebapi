package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectwebapi/controller"
	"projectwebapi/models"
)

func main() {

	r := gin.Default()
	//Models
	db := models.SetUpModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "golang web api",
		})
	})
	//GET All Data
	r.GET("/mahasiswa", controller.GetData)
	//POST Data >> Create Data
	r.POST("/mahasiswa", controller.CreateData)
	//Update Data >> Update Data
	r.PUT("/mahasiswa/:nim", controller.UpdateData)
	//Delete Data >> Delete data
	r.DELETE("/mahasiswa/:nim", controller.DeleteData)
		r.Run()
}
