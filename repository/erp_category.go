package repository

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	models "erp/models"
	"erp/utils"
	"github.com/pkg/errors"
)

type CategoryRepository interface {
	Create(ctx context.Context, Category *models.Category) (err error)
	Update(ctx context.Context, category *models.Category) (err error)
	Delete(ctx context.Context, id string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.Category, err error)
	GetList(ctx context.Context, category erpdto.GetListCategoryRequest) (res []*models.CategoryResult, total int64, err error)
}

type categoryRepo struct {
	db *infrastructure.Database
}

func NewCategoryRepository(db *infrastructure.Database) CategoryRepository {
	return &categoryRepo{db}
}

func (u *categoryRepo) Create(ctx context.Context, category *models.Category) (err error) {
	//currentUID, err := utils.GetUserUUIDFromContext(ctx)
	//if err != nil {
	//	return err
	//}
	//category.UpdaterID = currentUID

	err = u.db.Create(&category).Error
	return errors.Wrap(err, "create category failed")
}

func (u *categoryRepo) Update(ctx context.Context, category *models.Category) (err error) {
	//currentUID, err := utils.GetUserUUIDFromContext(ctx)
	//if err != nil {
	//	return err
	//}
	//category.UpdaterID = currentUID

	err = u.db.Save(&category).Error
	return errors.Wrap(err, "update category failed")
}

func (u *categoryRepo) Delete(ctx context.Context, id string) (err error) {
	if err := u.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		return errors.Wrap(err, "Delete product failed")
	}
	return nil
}

func (u *categoryRepo) GetOneByID(ctx context.Context, id string) (res *models.Category, err error) {
	err = u.db.Where("id = ?", id).First(&res).Error
	return res, errors.Wrap(err, "get category by id failed")
}

func (u *categoryRepo) GetList(ctx context.Context, req erpdto.GetListCategoryRequest) (res []*models.CategoryResult, total int64, err error) {
	var output []*models.CategoryResult
	query := u.db.Table("categories").Debug()
	if req.Search != "" {
		query = query.Where("unaccent(name) ilike ?", "%"+req.Search+"%")
	}

	query.Count(&total)

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &output).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return output, total, nil
}
