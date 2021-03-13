package model

type Amc struct {
	// ID        uint32 `json:"id"`
	AmcCode   string `json:"amc_code"`
	AmcNameEn string `json:"name_en"`
	AmcNameTh string `json:"name_th"`
}

func (Amc) TableName() string {
	return "amc"
}
