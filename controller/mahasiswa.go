package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"projectwebapi/models"
	"time"
)

type MahasiswaInput struct {
	Id	int `json:"id"`
	Nama string `json:"nama"`
	Prodi string `json:"prodi"`
	Fakultas string `json:"fakultas"`
	Nim int `json:"nim"`
	TahunAngkatan int `json:"tahunangkatan"`

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
	if err := c.ShouldBindJSON(&dataInput); err != nil {
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
	if err := c.ShouldBindJSON(&dataInput); err != nil {
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
	})
}
