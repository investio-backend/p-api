package model

type Amc struct {
	ID     uint32 `json:"id"`
	NameEn string `json:"name_en"`
	NameTh string `json:"name_th"`
	Funds  []Fund
}
