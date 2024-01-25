package repository

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	"erp/utils"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

type ERPRoleRepository interface {
	UpdateRolePermission(tx *TX, ctx context.Context, roleID string, permissionIDs []string) error
	CreateRole(tx *TX, ctx context.Context, name string, extends []string, storeID string) (*domain.Cashbook, error)
	AssignRevenueToUser(tx *TX, ctx context.Context, userID string, roleID string, storeID string, isStoreOwner bool) error
	FindRevenueByIDs(ids []string) ([]domain.Cashbook, error)
	GetRevenueByRevenueID(ctx context.Context, id string) (*domain.Cashbook, error)
	FindRoleByIDs(ids []string) ([]domain.Role, error)
	AssignRoleToUser(tx *TX, ctx context.Context, userID string, roleID string, storeID string, isStoreOwner bool) error
	FindRoleByID(ctx context.Context, id string) (*domain.Role, error)
	FindUserRoleByStoreIDAndUserID(ctx context.Context, storeID string, userID string) (*domain.UserRole, error)
}

type erpRoleRepository struct {
	db *infrastructure.Database
}

func (e *erpRoleRepository) CreateRole(tx *TX, ctx context.Context, name string, extends []string, storeID string) (*domain.Cashbook, error) {
	//TODO implement me
	panic("implement me")
}

func (e *erpRoleRepository) AssignRevenueToUser(tx *TX, ctx context.Context, userID string, roleID string, storeID string, isStoreOwner bool) error {
	//TODO implement me
	panic("implement me")
}

func (e *erpRoleRepository) FindRevenueByIDs(ids []string) ([]domain.Cashbook, error) {
	//TODO implement me
	panic("implement me")
}

func (e *erpRoleRepository) GetRevenueByRevenueID(ctx context.Context, id string) (*domain.Cashbook, error) {
	//TODO implement me
	panic("implement me")
}

func NewErpRoleRepo(db *infrastructure.Database) ERPRoleRepository {
	return &erpRoleRepository{
		db,
	}
}

//
//func (e *erpRoleRepository) CreateRole(tx *TX, ctx context.Context, name string, extends []string, storeID string) (*domain.Role, error) {
//	tx = GetTX(tx, *e.db)
//
//	sid, _ := uuid.FromString(storeID)
//
//	extendsRole := make([]domain.Role, 0)
//	for _, id := range extends {
//		extendsRole = append(extendsRole, domain.Role{
//			BaseModel: domain.BaseModel{
//				ID: uuid.FromStringOrNil(id),
//			},
//		})
//	}
//
//	role := &domain.Role{
//		Name:    name,
//		StoreID: sid,
//	}
//
//	if err := tx.db.Create(role).Error; err != nil {
//		return nil, errors.Wrap(err, "create role failed")
//	}
//
//	if err := tx.db.Model(role).Omit("Extends.*").Association("Extends").Append(extendsRole); err != nil {
//		return nil, errors.Wrap(err, "create role failed")
//	}
//
//	return role, nil
//}

func (e *erpRoleRepository) UpdateRolePermission(tx *TX, ctx context.Context, roleID string, permissionIDs []string) error {
	tx = GetTX(tx, *e.db)

	err := tx.db.Exec("DELETE FROM role_permissions WHERE role_id = ?", roleID).Error
	if err != nil {
		return errors.Wrap(err, "delete role permission failed")
	}

	var permissions []domain.Permission
	for _, id := range permissionIDs {
		permission := domain.Permission{
			ID: uuid.FromStringOrNil(id),
		}
		permissions = append(permissions, permission)
	}

	return tx.db.Model(&domain.Role{
		BaseModel: domain.BaseModel{
			ID: uuid.FromStringOrNil(roleID),
		},
	}).Omit("Permissions.*").Association("Permissions").Append(permissions)
}

func (e *erpRoleRepository) FindRoleByIDs(ids []string) ([]domain.Role, error) {
	var roles []domain.Role
	if err := e.db.Where("id IN (?)", ids).Find(&roles).Error; err != nil {
		return nil, errors.Wrap(err, "find role by ids failed")
	}

	return roles, nil
}

func (e *erpRoleRepository) AssignRoleToUser(tx *TX, ctx context.Context, userID string, roleID string, storeID string, isStoreOwner bool) error {
	tx = GetTX(tx, *e.db)

	err := tx.db.Debug().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}, {Name: "store_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"role_id": roleID,
		}),
	}).Create(&domain.UserRole{
		IsStoreOwner: isStoreOwner,
		UserID:       uuid.FromStringOrNil(userID),
		RoleID:       uuid.FromStringOrNil(roleID),
		StoreID:      uuid.FromStringOrNil(storeID),
	}).Error

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (e *erpRoleRepository) FindRoleByID(ctx context.Context, id string) (*domain.Role, error) {
	var role domain.Role
	err := e.db.WithContext(ctx).Where("id = ?", id).First(&role).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New("role not found")
		}
		return nil, errors.WithStack(err)
	}

	return &role, nil
}

func (e *erpRoleRepository) FindUserRoleByStoreIDAndUserID(ctx context.Context, storeID string, userID string) (*domain.UserRole, error) {
	var userRole domain.UserRole
	err := e.db.WithContext(ctx).Where("store_id = ? AND user_id = ?", storeID, userID).First(&userRole).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New("user role not found")
		}
		return nil, errors.WithStack(err)
	}

	return &userRole, nil
}
