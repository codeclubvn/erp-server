package erpdto

type PermissionResponse struct {
	Permission string `json:"permission"`
	ID         string `json:"id"`
}

type CreateRoleRequest struct {
	Name           string   `json:"name" binding:"required"`
	Extends        []string `json:"extends" binding:"required"`
	PersmissionIDs []string `json:"permission_ids" binding:"required"`
}

type ManagePermissionRequest struct {
	PermissionIDs []string `json:"permission_ids"`
	RoleID        string   `json:"role_id"`
}

type CreateEmployeeRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   string `json:"role_id" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}

type AssignRoleRequest struct {
	RoleID string `json:"role_id" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}
