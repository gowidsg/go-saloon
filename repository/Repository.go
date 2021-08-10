package repository

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
