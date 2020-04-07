package account

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mall/internal/pkg/ecode"
	"mall/internal/pkg/log"
	"mall/internal/pkg/uuid"
)

type Service interface {
	CreateAccount(acc *Account) (int64, error)
	ReadAccount(uid int64) (res *Account, err error)
	DeleteAccount(id int64) error
	Login(phone string, password string) (uid uint, err error)
}

type service struct {
}

func (g *service) CreateAccount(acc *Account) (int64, error) {
	if _, err := DReadAccountByPhone(acc.Phone); err == nil {
		return 0, ecode.PhoneHasBeenRegistered
	}
	if acc.Password != acc.PasswordReapt {
		return 0, ecode.PwdRepeatCheckErr
	}
	bcryptByte, err := bcrypt.GenerateFromPassword([]byte(acc.Password), 12)
	if err != nil {
		log.Logger.Errorf("generate password failed case: %s", err.Error())
		return 0, ecode.SavePwdErr
	}
	acc.Password = string(bcryptByte)
	acc.UId = uuid.UUID()
	return DCreateAccount(acc)
}

func (g *service) ReadAccount(uid int64) (res *Account, err error) {
	if res, err = DReadAccount(uid); err != nil {
		return nil, err
	}
	if res == nil {
		return nil, ecode.AccountNotFound
	}
	return res, nil
}

func (g *service) DeleteAccount(id int64) error {
	return DDeleteAccount(id)
}

func (g *service) Login(phone string, password string) (uid uint, err error) {
	acc, err := DReadAccountByPhone(phone)
	if err != nil {
		return 0, err
	}
	if bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)) != nil {
		return 0, errors.New("invalid password")
	}

	return acc.ID, nil
}
