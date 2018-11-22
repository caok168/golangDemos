package repositories

import (
	"github.com/jinzhu/gorm"
	"golangDemos/gormtest/models"
)

type ConfigRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) *ConfigRepository {
	db.AutoMigrate(&models.Config{})
	return &ConfigRepository{db: db}
}

func (r *ConfigRepository) List(offset, limit int, maps interface{}) ([]*models.Config, uint, error) {
	var configs []*models.Config
	err := r.db.Where(maps).Offset(offset).Limit(limit).Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	var total uint
	err = r.db.Where(maps).Model(&models.Config{}).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return configs, total, nil
}
