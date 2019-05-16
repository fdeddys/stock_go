package models

import "time"

// Barang ...
type Barang struct {
	ID        int32     `gorm:"primary_key:AUTO_INCREMENT"`
	Code      string    `gorm:"index:code"`
	Name      string    `gorm:"size:255"`
	UpdatedBy string    `gorm:"column(updated_by)"`
	UpdatedAt time.Time `gorm:"column(updated_at)"`
}

// TableName ...
func (barang *Barang) TableName() string {
	return "barang"
}
