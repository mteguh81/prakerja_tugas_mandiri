package models

type BaseResp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
