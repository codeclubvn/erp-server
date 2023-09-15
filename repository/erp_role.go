package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

type ERPRoleRepository interface {
	UpdateRolePermission(tx *TX, ctx context.Context, roleID string, permissionIDs []string) error
	CreateRole(tx *TX, ctx context.Context, name string, extends []string, storeID string) (*models.Role, error)
	AssignRoleToUser(tx *TX, ctx context.Context, userID string, roleID string, storeID string, isStoreOwner bool) error
	FindRoleByIDs(ids []string) ([]models.Role, error)
	FindRoleByID(ctx context.Context, id string) (*models.Role, error)
	FindUserRoleByStoreIDAndUserID(ctx context.Context, storeID string, userID string) (*models.UserRole, error)
}

type erpRoleRepository struct {
	db *infrastructure.Database
}

func NewErpRoleRepo(db *infrastructure.Database) ERPRoleRepository {
	return &erpRoleRepository{
		db,
	}
}

func (e *erpRoleRepository) CreateRole(tx *TX, ctx context.Context, name string, extends []string, storeID string) (*models.Role, error) {
	if tx == nil {
		tx = &TX{db: *e.db}
	}

	updaterID := utils.GetUserStringIDFromContext(ctx)
	sid, _ := uuid.FromString(storeID)

	extendsRole := make([]models.Role, 0)
	for _, id := range extends {
		extendsRole = append(extendsRole, models.Role{
			BaseModel: models.BaseModel{
				ID: uuid.FromStringOrNil(id),
			},
		})
	}

	role := &models.Role{
		Name:    name,
		StoreID: sid,
		BaseModel: models.BaseModel{
			UpdaterID: uuid.FromStringOrNil(updaterID),
		},
	}

	if err := tx.db.Create(role).Error; err != nil {
		return nil, errors.Wrap(err, "create role failed")
	}

	if err := tx.db.Model(role).Omit("Extends.*").Association("Extends").Append(extendsRole); err != nil {
		return nil, errors.Wrap(err, "create role failed")
	}

	return role, nil
}

func (e *erpRoleRepository) UpdateRolePermission(tx *TX, ctx context.Context, roleID string, permissionIDs []string) error {
	if tx == nil {
		tx = &TX{db: *e.db}
	}
	err := tx.db.Exec("DELETE FROM role_permissions WHERE role_id = ?", roleID).Error
	if err != nil {
		return errors.Wrap(err, "delete role permission failed")
	}

	var permissions []models.Permission
	for _, id := range permissionIDs {
		permission := models.Permission{
			ID: uuid.FromStringOrNil(id),
		}
		permissions = append(permissions, permission)
	}

	return tx.db.Model(&models.Role{
		BaseModel: models.BaseModel{
			ID: uuid.FromStringOrNil(roleID),
		},
	}).Omit("Permissions.*").Association("Permissions").Append(permissions)
}

func (e *erpRoleRepository) FindRoleByIDs(ids []string) ([]models.Role, error) {
	var roles []models.Role
	if err := e.db.Where("id IN (?)", ids).Find(&roles).Error; err != nil {
		return nil, errors.Wrap(err, "find role by ids failed")
	}

	return roles, nil
}

func (e *erpRoleRepository) AssignRoleToUser(tx *TX, ctx context.Context, userID string, roleID string, storeID string, isStoreOwner bool) error {
	currentUserID := utils.GetUserStringIDFromContext(ctx)
	if tx == nil {
		tx = &TX{db: *e.db}
	}

	updaterID := uuid.FromStringOrNil(currentUserID)

	err := tx.db.Debug().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}, {Name: "store_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"role_id": roleID,
		}),
	}).Create(&models.UserRole{
		UpdaterID:    updaterID,
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

func (e *erpRoleRepository) FindRoleByID(ctx context.Context, id string) (*models.Role, error) {
	var role models.Role
	err := e.db.WithContext(ctx).Where("id = ?", id).First(&role).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New("role not found")
		}
		return nil, errors.WithStack(err)
	}

	return &role, nil
}

func (e *erpRoleRepository) FindUserRoleByStoreIDAndUserID(ctx context.Context, storeID string, userID string) (*models.UserRole, error) {
	var userRole models.UserRole
	err := e.db.WithContext(ctx).Where("store_id = ? AND user_id = ?", storeID, userID).First(&userRole).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New("user role not found")
		}
		return nil, errors.WithStack(err)
	}

	return &userRole, nil
}
