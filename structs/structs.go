package structs

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Nama_Lengkap string `gorm:"type:varchar(300);index"`
	Nama         string `gorm:"type:varchar(300);unique"`
	Username     string `gorm:"type:varchar(50);unique"`
	Password     string `grom:"type:text"`
}

type Penyakit struct {
	gorm.Model

	Nama_Penyakit string `gorm:"type:varchar(300);index"`
	Desc_Penyakit string `grom:"type:text"`
	Id_Penyebab   string `grom:"type:text"`
	Id_Solusi     string `grom:"type:text"`
}

type Penyebab struct {
	Id_Penyebab   string `grom:"type:int(300);primary"`
	Nama_Penyebab string `grom:"type:text"`
	Desc_Penyebab string `grom:"type:text"`
}

type Solusi struct {
	Id_Solusi   string `grom:"type:int(300);primary"`
	Nama_Solusi string `grom:"type:text"`
	Desc_Solusi string `grom:"type:text"`
}
