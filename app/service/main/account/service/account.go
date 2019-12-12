package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"mall/app/service/main/account/internal/model"
)

func (s *service) GetAccount(id int) (*model.Account, error) {
	return s.dao.GetAccount(uint(id))
}

func (s *service) CreateAccount(acc *model.Account) (uint, error) {
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

func (s *service) DelAccount(id int) error {
	return s.dao.DelAccount(uint(id))
}