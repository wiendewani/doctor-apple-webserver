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

	Nama_penyakit string `gorm:"type:varchar(300);index"`
	Desc_penyakit string `grom:"type:text"`
	Id_penyebab   string `grom:"type:text"`
	Id_solusi     string `grom:"type:text"`
}

type Penyebab struct {
	Id_penyebab   string `grom:"type:int(300);primary"`
	Nama_penyakit string `grom:"type:text"`
	Desc_penyakit string `grom:"type:text"`
}

type Solusi struct {
	Id_solusi   string `grom:"type:int(300);primary"`
	Nama_solusi string `grom:"type:text"`
	Desc_solusi string `grom:"type:text"`
}
