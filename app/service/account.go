package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"mall/app/internal/model"
)

func (s *Service) GetAccount(id int) (*model.Account, error) {
	return s.dao.GetAccount(uint(id))
}

func (s *Service) CreateAccount(acc *model.Account) (uint, error) {
	if acc.Password != acc.PasswordReapt {
		return 0, fmt.Errorf("repeat password check fail")
	}
	bcryptByte, err := bcrypt.GenerateFromPassword([]byte(acc.Password), 12)
	if err != nil {
		return 0, fmt.Errorf("generate password failed case: %s", err.Error())
	}
	acc.Password = string(bcryptByte)
	return s.dao.CreateAccount(acc)
}

func (s *Service) DelAccount(id int) error {
	return s.dao.DelAccount(uint(id))
}

func (s *Service) Login(name string, password string) (uid uint, err error) {
	acc, err := s.dao.GetAccountByName(name)
	if err != nil {
		return 0, err
	}
	if bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)) != nil {
		return 0, errors.New("invalid password")
	}

	return acc.ID, nil
}
