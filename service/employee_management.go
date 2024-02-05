package service

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	"erp/handler/dto/erp"
	"erp/repository"
	"erp/utils"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type (
	ERPEmployeeManagementService interface {
		ListPermission() (data []erpdto.PermissionResponse, total *int64, err error)
		CreateRole(ctx context.Context, req erpdto.CreateRoleRequest) (id string, err error)
		// AssignRole(ctx context.Context, req erpdto.AssignRoleRequest) (err error)
		CreateEmployee(ctx context.Context, req erpdto.CreateEmployeeRequest) (id string, err error)
	}

	erpEmployeeManagementService struct {
		erpPermissionRepo repository.ErpPermissionRepo
		erpRoleRepo       repository.ERPRoleRepository
		userRepo          repository.UserRepository
		db                *infrastructure.Database
	}
)

func NewERPEmployeeManagementService(erpPermissionRepo repository.ErpPermissionRepo, erpRoleRepo repository.ERPRoleRepository, userRepo repository.UserRepository, db *infrastructure.Database) ERPEmployeeManagementService {
	return &erpEmployeeManagementService{
		erpPermissionRepo,
		erpRoleRepo,
		userRepo,
		db,
	}
}

func (e *erpEmployeeManagementService) ListPermission() (res []erpdto.PermissionResponse, total *int64, err error) {
	permissions, total, err := e.erpPermissionRepo.List()
	if err != nil {
		return nil, nil, err
	}

	for _, permission := range permissions {
		res = append(res, erpdto.PermissionResponse{
			Permission: permission.Name,
			ID:         permission.ID.String(),
		})
	}

	return
}

func (e *erpEmployeeManagementService) CreateRole(ctx context.Context, req erpdto.CreateRoleRequest) (id string, err error) {
	storeID := utils.GetStoreIDFromContext(ctx)

	err = repository.WithTransaction(e.db, func(tx *repository.TX) error {
		role, err := e.erpRoleRepo.CreateRole(tx, ctx, req.Name, req.Extends, storeID)
		if err != nil {
			return err
		}

		if err := e.erpRoleRepo.UpdateRolePermission(tx, ctx, role.ID.String(), req.PersmissionIDs); err != nil {
			return err
		}

		id = role.ID.String()
		return nil
	})

	if err != nil {
		return "", err
	}

	return
}

func (e *erpEmployeeManagementService) CreateEmployee(ctx context.Context, req erpdto.CreateEmployeeRequest) (id string, err error) {
	var user *domain.User

	err = repository.WithTransaction(e.db, func(tx *repository.TX) error {
		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return errors.WithStack(err)
		}

		user, err = e.userRepo.Create(tx, ctx, domain.User{
			FullName: req.FullName,
			Email:    req.Email,
			Password: string(encryptedPassword),
		})
		if err != nil {
			return err
		}

		role, err := e.erpRoleRepo.FindRoleByID(ctx, req.RoleID)
		if err != nil {
			return err
		}

		storeID := utils.GetStoreIDFromContext(ctx)

		if err := e.erpRoleRepo.AssignRoleToUser(tx, ctx, user.ID.String(), role.ID.String(), storeID, false); err != nil {
			return err
		}

		return nil
	})

	return user.ID.String(), err
}
