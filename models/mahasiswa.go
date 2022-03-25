package models

type Mahasiswa struct {
	Id	int `json:"id" gorm:"primary_key"`
	Nama string `json:"nama"`
	Prodi string `json:"prodi"`
	Fakultas string `json:"fakultas"`
	Nim int `json:"nim"`
	TahunAngkatan int `json:"tahunangkatan"`
}
