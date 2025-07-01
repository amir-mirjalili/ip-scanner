package repositories

import (
	"github.com/amir-mirjalili/ip-scanner/internal/models"
	"gorm.io/gorm"
)

type ScanRepository interface {
	CreateScan(scan *models.Scan) error
	UpdateScan(scan *models.Scan) error
	CreateScanResults(results []models.ScanResult) error
}

type ScanGormRepository struct {
	DB *gorm.DB
}

func NewGormScanRepository(db *gorm.DB) *ScanGormRepository {
	return &ScanGormRepository{DB: db}
}

func (r *ScanGormRepository) CreateScan(scan *models.Scan) error {
	return r.DB.Create(scan).Error
}

func (r *ScanGormRepository) UpdateScan(scan *models.Scan) error {
	return r.DB.Save(scan).Error
}

func (r *ScanGormRepository) CreateScanResults(results []models.ScanResult) error {
	return r.DB.Create(&results).Error
}
