package models

type Role struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;type:varchar(200);not null"`
	Key  string `json:"key" gorm:"column:key;type:varchar(50);not null"`
}
