package dto

import "gitlab.com/investio/backend/api/v1/model"

// SocketDTO - Fund searching
type SocketDTO struct {
	Type  string `json:"type"`
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

type SocketArrayDTO struct {
	Type  string                     `json:"type"`
	Topic string                     `json:"topic"`
	Data  []model.FundSearchResponse `json:"data"`
}
