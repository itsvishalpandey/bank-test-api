package model

import "time"

type BankMaster struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Bank     string `gorm:"column:bank" json:"bank"`
	Ifsc     string `gorm:"column:ifsc" json:"ifsc"`
	Micr     string `gorm:"column:micr" json:"micr"`
	Branch   string `gorm:"column:branch" json:"branch"`
	Address  string `gorm:"column:address" json:"address"`
	Contact  string `gorm:"column:contact" json:"contact"`
	City     string `gorm:"column:city" json:"city"`
	District string `gorm:"column:district" json:"district"`
	State    string `gorm:"column:state" json:"state"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
