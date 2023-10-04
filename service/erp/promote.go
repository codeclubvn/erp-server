package erpservice

import (
	"context"
	"erp/api_errors"
	"erp/constants"
	erpdto "erp/dto/erp"
	"erp/models"
	"erp/repository"
	"erp/utils"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"log"
)

type IPromoteService interface {
	CreateFlow(ctx context.Context, req erpdto.CreatePromoteRequest) (*models.Promote, error)
	GetPromoteById(ctx context.Context, id string) (*models.Promote, error)
	CountCustomerUsePromote(ctx context.Context, customerId, code string) (int64, error)
	UpdateQuantityUse(ctx context.Context, code string, quantityUse int) error
	CreatePromoteUse(ctx context.Context, req erpdto.CreatePromoteUseRequest) error
	GetPromoteByCode(ctx context.Context, code string) (*models.Promote, error)
}

type promoteService struct {
	promoteRepo repository.IPromoteRepo
}

func NewPromoteService(promoteRepo repository.IPromoteRepo) IPromoteService {
	return &promoteService{
		promoteRepo: promoteRepo,
	}
}

func (s *promoteService) CreateFlow(ctx context.Context, req erpdto.CreatePromoteRequest) (*models.Promote, error) {
	promote, err := s.promoteRepo.GetPromoteByCode(ctx, req.Code)
	if err != nil {
		if !utils.ErrNoRows(err) {
			return nil, errors.Wrap(err, "Find promote failed")
		}
	}

	if promote != nil {
		return nil, errors.New(api_errors.ErrPromoteCodeExist)
	}

	if req.PromoteType == constants.TypePercent {
		if req.DiscountValue > 100 {
			return nil, errors.New(api_errors.ErrValidation)
		}
	}

	promote = &models.Promote{}
	if err := copier.Copy(promote, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	if err := s.promoteRepo.Create(ctx, promote); err != nil {
		return nil, err
	}

	return promote, nil
}

func (s *promoteService) GetPromoteById(ctx context.Context, id string) (*models.Promote, error) {
	return s.promoteRepo.GetPromoteById(ctx, id)
}

func (s *promoteService) GetPromoteByCode(ctx context.Context, code string) (*models.Promote, error) {
	return s.promoteRepo.GetPromoteByCode(ctx, code)
}

func (s *promoteService) CountCustomerUsePromote(ctx context.Context, customerId, code string) (int64, error) {
	return s.promoteRepo.CountCustomerUsePromote(ctx, customerId, code)
}
func (s *promoteService) UpdateQuantityUse(ctx context.Context, code string, quantityUse int) error {
	return s.promoteRepo.UpdateQuantityUse(ctx, code, quantityUse)
}

func (s *promoteService) CreatePromoteUse(ctx context.Context, req erpdto.CreatePromoteUseRequest) error {
	promoteUse := &models.PromoteUse{}
	if err := copier.Copy(&promoteUse, &req); err != nil {
		return err
	}
	return s.promoteRepo.CreatePromoteUse(ctx, promoteUse)
}
