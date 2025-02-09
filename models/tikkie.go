package models

type Tikkie struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Nickname string `json:"nickname" gorm:"unique"`
	Link     string `json:"link"`
}
