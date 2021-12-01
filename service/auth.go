package service

import (
	"errors"
	"fmt"
	"popaket/entity"
	"popaket/formatter"
	"popaket/storage"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(userUser entity.AuthInputs) (formatter.UserFormat, error)
	LoginUser(input entity.LoginAuthInputs) (formatter.UserFormat, error)
}

type userservice struct {
	dao storage.UserDao
}

func NewUserService(dao storage.UserDao) *userservice {
	return &userservice{dao}
}

func (s *userservice) RegisterUser(userUser entity.AuthInputs) (formatter.UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(userUser.Password), bcrypt.MinCost)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if err != nil {
		return formatter.UserFormat{}, err
	}

	useruuid, err := uuid.NewV4()
	if err != nil {
		return formatter.UserFormat{}, err
	}

	msisdn := "62" + userUser.Msisdn

	var newUser = entity.Auths{
		ID:        useruuid.String(),
		Username:  userUser.Username,
		Password:  string(genPassword),
		Name:      userUser.Name,
		Msisdn:    msisdn,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.dao.RegisterUser(newUser)
	formatUser := formatter.FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *userservice) LoginUser(input entity.LoginAuthInputs) (formatter.UserFormat, error) {
	userUser, err := s.dao.FindUserByUsername(input.Username)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if len(userUser.ID) != 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(userUser.Password), []byte(input.Password)); err != nil {
			return formatter.UserFormat{}, errors.New("password invalid")
		}

		formatter := formatter.FormatUser(userUser)

		return formatter, nil
	}

	newError := fmt.Sprintf("user id %v not found", userUser.ID)
	return formatter.UserFormat{}, errors.New(newError)
}
