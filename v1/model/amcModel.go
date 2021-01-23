package model

import "github.com/jinzhu/gorm"

type Amc struct {
	gorm.Model
	ID     uint32 `json:"id"`
	NameEn string `json:"name_en"`
	NameTh string `json:"name_th"`
	Funds  []Fund
}
