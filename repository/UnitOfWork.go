package repository

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

type UnitOfWork struct {
	DB        *gorm.DB
	ReadOnly  bool
	Committed bool
}

func NewUnitOfWork(db *gorm.DB, readonly bool) *UnitOfWork {
	if readonly {
		return &UnitOfWork{DB: db.New(), ReadOnly: true, Committed: false}
	}
	return &UnitOfWork{DB: db.New().Begin(), ReadOnly: false, Committed: false}
}

func (uow *UnitOfWork) RollingBack() {
	if !uow.ReadOnly && !uow.Committed {
		uow.DB.Rollback()
	}
}

func (uow *UnitOfWork) Committing() {
	if !uow.ReadOnly && !uow.Committed {
		uow.DB.Commit()
	}
}
