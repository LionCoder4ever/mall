package account

import (
	"errors"
	"github.com/jinzhu/gorm"
	"mall/internal/pkg/dao"
	"mall/internal/pkg/log"
)

func CreateAccount(acc *Account) (id int64, err error) {
	// if not found, create the record
	if err = dao.Db.Where("uid = ?", acc.UId).First(acc).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			if err = dao.Db.Create(acc).Error; err != nil {
				log.Logger.Error("create account failed cause %s", err.Error())
				return
			}
			return acc.UId, nil
		}
		log.Logger.Error("create account failed when check is exit cause %s", err.Error())
		return
	}
	return
}

func ReadAccount(uid int64) (*Account, error) {
	acc := new(Account)
	if err := dao.Db.Where("uid = ?", uid).First(acc).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		log.Logger.Error("read account failed cause %s", err.Error())
		return nil, err
	}
	return acc, nil
}

func ReadAccountByPhone(phone string) (*Account, error) {
	acc := new(Account)
	if dao.Db.Where("phone = ?", phone).First(acc).RecordNotFound() {
		return acc, errors.New("id not found")
	}
	return acc, nil
}

func DeleteAccount(uid int64) error {
	// set the delete_at field , use db.Unscoped().Delete() delete the row
	if err := dao.Db.Where("uid = ?", uid).Delete(Account{}).Error; err != nil {
		return err
	}
	return nil
}
