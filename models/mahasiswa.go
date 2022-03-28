package models

type Mahasiswa struct {
	Id	int `json:"id"  binding:"required,uuid" gorm:"primary_key"`
	Nama string `json:"nama" binding: "required,min=5"`
	Prodi string `json:"prodi" binding:"required"`
	Fakultas string `json:"fakultas" binding:"required"`
	Nim int `json:"nim" binding:"required,gte=6"`
	TahunAngkatan int `json:"tahunangkatan" binding:"required"`
}
