package structs

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Nama_Lengkap string `gorm:"type:varchar(300);index"`
}
