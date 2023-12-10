package service

import (
	"context"
	"erp/constants"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type TransactionService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateTransactionRequest) (*models.Transaction, error)
	Update(ctx context.Context, req erpdto.UpdateTransactionRequest) (*models.Transaction, error)
	GetList(ctx context.Context, req erpdto.ListTransactionRequest) ([]*models.Transaction, int64, error)
	Delete(ctx context.Context, transactionID string) error
	GetTransactionByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Transaction, error)
	GetOne(ctx context.Context, id string) (*models.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	walletRepo      repository.WalletRepository
	db              *infrastructure.Database
	logger          *zap.Logger
}

func NewTransactionService(transactionRepo repository.TransactionRepository, db *infrastructure.Database, logger *zap.Logger, walletRepo repository.WalletRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		walletRepo:      walletRepo,
		db:              db,
		logger:          logger,
	}
}

func (s *transactionService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateTransactionRequest) (*models.Transaction, error) {
	output := &models.Transaction{}
	err := repository.WithTransaction(s.db, func(tx *repository.TX) error {
		if err := s.CalculateWalletAmount(tx, ctx, req.WalletId.String(), req.Amount, req.Status); err != nil {
			return err
		}

		if err := copier.Copy(&output, &req); err != nil {
			return err
		}
		return s.transactionRepo.Create(tx, ctx, output)
	})
	return output, err
}

func (s *transactionService) Update(ctx context.Context, req erpdto.UpdateTransactionRequest) (*models.Transaction, error) {
	output, err := s.transactionRepo.GetOneById(ctx, req.Id.String())
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
		return s.transactionRepo.Update(tx, ctx, output)
	})

	return output, err
}

func (s *transactionService) GetList(ctx context.Context, req erpdto.ListTransactionRequest) ([]*models.Transaction, int64, error) {
	return s.transactionRepo.GetList(ctx, req)
}

func (s *transactionService) GetOne(ctx context.Context, id string) (*models.Transaction, error) {
	return s.transactionRepo.GetOneById(ctx, id)
}

func (s *transactionService) GetTransactionByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Transaction, error) {
	return s.transactionRepo.GetTransactionByOrderId(tx, ctx, orderId)
}

func (s *transactionService) Delete(ctx context.Context, transactionID string) error {
	output, err := s.transactionRepo.GetOneById(ctx, transactionID)
	if err != nil {
		return err
	}
	err = repository.WithTransaction(s.db, func(tx *repository.TX) error {
		if err = s.CalculateWalletAmount(tx, ctx, output.WalletId.String(), -output.Amount, output.Status); err != nil {
			return err
		}
		return s.transactionRepo.Delete(nil, ctx, transactionID)
	})
	return err
}

func (s *transactionService) CalculateWalletAmount(tx *repository.TX, ctx context.Context, walletId string, amount float64, status string) error {
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
