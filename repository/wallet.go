package repository

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	erpdto "erp/handler/dto/finance"
	"erp/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type WalletRepository interface {
	Create(tx *TX, ctx context.Context, trans *domain.Wallet) error
	Update(tx *TX, ctx context.Context, trans *domain.Wallet) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetOneById(ctx context.Context, id string) (*domain.Wallet, error)
	GetOneByName(ctx context.Context, name string) (*domain.Wallet, error)
	GetList(ctx context.Context, req erpdto.ListWalletRequest) (res []*domain.Wallet, total int64, err error)
	SetAllWalletToFalse(ctx context.Context) error
}

type walletRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewWalletRepository(db *infrastructure.Database, logger *zap.Logger) WalletRepository {
	return &walletRepo{
		db:     db,
		logger: logger,
	}
}

func (r *walletRepo) Create(tx *TX, ctx context.Context, trans *domain.Wallet) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Create(trans).Error
}

func (r *walletRepo) GetOneById(ctx context.Context, id string) (*domain.Wallet, error) {
	trans := &domain.Wallet{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(trans).Error
	return trans, err
}

func (r *walletRepo) GetOneByName(ctx context.Context, name string) (*domain.Wallet, error) {
	trans := &domain.Wallet{}
	err := r.db.WithContext(ctx).Where("name = ?", name).First(trans).Error
	return trans, err
}

func (r *walletRepo) GetList(ctx context.Context, req erpdto.ListWalletRequest) (res []*domain.Wallet, total int64, err error) {
	query := r.db.Model(&domain.Wallet{})
	if req.Search != "" {
		query = query.Where("name ilike ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res).
		Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}

func (r *walletRepo) Update(tx *TX, ctx context.Context, trans *domain.Wallet) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", trans.ID).Save(trans).Error
}

func (r *walletRepo) Delete(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Wallet{}).Error
}

func (r *walletRepo) SetAllWalletToFalse(ctx context.Context) error {
	userId := ctx.Value("x-user-id").(string)
	return r.db.WithContext(ctx).Model(&domain.Wallet{}).Where("updater_id = ?", userId).Update("is_default", false).Error
}
