package service

import (
	"stock/models"

	postgres "stock/db"
)

// GroupBarangService ...
type GroupBarangService struct {
}

// Add ,,,
func (groupBarangService *GroupBarangService) Add(groupBarang models.GroupBarang) models.GroupBarang {
	db := postgres.GetDbCon()

	db.Create(&groupBarang)
	return groupBarang
}

// FindByID Param id
func (groupBarangService *GroupBarangService) FindByID(id int32) models.GroupBarang {
	db := postgres.GetDbCon()

	var groupBarang models.GroupBarang
	db.Where("id = ?", id).Find(&groupBarang)
	return groupBarang
}

// FindByCode Param id
func (groupBarangService *GroupBarangService) FindByCode(code string) models.GroupBarang {
	db := postgres.GetDbCon()

	var groupBarang models.GroupBarang
	db.Where("code = ?", code).Find(&groupBarang)
	return groupBarang
}
