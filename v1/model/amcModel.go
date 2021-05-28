package model

type Amc struct {
	ID        uint32 `json:"amc_id"`
	AmcCode   string `json:"amc_code"`
	AmcNameEn string `json:"amc_name_en"`
	AmcNameTh string `json:"amc_name_th"`
}

func (Amc) TableName() string {
	return "amc"
}
