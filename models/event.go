package models

type Event struct {
	Id        int    `gorm:"primaryKey autoIncrement" json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Eventdate string `json:"eventdate"`
}
