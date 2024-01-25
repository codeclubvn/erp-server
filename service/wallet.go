package service

import (
	"context"
	erpdto "erp/api/dto/finance"
	"erp/domain"
	"erp/infrastructure"
	"erp/repository"
	"erp/utils/api_errors"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type WalletService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateWalletRequest) (*domain.Wallet, error)
	Update(ctx context.Context, req erpdto.UpdateWalletRequest) (*domain.Wallet, error)
	GetList(ctx context.Context, req erpdto.ListWalletRequest) ([]*domain.Wallet, int64, error)
	Delete(ctx context.Context, walletID string) error
	GetOne(ctx context.Context, id string) (*domain.Wallet, error)
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

func (p *walletService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateWalletRequest) (*domain.Wallet, error) {
	// check if wallet name is already exist
	if _, err := p.walletRepo.GetOneByName(ctx, req.Name); err == nil {
		return nil, errors.New(api_errors.ErrWalletNameAlreadyExist)
	}

	output := &domain.Wallet{}
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

func (p *walletService) Update(ctx context.Context, req erpdto.UpdateWalletRequest) (*domain.Wallet, error) {
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

func (p *walletService) GetList(ctx context.Context, req erpdto.ListWalletRequest) ([]*domain.Wallet, int64, error) {
	return p.walletRepo.GetList(ctx, req)
}

func (p *walletService) GetOne(ctx context.Context, id string) (*domain.Wallet, error) {
	return p.walletRepo.GetOneById(ctx, id)
}

func (p *walletService) Delete(ctx context.Context, walletID string) error {
	return p.walletRepo.Delete(nil, ctx, walletID)
}
