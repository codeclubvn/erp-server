package repository

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	"go.uber.org/zap"
)

type ResetPasswordToken interface {
	Create(tx *TX, ctx context.Context, input *domain.ResetPasswordToken) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetOneById(ctx context.Context, id string) (*domain.ResetPasswordToken, error)
}

type resetPasswordToken struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewResetPassToken(db *infrastructure.Database, logger *zap.Logger) ResetPasswordToken {
	return &resetPasswordToken{
		db:     db,
		logger: logger,
	}
}

func (r *resetPasswordToken) Create(tx *TX, ctx context.Context, input *domain.ResetPasswordToken) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Create(input).Error
}

func (r *resetPasswordToken) GetOneById(ctx context.Context, id string) (*domain.ResetPasswordToken, error) {
	output := &domain.ResetPasswordToken{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(output).Error
	return output, err
}

func (r *resetPasswordToken) Delete(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.ResetPasswordToken{}).Error
}
