package test

import (
	"erp/models"
	erpservice "erp/service"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestListPermission(t *testing.T) {
	expectedPermissionCount := 3
	mockRepo := &MockPermissionRepo{}

	service := erpservice.NewERPEmployeeManagementService(mockRepo, nil, nil, nil)

	res, total, err := service.ListPermission()

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

	if len(res) != expectedPermissionCount {
		t.Errorf("Expected %d permissions, but got %d", expectedPermissionCount, len(res))
	}

	if total == nil {
		t.Error("Expected non-nil total, but got nil")
	}
}

type MockPermissionRepo struct {
}

func (*MockPermissionRepo) List() ([]*models.Permission, *int64, error) {
	total := int64(3)
	return []*models.Permission{
		{ID: uuid.NewV4(), Name: "permission1"},
		{ID: uuid.NewV4(), Name: "permission2"},
		{ID: uuid.NewV4(), Name: "permission3"},
	}, &total, nil
}
