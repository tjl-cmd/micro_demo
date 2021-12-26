package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/domain/model"
	"user/domain/repository"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(*model.User, bool) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

func NewUserDataService(UserRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: UserRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

// GeneratePassword 加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword 验证用户密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码对比错误")
	}
	return true, nil
}

// AddUser 插入用户
func (u UserDataService) AddUser(user *model.User) (int64, error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

// DeleteUser 删除用户
func (u UserDataService) DeleteUser(i int64) error {
	return u.UserRepository.DeleteUserByID(i)
}

// UpdateUser 更新用户
func (u UserDataService) UpdateUser(user *model.User, b bool) (err error) {
	if b  {
		pwdByte,err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

// FindUserByName 查找用户
func (u UserDataService) FindUserByName(s string) (*model.User, error) {
	return u.UserRepository.FindUserByName(s)
}

// CheckPwd 验证用户密码
func (u UserDataService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user,err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false,err
	}
	return ValidatePassword(pwd,user.HashPassword)
}
