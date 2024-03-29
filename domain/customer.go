package domain

type Customer struct {
	BaseModel
	FullName    string `json:"full_name" gorm:"column:full_name;type:varchar(100);not null"`
	Gender      string `json:"gender" gorm:"column:gender;type:varchar(12)"`
	Age         int    `json:"age" gorm:"column:age;type:int"`
	Address     string `json:"address" gorm:"column:address;type:text"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;type:varchar(15)"`
	Email       string `json:"email" gorm:"column:email;type:varchar(100);not null"`
}

type CustomerDetailResponse struct {
	BaseModel
	FullName    string  `json:"full_name" gorm:"column:full_name;type:varchar(100);not null"`
	Gender      string  `json:"gender" gorm:"column:gender;type:varchar(12)"`
	Age         int     `json:"age" gorm:"column:age;type:int"`
	Address     string  `json:"address" gorm:"column:address;type:text"`
	PhoneNumber string  `json:"phone_number" gorm:"column:phone_number;type:varchar(15)"`
	Email       string  `json:"email" gorm:"column:email;type:varchar(100);not null"`
	TotalOrder  int     `json:"total_order"`
	TotalPaid   float64 `json:"total_paid"`
	TotalDebt   float64 `json:"total_debt"`
}

func (Customer) TableName() string {
	return "customers"
}
