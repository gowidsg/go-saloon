package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	UserID    string `gorm:"type:varchar(20);unique_key;not_null;" json:"id"`
	Fname     string `gorm:"type:varchar(20);" json:"fname"`
	Lname     string `gorm:"type:varchar(20);" json:"lname"`
	ContactNo string `gorm:"type:varchar(10);" json:"contact"`
	Email     string `gorm:"type:varchar(50);" json:"mail"`
	DOB       string `gorm:"type:varchar(10);" json:"dob"`
}

// func NewCustomerwithUserID() *Customer {
// 	return

// }

// func NewCustomerwithID() *Customer {
// 	return

// }

func (customer *Customer) GetFname() string {
	return customer.Fname
}

func (customer *Customer) GetLname() string {
	return customer.Lname
}

func (customer *Customer) GetContactNo() string {
	return customer.ContactNo
}

func (customer *Customer) GetEmail() string {
	return customer.Email
}

func (customer *Customer) GetDOB() string {
	return customer.DOB
}
