package structs

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Nama_Lengkap string `gorm:"type:varchar(300);index"`
	Nama         string `gorm:"type:varchar(300);unique"`
	Username     string `gorm:"type:varchar(50);unique"`
	password     string `grom:"type:text"`
}

type penyakit struct {
	gorm.Model

	nama_penyakit string `gorm:"type:varchar(300);index"`
	desc_penyakit string `grom:"type:text"`
	id_penyebab   string `grom:"type:text"`
	id_solusi     string `grom:"type:text"`
	gambar        string `grom:"type:text"`
}

type penyebab struct {
	id_penyebab   string `grom:"type:int(300);primary"`
	nama_penyakit string `grom:"type:text"`
	desc_penyakit string `grom:"type:text"`
}

type solusi struct {
	id_solusi   string `grom:"type:int(300);primary"`
	nama_solusi string `grom:"type:text"`
	desc_solusi string `grom:"type:text"`
}
