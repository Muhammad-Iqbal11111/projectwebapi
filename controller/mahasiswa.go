package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator/v10"
	"net/http"
	"projectwebapi/models"
	"time"
	"fmt"
)

type MahasiswaInput struct {
	Id	int `json:"id"  binding:"required,uuid" gorm:"primary_key"`
	Nama string `json:"nama" binding: "required,min=5"`
	Prodi string `json:"prodi" binding:"required"`
	Fakultas string `json:"fakultas" binding:"required"`
	Nim int `json:"nim" binding:"required,gte=6,number"`
	TahunAngkatan int `json:"tahunangkatan" binding:"required,number"`
}

//GET Data
func GetData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"data" : mhs,
		"time" : time.Now(),
	})
}

//POST Data >> Create Data
func CreateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi inputan
	var dataInput MahasiswaInput
	err := c.ShouldBindJSON(&dataInput) 
	if err != nil {
		errorMessages := []string{}
		for _,e := range err.(validator.ValidationErrors){
			switch  e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append (errorMessages, report)
			case "min":
				report := fmt.Sprintf("%s must be more than 5 characters", e.Field())
				errorMessages = append (errorMessages, report)
			case "number":
				report := fmt.Sprintf("%s must be numbers", e.Field())
				errorMessages = append (errorMessages, report)
			case "gte":
				report := fmt.Sprintf("%s must be more than 5", e.Field())
				errorMessages = append (errorMessages, report)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
//	prose input data
	mhs := models.Mahasiswa{
		Id:  dataInput.Id,
		Nama: dataInput.Nama,
		Prodi: dataInput.Prodi,
		Fakultas:dataInput.Fakultas,
		Nim: dataInput.Nim,
		TahunAngkatan: dataInput.TahunAngkatan,
	}

	db.Create(&mhs)

//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Sukses input data",
		"Data" : mhs,
		"time" : time.Now(),
	})
}

//UPDATE Data >> Update Data
func UpdateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mhs models.Mahasiswa
	if err := db.Where("nim =? ", c.Param("nim")).First(&mhs).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Data mahasiswa tidak di temukan",
		})
		return
	}

	//validasi inputan
	var dataInput MahasiswaInput
	err := c.ShouldBindJSON(&dataInput); 
	if err != nil {
		errorMessages := []string{}
		for _,e := range err.(validator.ValidationErrors){
			switch  e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append (errorMessages, report)
			case "min":
				report := fmt.Sprintf("%s must be more than 5 characters", e.Field())
				errorMessages = append (errorMessages, report)
			case "number":
				report := fmt.Sprintf("%s must be numbers", e.Field())
				errorMessages = append (errorMessages, report)
			case "gte":
				report := fmt.Sprintf("%s must be more than 5", e.Field())
				errorMessages = append (errorMessages, report)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	//	prose Ubah data
	db.Model(&mhs).Update(&dataInput)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Sukses ubah data",
		"Data" : mhs,
		"time" : time.Now(),
	})
}

// Delete Data >> Hapus Data
func DeleteData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mhs models.Mahasiswa
	if err := db.Where("nim = ? ", c.Param("nim")).First(&mhs).Error;
		err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Data mahasiswa tidak di temukan",
		})
		return
	}
	//	prose hapus data
	db.Delete(&mhs)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data" : true,
		"Message": "Data berhasil dihapus",
	})
}
