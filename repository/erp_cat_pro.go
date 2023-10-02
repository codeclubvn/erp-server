package repository

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	models "erp/models"
	"erp/utils"
	"github.com/pkg/errors"
)

type CategoryProductRepository interface {
	Create(ctx context.Context, CategoryProduct *models.CategoryProduct) (err error)
	Update(ctx context.Context, categoryProduct *models.CategoryProduct) (err error)
	Delete(ctx context.Context, id string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.CategoryProduct, err error)
	GetList(ctx context.Context, req erpdto.GetListCatProRequest) (res []*models.CategoryProduct, total *int64, err error)
}

type catProRepo struct {
	db *infrastructure.Database
}

func NewCategoryProductRepository(db *infrastructure.Database) CategoryProductRepository {
	return &catProRepo{db}
}

func (u *catProRepo) Create(ctx context.Context, categoryProduct *models.CategoryProduct) (err error) {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return err
	}
	categoryProduct.UpdaterID = currentUID

	err = u.db.Create(&categoryProduct).Error
	return errors.Wrap(err, "Create category_product failed")
}

func (u *catProRepo) Update(ctx context.Context, categoryProduct *models.CategoryProduct) (err error) {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return err
	}
	categoryProduct.UpdaterID = currentUID

	err = u.db.Updates(&categoryProduct).Error
	return errors.Wrap(err, "Update category_product failed")
}

func (u *catProRepo) Delete(ctx context.Context, id string) (err error) {
	if err := u.db.WithContext(ctx).Where("id = ?", id).Delete(&models.CategoryProduct{}).Error; err != nil {
		return errors.Wrap(err, "Delete category_product failed")
	}
	return nil
}

func (u *catProRepo) GetOneByID(ctx context.Context, id string) (res *models.CategoryProduct, err error) {
	err = u.db.Where("id = ?", id).First(&res).Error
	return res, errors.Wrap(err, "GetOneByID category_product failed")
}

func (u *catProRepo) GetList(ctx context.Context, req erpdto.GetListCatProRequest) (res []*models.CategoryProduct, total *int64, err error) {
	query := u.db.Model(&models.CategoryProduct{})
	if req.Search != "" {
		query = query.Where("name like ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res).Count(total).Error(); err != nil {
		return nil, nil, errors.Wrap(err, "GetList category_product failed")
	}

	return res, total, nil
}
