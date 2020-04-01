package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"mall/app/internal/model"
	"mall/library/uuid"
)

func (s *Service) CreateAccount(acc *model.Account) (int64, error) {
	if _, err := s.dao.ReadAccountByPhone(acc.Phone); err == nil {
		return 0, fmt.Errorf("phone has been registered")
	}
	if acc.Password != acc.PasswordReapt {
		return 0, fmt.Errorf("repeat password check fail")
	}
	bcryptByte, err := bcrypt.GenerateFromPassword([]byte(acc.Password), 12)
	if err != nil {
		return 0, fmt.Errorf("generate password failed case: %s", err.Error())
	}
	acc.Password = string(bcryptByte)
	if uid, err := uuid.UUID(); err != nil {
		return 0, fmt.Errorf("uid generated failed %s", err.Error())
	} else {
		acc.UId = uid
	}
	return s.dao.CreateAccount(acc)
}

func (s *Service) ReadAccount(uid int64) (*model.Account, error) {
	return s.dao.ReadAccount(uid)
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
