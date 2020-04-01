package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"mall/app/internal/model"
)

func (s *Service) CreateAccount(acc *model.Account) (uint, error) {
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
	return s.dao.CreateAccount(acc)
}

func (s *Service) ReadAccount(id int) (*model.Account, error) {
	return s.dao.ReadAccount(uint(id))
}

func (s *Service) DeleteAccount(id int) error {
	return s.dao.DeleteAccount(uint(id))
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
