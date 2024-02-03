package repository

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	"erp/utils/api_errors"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type FileRepository interface {
	Create(ctx context.Context, file *domain.File) (err error)
	Update(ctx context.Context, file *domain.File) (err error)
	GetOneById(ctx context.Context, id string) (res *domain.File, err error)
	DeleteById(ctx context.Context, id string) (err error)
}

type fileRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewFileRepository(db *infrastructure.Database, logger *zap.Logger) FileRepository {
	return &fileRepository{
		db:     db,
		logger: logger,
	}
}

func (r *fileRepository) Create(ctx context.Context, file *domain.File) (err error) {
	err = r.db.Create(&file).Error
	return errors.Wrap(err, "create file failed")
}

func (r *fileRepository) Update(ctx context.Context, file *domain.File) (err error) {
	err = r.db.Updates(&file).Error
	return errors.Wrap(err, "update file failed")
}

func (r *fileRepository) GetOneById(ctx context.Context, id string) (res *domain.File, err error) {
	var file domain.File
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&file).Error; err != nil {
		return nil, errors.New(api_errors.ErrFileNotFound)
	}
	return &file, nil
}

func (r *fileRepository) DeleteById(ctx context.Context, id string) (err error) {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.File{}).Error; err != nil {
		return errors.Wrap(err, "Delete file failed")
	}
	return nil
}
