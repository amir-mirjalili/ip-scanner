package repositories

import (
	"errors"
	"github.com/amir-mirjalili/ip-scanner/internal/models"
	"gorm.io/gorm"
)

type AssetRepository interface {
	FindAssetByIP(ip string) (*models.Asset, error)
	CreateAsset(asset *models.Asset) error
	GetAssetByID(id uint) (*models.Asset, error)
	UpdateAsset(asset *models.Asset) error
	DeleteAsset(id uint) error
	ListAssets() ([]models.Asset, error)
}

type AssetGormRepository struct {
	DB *gorm.DB
}

func NewAssetGormRepository(db *gorm.DB) *AssetGormRepository {
	return &AssetGormRepository{DB: db}
}

func (r *AssetGormRepository) FindAssetByIP(ip string) (*models.Asset, error) {
	var asset models.Asset
	err := r.DB.Where("ip_address = ?", ip).First(&asset).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &asset, err
}

func (r *AssetGormRepository) CreateAsset(asset *models.Asset) error {
	return r.DB.Create(asset).Error
}

func (r *AssetGormRepository) GetAssetByID(id uint) (*models.Asset, error) {
	var asset models.Asset
	err := r.DB.First(&asset, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &asset, err
}

func (r *AssetGormRepository) UpdateAsset(asset *models.Asset) error {
	return r.DB.Save(asset).Error
}

func (r *AssetGormRepository) DeleteAsset(id uint) error {
	return r.DB.Delete(&models.Asset{}, id).Error
}

func (r *AssetGormRepository) ListAssets() ([]models.Asset, error) {
	var assets []models.Asset
	err := r.DB.Find(&assets).Error
	return assets, err
}
