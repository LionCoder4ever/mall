package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"mall/app/internal/model"
	"mall/library/log"
)

func (d *Dao) CreateAccount(acc *model.Account) (id int64, err error) {
	// if not found, create the record
	if err = d.db.Where("uid = ?", acc.UId).First(acc).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			if err = d.db.Create(acc).Error; err != nil {
				log.Logger.Error("create account failed cause %s", err.Error())
				return
			}
			return acc.UId, nil
		}
		log.Logger.Error("create account failed when check is exist cause %s", err.Error())
		return
	}
	return
}

func (d *Dao) ReadAccount(uid int64) (*model.Account, error) {
	acc := new(model.Account)
	if err := d.db.Where("uid = ?", uid).First(acc).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		log.Logger.Error("read account failed cause %s", err.Error())
		return nil, err
	}
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
