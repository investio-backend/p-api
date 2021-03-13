package model

type AimcCat struct {
	CatID     string `json:"id"`
	CatNameEn string `json:"name_en"`
	CatNameTh string `json:"name_th"`
}

func (AimcCat) TableName() string {
	return "aimc_cat"
}

type AimcBrdCat struct {
	CatID     string `json:"id"`
	CatNameEn string `json:"name_en"`
	CatNameTh string `json:"name_th"`
}

func (AimcBrdCat) TableName() string {
	return "aimc_brd_cat"
}
