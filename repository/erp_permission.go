package repository

import (
	"erp/cmd/infrastructure"
	"erp/domain"
	"erp/handler/dto/request"
	"erp/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ErpPermissionRepo interface {
	List() ([]*domain.Permission, *int64, error)
}

type erpPermissionRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewErpPermissionRepo(db *infrastructure.Database, logger *zap.Logger) ErpPermissionRepo {
	return &erpPermissionRepo{db, logger}
}

func (e *erpPermissionRepo) List() ([]*domain.Permission, *int64, error) {
	var total int64 = 0
	var res []*domain.Permission
	query := e.db.Model(&domain.Permission{})
	err := utils.QueryPagination(query, request.PageOptions{
		Limit: 1000,
		Page:  1,
	}, &res).Count(&total).Error()

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return res, &total, nil
}
