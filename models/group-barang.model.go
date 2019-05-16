package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// GroupBarang ...
type GroupBarang struct {
	ID        int32     `gorm:"column(id)"`
	Code      string    `gorm:"column(code)"`
	Name      string    `gorm:"column(name)"`
	UpdatedBy string    `gorm:"column(updated_by)"`
	UpdatedAt time.Time `gorm:"column(updated_at)"`
}

// TableName ...
func (GroupBarang) TableName() string {
	return "group_barang"
}

// BeforeCreate ...
func (GroupBarang) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("updated_at", time.Now().UTC())
	return nil
}
