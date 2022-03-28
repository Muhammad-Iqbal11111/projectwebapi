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

	v1 := r.Group("/v1")
	db := models.SetUpModels()
	v1.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "Selamat Berhasil GET",
		})
	})
	//GET All Data
	v1.GET("/mahasiswa", controller.GetData)
	//POST Data >> Create Data
	v1.POST("/mahasiswa", controller.CreateData)
	//Update Data >> Update Data
	v1.PUT("/mahasiswa/:nim", controller.UpdateData)
	//Delete Data >> Delete data
	v1.DELETE("/mahasiswa/:nim", controller.DeleteData)
		r.Run()
}
