package adapter

import (
	"github.com/amir-mirjalili/ip-scanner/internal/models"
	"github.com/amir-mirjalili/ip-scanner/internal/repositories"
)

type AssetAdapter struct {
	repo repositories.AssetRepository
}

func NewAssetAdapter(repo *repositories.AssetGormRepository) *AssetAdapter {
	return &AssetAdapter{repo}
}

func (a *AssetAdapter) FindByIP(ip string) (*models.Asset, error) {
	return a.repo.FindAssetByIP(ip)
}

func (a *AssetAdapter) CreateAsset(asset *models.Asset) error {
	return a.repo.CreateAsset(asset)
}
