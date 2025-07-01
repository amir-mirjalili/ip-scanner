package services

import (
	"fmt"
	"github.com/amir-mirjalili/ip-scanner/internal/scanner"
	"time"

	"github.com/amir-mirjalili/ip-scanner/internal/adapter"
	"github.com/amir-mirjalili/ip-scanner/internal/models"
	"github.com/amir-mirjalili/ip-scanner/internal/repositories"
)

type ScanService struct {
	ScanRepo     repositories.ScanRepository
	AssetAdapter *adapter.AssetAdapter
}

func NewScanService(scanRepo repositories.ScanRepository, assetAdapter *adapter.AssetAdapter) *ScanService {
	return &ScanService{
		ScanRepo:     scanRepo,
		AssetAdapter: assetAdapter,
	}
}

func (s *ScanService) RunAndSaveScan(cidr string) (*models.Scan, error) {
	scan := &models.Scan{
		CIDR:      cidr,
		StartedAt: time.Now(),
		Status:    models.ScanStatusInProgress,
	}
	if err := s.ScanRepo.CreateScan(scan); err != nil {
		return nil, fmt.Errorf("failed to create scan session: %w", err)
	}

	rawResults, err := scanner.ScanNetwork(cidr)
	if err != nil {
		scan.Status = models.ScanStatusFailed
		_ = s.ScanRepo.UpdateScan(scan)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	var scanResults []models.ScanResult
	for _, r := range rawResults {
		existingAsset, err := s.AssetAdapter.FindByIP(r.IP)
		if err != nil {
			return nil, fmt.Errorf("error finding asset: %w", err)
		}
		if existingAsset == nil {
			asset := &models.Asset{
				IPAddress:  r.IP,
				MACAddress: r.MAC,
				Hostname:   r.Hostname,
				OS:         r.OS,
			}
			if err := s.AssetAdapter.CreateAsset(asset); err != nil {
				return nil, fmt.Errorf("error creating asset: %w", err)
			}
			existingAsset = asset
		}

		scanResults = append(scanResults, models.ScanResult{
			ScanID:     scan.ID,
			AssetID:    existingAsset.ID,
			DetectedAt: time.Now(),
		})
	}

	if len(scanResults) > 0 {
		if err := s.ScanRepo.CreateScanResults(scanResults); err != nil {
			return nil, fmt.Errorf("failed to save scan results: %w", err)
		}
	}

	finishedAt := time.Now()
	scan.Status = models.ScanStatusCompleted
	scan.FinishedAt = &finishedAt
	if err := s.ScanRepo.UpdateScan(scan); err != nil {
		return nil, fmt.Errorf("failed to update scan session: %w", err)
	}

	return scan, nil
}

func (s *ScanService) GetScanByID(scanId uint) (*models.Scan, error) {
	return s.ScanRepo.GetAllScans(scanId)
}
