package repository

import (
	"errors"

	"github.com/gowidsg/go-saloon/models"
)

type Repository interface {
	Add(ufw *UnitOfWork, in interface{}) error
}

type RepositorySRV struct {
}

func NewRepositorySRV() *RepositorySRV {
	return &RepositorySRV{}
}

func (srv *RepositorySRV) Add(ufw *UnitOfWork, in interface{}) error {
	return ufw.DB.Create(in).Error

}
func (srv *RepositorySRV) GetAll(ufw *UnitOfWork, in interface{}) error {
	c := in.(*[]models.Customer)
	return ufw.DB.Find(&c).Error
}
func (srv *RepositorySRV) GetByUserID(ufw *UnitOfWork, in interface{}) error {
	c := in.(*models.Customer)
	return ufw.DB.Where("user_id = ?", &c.UserID).Find(&c).Error
}
func (srv *RepositorySRV) Delete(ufw *UnitOfWork, in interface{}) error {
	c := in.(*models.Customer)
	res := ufw.DB.Where("user_id = ?", &c.UserID).Delete(&c)
	if res.RowsAffected == 0 {
		return errors.New("customer not found")
	}
	return res.Error
}

func (srv *RepositorySRV) Update(ufw *UnitOfWork, in interface{}) error {
	c := in.(*models.Customer)
	res := ufw.DB.Model(models.Customer{}).Where("user_id = ?", &c.UserID).Update(&c)
	if res.RowsAffected == 0 {
		return errors.New("customer not found")
	}
	return res.Error
}
