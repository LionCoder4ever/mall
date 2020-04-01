package dao

import (
	"errors"
	"mall/app/internal/model"
)

func (d *Dao) CreateAccount(acc *model.Account) (id int64, err error) {
	// if not found, create the record
	if err := d.db.FirstOrCreate(&model.Account{UId: acc.UId}, acc).Error; err != nil {
		return 0, err
	}
	return acc.UId, nil
}

func (d *Dao) ReadAccount(uid int64) (*model.Account, error) {
	acc := new(model.Account)
	if d.db.Where("uid = ?", uid).First(acc).RecordNotFound() {
		return acc, errors.New("account not found")
	}
	acc.AccountPrivacy = model.AccountPrivacy{}
	return acc, nil
}

func (d *Dao) ReadAccountByPhone(phone string) (*model.Account, error) {
	acc := new(model.Account)
	if d.db.Where("phone = ?", phone).First(acc).RecordNotFound() {
		return acc, errors.New("id not found")
	}
	return acc, nil
}

func (d *Dao) DeleteAccount(uid int64) error {
	// set the delete_at field , use db.Unscoped().Delete() delete the row
	if err := d.db.Where("uid = ?", uid).Delete(&model.Account{}).Error; err != nil {
		return err
	}
	return nil
}
