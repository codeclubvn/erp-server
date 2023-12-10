package models

import (
	"github.com/lib/pq"
)

type Customer struct {
	BaseModel
	FullName       string         `json:"full_name" gorm:"column:full_name;type:varchar(100);not null"`
	Gender         string         `json:"gender" gorm:"column:gender;type:varchar(12)"`
	Age            int            `json:"age" gorm:"column:age;type:int"`
	AddressStrings pq.StringArray `json:"address_strings" gorm:"column:address_strings;type:text[]"`
	PhoneNumber    string         `json:"phone_number" gorm:"column:phone_number;type:varchar(15)"`
	Email          string         `json:"email" gorm:"column:email;type:varchar(100);not null"`
}

func (Customer) TableName() string {
	return "customers"
}
