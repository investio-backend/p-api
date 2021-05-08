package model

type AimcCat struct {
	ID        uint32 `gorm:"primary_key" json:"-"`
	CatID     string `json:"cid"`
	CatNameEn string `json:"name_en"`
	CatNameTh string `json:"name_th"`
	BrdCatID  uint32 `json:"-"`
}

func (AimcCat) TableName() string {
	return "aimc_cat"
}

type AimcBrdCat struct {
	ID         uint32    `gorm:"primary_key" json:"-"`
	BcatID     string    `json:"bid"`
	BcatNameEn string    `json:"name_en"`
	BcatNameTh string    `json:"name_th"`
	Cats       []AimcCat `json:"cats" gorm:"foreignKey:BrdCatID;references:ID"`
}

func (AimcBrdCat) TableName() string {
	return "aimc_brd_cat"
}
