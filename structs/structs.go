package structs

import "github.com/jinzhu/gorm"

type Disease struct {
	gorm.Model

	Nama_Penyakit string `gorm:"type:varchar(300);index"`
	Desc_Penyakit string `grom:"type:text"`
	Desc_Penyebab string `grom:"type:text"`
	Desc_Solusi   string `grom:"type:text"`
}

type Analyzer struct {
	Image string `gorm:"type:blob;primary_key"`
}
