package models

type Routes struct {
	BaseModel
	Method string `json:"method" gorm:"column:method;type:varchar(50);not null"`
	Path   string `json:"path" gorm:"column:path;type:varchar(255);not null"`
}

type RouteRole struct {
	BaseModel
	RoutePath string `json:"route_path" gorm:"column:route_path;type:varchar(255);not null"`
	RoleID    string `json:"role_id" gorm:"column:role_id;type:varchar(50);not null"`
}
