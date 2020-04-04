package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mall/app/internal/model"
	"mall/library/ecode"
	"mall/library/log"
	"mall/library/uuid"
)

func (s *Service) CreateAccount(acc *model.Account) (int64, error) {
	if _, err := s.dao.ReadAccountByPhone(acc.Phone); err == nil {
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
	return s.dao.CreateAccount(acc)
}

func (s *Service) ReadAccount(uid int64) (res *model.Account, err error) {
	if res, err = s.dao.ReadAccount(uid); err != nil {
		return nil, err
	}
	if res == nil {
		return nil, ecode.AccountNotFound
	}
	return res, nil
}

func (s *Service) DeleteAccount(id int64) error {
	return s.dao.DeleteAccount(id)
}

func (s *Service) Login(phone string, password string) (uid uint, err error) {
	acc, err := s.dao.ReadAccountByPhone(phone)
	if err != nil {
		return 0, err
	}
	if bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)) != nil {
		return 0, errors.New("invalid password")
	}

	return acc.ID, nil
}
