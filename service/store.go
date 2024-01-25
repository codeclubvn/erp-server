package service

import (
	"context"
	"erp/api/dto/erp"
	"erp/domain"
	"erp/infrastructure"
	"erp/repository"
	"erp/utils"
	"erp/utils/api_errors"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type ERPStoreService interface {
	CreateStoreAndAssignOwner(ctx context.Context, req erpdto.CreateStoreRequest) (*domain.Store, error)
	UpdateStore(ctx context.Context, storeID string, req erpdto.UpdateStoreRequest) (*domain.Store, error)
	ListStore(ctx context.Context, req erpdto.ListStoreRequest) ([]*domain.Store, *int64, error)
	DeleteStore(ctx context.Context, storeID string) error
}

type erpStoreService struct {
	erpStoreRepo repository.ERPStoreRepository
	erpRoleRepo  repository.ERPRoleRepository
	db           *infrastructure.Database
	logger       *zap.Logger
}

func NewStoreService(erpStoreRepo repository.ERPStoreRepository, erpRoleRepo repository.ERPRoleRepository, db *infrastructure.Database, logger *zap.Logger) ERPStoreService {
	return &erpStoreService{
		erpStoreRepo: erpStoreRepo,
		erpRoleRepo:  erpRoleRepo,
		db:           db,
		logger:       logger,
	}
}

func (p *erpStoreService) CreateStoreAndAssignOwner(ctx context.Context, req erpdto.CreateStoreRequest) (*domain.Store, error) {
	ownerID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	store := &domain.Store{
		Name:         req.Name,
		Avatar:       req.Avatar,
		Thumbnail:    req.Thumbnail,
		Bio:          req.Bio,
		Domain:       req.Domain,
		BusinessType: req.BusinessType,
		OpendAt:      req.OpendAt,
		ClosedAt:     req.ClosedAt,
		Phone:        req.Phone,
		Location:     req.Location,
	}

	err = repository.WithTransaction(p.db, func(tx *repository.TX) error {

		store, err := p.erpStoreRepo.Create(tx, ctx, store)
		if err != nil {
			return err
		}

		role, err := p.erpRoleRepo.CreateRole(tx, ctx, "store_owner", []string{}, store.ID.String())
		if err != nil {
			return err
		}

		if err = p.erpRoleRepo.AssignRoleToUser(tx, ctx, ownerID.String(), role.ID.String(), store.ID.String(), true); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return store, nil
}

func (p *erpStoreService) UpdateStore(ctx context.Context, storeID string, req erpdto.UpdateStoreRequest) (*domain.Store, error) {
	u := &domain.Store{
		Name:         req.Name,
		Avatar:       req.Avatar,
		Thumbnail:    req.Thumbnail,
		Bio:          req.Bio,
		Domain:       req.Domain,
		BusinessType: req.BusinessType,
		OpendAt:      req.OpendAt,
		ClosedAt:     req.ClosedAt,
		Phone:        req.Phone,
		Location:     req.Location,
	}

	u.ID = uuid.FromStringOrNil(storeID)
	store, err := p.erpStoreRepo.Update(nil, ctx, u)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (p *erpStoreService) ListStore(ctx context.Context, req erpdto.ListStoreRequest) ([]*domain.Store, *int64, error) {
	userID := utils.GetUserStringIDFromContext(ctx)

	_, err := p.erpStoreRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	stores, total, err := p.erpStoreRepo.List(ctx, req.Search, req.PageOptions, userID)
	if err != nil {
		return nil, nil, err
	}

	return stores, total, nil
}

func (p *erpStoreService) DeleteStore(ctx context.Context, storeID string) error {
	ownerID := utils.GetUserStringIDFromContext(ctx)

	ur, err := p.erpRoleRepo.FindUserRoleByStoreIDAndUserID(ctx, storeID, ownerID)
	if err != nil {
		return err
	}

	if !ur.IsStoreOwner {
		return errors.New(api_errors.ErrPermissionDenied)
	}

	return p.erpStoreRepo.DeleteByID(nil, ctx, storeID)
}
