package dao

import (
	"fmt"
	"mall/app/service/main/account/internal/model"
)

func (d *Dao) GetAccount(id uint) *model.Account {
	acc := new(model.Account)
	d.db.First(acc, id)
	return acc
}

func (d *Dao) CreateAccount(acc *model.Account) (id uint, err error) {
	if ok := d.db.NewRecord(acc); ok == false {
		return 0, fmt.Errorf("primary key is not null")
	}
	d.db.Create(acc)
	return acc.ID, nil
}

func (d *Dao) DelAccount(id uint) error {
	acc := new(model.Account)
	acc.ID = id
	d.db.Delete(acc)
	return nil
}
