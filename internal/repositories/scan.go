package repositories

import (
	"github.com/amir-mirjalili/ip-scanner/internal/models"
	"gorm.io/gorm"
)

type ScanRepository interface {
	CreateScan(scan *models.Scan) error
	UpdateScan(scan *models.Scan) error
	CreateScanResults(results []models.ScanResult) error
	GetScansResults(scanId uint) (*models.Scan, error)
	GetAllScansResults() (*[]models.Scan, error)
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

func (r *ScanGormRepository) GetScansResults(scanID uint) (*models.Scan, error) {
	var scan models.Scan
	err := r.DB.Preload("Results.Asset").First(&scan, scanID).Error
	if err != nil {
		return nil, err
	}
	return &scan, nil
}

func (r *ScanGormRepository) GetAllScansResults() (*[]models.Scan, error) {
	var scans []models.Scan
	err := r.DB.Preload("Results.Asset").Find(&scans).Error
	if err != nil {
		return nil, err
	}
	return &scans, nil
}
