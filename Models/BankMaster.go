package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankMaster struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Bank     string    `gorm:"column:bank" json:"bank"`
	Ifsc     string    `gorm:"column:ifsc" json:"ifsc"`
	Micr     string    `gorm:"column:micr" json:"micr"`
	Branch   string    `gorm:"column:branch" json:"branch"`
	Address  string    `gorm:"column:address" json:"address"`
	Contact  string    `gorm:"column:contact" json:"contact"`
	City     string    `gorm:"column:city" json:"city"`
	District string    `gorm:"column:district" json:"district"`
	State    string    `gorm:"column:state" json:"state"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BankMaster) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
