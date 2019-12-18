package dao

import (
	"errors"
	"fmt"
	"mall/app/service/main/account/internal/model"
)

func (d *Dao) GetAccount(id uint) (*model.Account, error) {
	acc := new(model.Account)
	if d.db.First(acc, id).RecordNotFound() {
		return acc, errors.New("id not found")
	}
	return acc, nil
}

func (d *Dao) GetAccountByName(name string) (*model.Account, error) {
	acc := new(model.Account)
	if d.db.Where("name = ?", name).First(acc).RecordNotFound() {
		return acc, errors.New("id not found")
	}
	return acc, nil
}

func (d *Dao) CreateAccount(acc *model.Account) (id uint, err error) {

	if ok := d.db.NewRecord(acc); ok == false {
		return 0, fmt.Errorf("primary key is not null")
	}
	if err := d.db.Create(acc).Error; err != nil {
		return 0, err
	}
	return acc.ID, nil
}

func (d *Dao) DelAccount(id uint) error {
	acc := new(model.Account)
	acc.ID = id
	// set the delete_at field , use db.Unscoped().Delete() delete the row
	if err := d.db.Delete(acc).Error; err != nil {
		return err
	}
	return nil
}
