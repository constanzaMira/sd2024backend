package service

import (
	"hello-run/model"

	"errors"
)

type UserParams struct {
	Username     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(ur model.UserClient, userParams UserParams) (*model.User, error) {
	collisionUser, err := ur.UserFirst("email = ?", userParams.Email)
	if collisionUser != nil {
		return nil, errors.New("user already exists")
	}

	user := &model.User{
		Name:     userParams.Username,
		Email:    userParams.Email,
		Password: hashPassword(userParams.Password),
	}

	err = ur.SaveUser(user)
	if err != nil {
		return nil, errors.New("error trying to save user")
	}

	return user, nil
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
	FullName string `json:"fullname"`
}

func LoginUser(ur model.UserClient, loginParams LoginParams) (*LoginResponse, error) {
	user, err := ur.UserFirst("email = ?", loginParams.Email)
	if err != nil {
		return nil, errors.New("error trying to find user")
	}

	err = ValidatePassword(user.Password, loginParams.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := GenerateToken(user.Email)
	if err != nil {
		return nil, errors.New("error trying to generate token")
	}
	
	return &LoginResponse{
		Email:    user.Email,
		Token: token,
		FullName: user.Name,
	}, nil
}

type GoogleLoginParams struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Username string `json:"username"`
}

func GoogleLoginUser(ur model.UserClient, googleLoginParams GoogleLoginParams) (*LoginResponse, error) {
	user, err := ur.UserFirst("email = ?", googleLoginParams.Email)
	if err != nil {
		user = &model.User{
			Name:     googleLoginParams.Username,
			Email:    googleLoginParams.Email,
		}
	
		err = ur.SaveUser(user)
		if err != nil {
			return nil, errors.New("error trying to save user")
		}
	}

	token, err := GenerateToken(user.Email)
	if err != nil {
		return nil, errors.New("error trying to generate token")
	}
	
	return &LoginResponse{
		Email:    user.Email,
		Token: token,
	}, nil
}