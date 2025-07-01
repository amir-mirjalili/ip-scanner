package services

import (
	"github.com/amir-mirjalili/ip-scanner/internal/models"
	"github.com/amir-mirjalili/ip-scanner/internal/repositories"
)

type AssetService struct {
	repo repositories.AssetRepository
}

func NewAssetService(repo repositories.AssetRepository) *AssetService {
	return &AssetService{repo: repo}
}

func (s *AssetService) FindByIP(ip string) (*models.Asset, error) {
	return s.repo.FindAssetByIP(ip)
}

func (s *AssetService) Create(asset *models.Asset) error {
	return s.repo.CreateAsset(asset)
}

func (s *AssetService) GetByID(id uint) (*models.Asset, error) {
	return s.repo.GetAssetByID(id)
}

func (s *AssetService) Update(asset *models.Asset) error {
	return s.repo.UpdateAsset(asset)
}

func (s *AssetService) Delete(id uint) error {
	return s.repo.DeleteAsset(id)
}

func (s *AssetService) List() ([]models.Asset, error) {
	return s.repo.ListAssets()
}
