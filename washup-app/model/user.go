package model

import (

	"gorm.io/gorm"
)

type UserClient struct {
    DB *gorm.DB
}

type User struct {
    gorm.Model
    Name  string `json:"name"`
    Email string `json:"email" gorm:"uniqueIndex"`
    Password string 
    Reservations []Reservation `gorm:"foreignKey:UserID"`
}

type UserRepository interface {
    SaveUser(user *User) error
    UserFirst(query string, args ...interface{}) (*User, error)
    DeleteUser(user *User) error
}

func (u UserClient) SaveUser(user *User) error {
    return u.DB.Save(user).Error
}

func (u UserClient) UserFirst(query string, args ...interface{}) (*User, error) {
    var user User
    if err := u.DB.Where(query, args...).First(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil
}

func (u *UserClient) DeleteUser(user *User) error {
    return u.DB.Delete(user).Error
}