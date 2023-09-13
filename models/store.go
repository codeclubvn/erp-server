package models

type Store struct {
	BaseModel
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	Avatar       string `gorm:"type:varchar(255)" json:"avatar"`
	Thumbnail    string `gorm:"type:varchar(255)" json:"thumbnail"`
	Bio          string `gorm:"type:text" json:"bio"`
	Domain       string `gorm:"type:varchar(255);unique" json:"domain"`
	BusinessType string `gorm:"type:varchar(255)" json:"business_type"`
	OpendAt      string `gorm:"type:varchar(255)" json:"opend_at"`
	ClosedAt     string `gorm:"type:varchar(255)" json:"closed_at"`
	Phone        string `gorm:"type:varchar(255)" json:"phone"`
	Location     string `gorm:"type:varchar(255)" json:"location"`
}
