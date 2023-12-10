package service

import (
	"context"
	"erp/api_errors"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type WalletService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateWalletRequest) (*models.Wallet, error)
	Update(ctx context.Context, req erpdto.UpdateWalletRequest) (*models.Wallet, error)
	GetList(ctx context.Context, req erpdto.ListWalletRequest) ([]*models.Wallet, int64, error)
	Delete(ctx context.Context, walletID string) error
	GetOne(ctx context.Context, id string) (*models.Wallet, error)
}

type walletService struct {
	walletRepo repository.WalletRepository
	db         *infrastructure.Database
	logger     *zap.Logger
}

func NewWalletService(walletRepo repository.WalletRepository, db *infrastructure.Database, logger *zap.Logger) WalletService {
	return &walletService{
		walletRepo: walletRepo,
		db:         db,
		logger:     logger,
	}
}

func (p *walletService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateWalletRequest) (*models.Wallet, error) {
	// check if wallet name is already exist
	if _, err := p.walletRepo.GetOneByName(ctx, req.Name); err == nil {
		return nil, errors.New(api_errors.ErrWalletNameAlreadyExist)
	}

	output := &models.Wallet{}
	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	// if isDefault is true, then set all other wallet to false
	if req.IsDefault {
		if err := p.walletRepo.SetAllWalletToFalse(ctx); err != nil {
			return nil, err
		}
	}

	if err := p.walletRepo.Create(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *walletService) Update(ctx context.Context, req erpdto.UpdateWalletRequest) (*models.Wallet, error) {
	output, err := p.walletRepo.GetOneById(ctx, req.Id.String())
	if err != nil {
		return nil, err
	}

	// if isDefault is true, then set all other wallet to false
	if req.IsDefault {
		if err := p.walletRepo.SetAllWalletToFalse(ctx); err != nil {
			return nil, err
		}
	}

	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.walletRepo.Update(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *walletService) GetList(ctx context.Context, req erpdto.ListWalletRequest) ([]*models.Wallet, int64, error) {
	return p.walletRepo.GetList(ctx, req)
}

func (p *walletService) GetOne(ctx context.Context, id string) (*models.Wallet, error) {
	return p.walletRepo.GetOneById(ctx, id)
}

func (p *walletService) Delete(ctx context.Context, walletID string) error {
	return p.walletRepo.Delete(nil, ctx, walletID)
}
