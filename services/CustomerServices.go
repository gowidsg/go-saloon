package services

import (
	"github.com/gowidsg/go-saloon/models"
	"github.com/gowidsg/go-saloon/repository"
	"github.com/jinzhu/gorm"
)

type CustomerServices struct {
	DB         *gorm.DB
	Repository *repository.RepositorySRV
}

func NewCustomerServices(db *gorm.DB, repo *repository.RepositorySRV) *CustomerServices {
	db.AutoMigrate(models.Customer{})
	return &CustomerServices{DB: db, Repository: repo}
}

func (custsrv *CustomerServices) AddCustomer(cust *models.Customer) error {
	ufw := repository.NewUnitOfWork(custsrv.DB, false)
	err := custsrv.Repository.Add(ufw, cust)
	if err != nil {
		ufw.RollingBack()
		return err
	}
	ufw.Committing()
	return err
}

func (custsrv *CustomerServices) GetAllCustomer(cust *[]models.Customer) error {
	ufw := repository.NewUnitOfWork(custsrv.DB, true)
	return custsrv.Repository.GetCustomer(ufw, cust)

}
