package repository

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	"erp/handler/dto/request"
	"erp/utils"
	"erp/utils/api_errors"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ERPStoreRepository interface {
	Create(tx *TX, ctx context.Context, store *domain.Store) (*domain.Store, error)
	Update(tx *TX, ctx context.Context, store *domain.Store) (*domain.Store, error)
	FindByID(ctx context.Context, id string) (*domain.Store, error)
	List(ctx context.Context, search string, o request.PageOptions, userID string) ([]*domain.Store, *int64, error)
	DeleteByID(tx *TX, ctx context.Context, id string) error
}

type erpStoreRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewERPStoreRepository(db *infrastructure.Database, logger *zap.Logger) ERPStoreRepository {
	return &erpStoreRepository{
		db:     db,
		logger: logger,
	}
}

func (p *erpStoreRepository) Create(tx *TX, ctx context.Context, store *domain.Store) (*domain.Store, error) {
	tx = GetTX(tx, *p.db)

	if err := tx.db.WithContext(ctx).Create(store).Error; err != nil {
		return nil, errors.Wrap(err, "create store failed")
	}

	return store, nil
}

func (p *erpStoreRepository) Update(tx *TX, ctx context.Context, store *domain.Store) (*domain.Store, error) {
	tx = GetTX(tx, *p.db)

	if err := tx.db.WithContext(ctx).Save(store).Error; err != nil {
		return nil, errors.Wrap(err, "UpdateById store failed")
	}

	return store, nil
}

func (p *erpStoreRepository) FindByID(ctx context.Context, id string) (*domain.Store, error) {
	var store domain.Store
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(&store).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrStoreNotFound)
		}
		return nil, errors.Wrap(err, "Find store failed")
	}

	return &store, nil
}

func (p *erpStoreRepository) List(ctx context.Context, search string, o request.PageOptions, userID string) ([]*domain.Store, *int64, error) {
	var stores []*domain.Store
	var total int64 = 0

	q := p.db.WithContext(ctx).Model(&domain.Store{})

	if search != "" {
		q = q.Where("name LIKE ?", "%"+search+"%")
	}

	q.Order("created_at DESC")

	if err := utils.QueryPagination(q, o, &stores).Count(&total).Error(); err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return stores, &total, nil
}

func (p *erpStoreRepository) DeleteByID(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *p.db)

	if err := tx.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Store{}).Error; err != nil {
		return errors.Wrap(err, "Delete store failed")
	}

	return nil
}
