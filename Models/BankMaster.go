package model

import "gorm.io/gorm"

type BankMaster struct {
	Bank   string `gorm:"column:bank" json:"bank"`
	Ifsc   string `gorm:"column:ifsc" json:"ifsc"`
	Micr   string `gorm:"column:micr" json:"micr"`
	Branch string `gorm:"column:branch" json:"branch"`
	gorm.Model
}
