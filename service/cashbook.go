package service

import (
	"context"
	erpdto "erp/api/dto/finance"
	"erp/constants"
	"erp/domain"
	"erp/infrastructure"
	"erp/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type CashbookService interface {
	Create(ctx context.Context, req erpdto.CreateCashbookRequest) (*domain.Cashbook, error)
	Update(ctx context.Context, req erpdto.UpdateCashbookRequest) (*domain.Cashbook, error)
	GetList(ctx context.Context, req erpdto.ListCashbookRequest) ([]*domain.Cashbook, int64, error)
	GetListDebt(ctx context.Context, req erpdto.ListCashbookRequest) ([]*domain.Cashbook, int64, error)
	Delete(ctx context.Context, transactionID string) error
	GetCashbookByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*domain.Cashbook, error)
	GetOne(ctx context.Context, id string) (*domain.Cashbook, error)
}

type cashbookService struct {
	cashbookRepo repository.CashbookRepository
	walletRepo   repository.WalletRepository
	db           *infrastructure.Database
	logger       *zap.Logger
}

func NewCashbookService(cashbookRepo repository.CashbookRepository, db *infrastructure.Database, logger *zap.Logger, walletRepo repository.WalletRepository) CashbookService {
	return &cashbookService{
		cashbookRepo: cashbookRepo,
		walletRepo:   walletRepo,
		db:           db,
		logger:       logger,
	}
}

func (s *cashbookService) Create(ctx context.Context, req erpdto.CreateCashbookRequest) (*domain.Cashbook, error) {
	output := &domain.Cashbook{}
	err := repository.WithTransaction(s.db, func(tx *repository.TX) error {
		if err := s.CalculateWalletAmount(tx, ctx, req.WalletId.String(), req.Amount, req.Status); err != nil {
			return err
		}

		if err := copier.Copy(&output, &req); err != nil {
			return err
		}
		return s.cashbookRepo.Create(tx, ctx, output)
	})
	return output, err
}

func (s *cashbookService) Update(ctx context.Context, req erpdto.UpdateCashbookRequest) (*domain.Cashbook, error) {
	output, err := s.cashbookRepo.GetOneById(ctx, req.Id.String())
	if err != nil {
		return nil, err
	}

	err = repository.WithTransaction(s.db, func(tx *repository.TX) error {
		amount := req.Amount - output.Amount
		if err = s.CalculateWalletAmount(tx, ctx, req.WalletId.String(), amount, req.Status); err != nil {
			return err
		}

		if err = copier.Copy(&output, &req); err != nil {
			return err
		}
		return s.cashbookRepo.Update(tx, ctx, output)
	})

	return output, err
}

func (s *cashbookService) GetList(ctx context.Context, req erpdto.ListCashbookRequest) ([]*domain.Cashbook, int64, error) {
	return s.cashbookRepo.GetList(ctx, req)
}

func (s *cashbookService) GetListDebt(ctx context.Context, req erpdto.ListCashbookRequest) ([]*domain.Cashbook, int64, error) {
	return s.cashbookRepo.GetListDebt(ctx, req)
}

func (s *cashbookService) GetOne(ctx context.Context, id string) (*domain.Cashbook, error) {
	return s.cashbookRepo.GetOneById(ctx, id)
}

func (s *cashbookService) GetCashbookByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*domain.Cashbook, error) {
	return s.cashbookRepo.GetCashbookByOrderId(tx, ctx, orderId)
}

func (s *cashbookService) Delete(ctx context.Context, transactionID string) error {
	output, err := s.cashbookRepo.GetOneById(ctx, transactionID)
	if err != nil {
		return err
	}
	err = repository.WithTransaction(s.db, func(tx *repository.TX) error {
		if err = s.CalculateWalletAmount(tx, ctx, output.WalletId.String(), -output.Amount, output.Status); err != nil {
			return err
		}
		return s.cashbookRepo.Delete(nil, ctx, transactionID)
	})
	return err
}

func (s *cashbookService) CalculateWalletAmount(tx *repository.TX, ctx context.Context, walletId string, amount float64, status string) error {
	if walletId == "" || amount == 0 {
		return nil
	}

	wallet, err := s.walletRepo.GetOneById(ctx, walletId)
	if err != nil {
		return err
	}

	if status == constants.StatusIn {
		wallet.Amount += amount
	}
	if status == constants.StatusOut {
		wallet.Amount -= amount
	}

	if err := s.walletRepo.Update(tx, ctx, wallet); err != nil {
		return err
	}
	return nil
}
