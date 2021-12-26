package repository

import (
	"github.com/jinzhu/gorm"
	"user/domain/model"
)

type IUserRepository interface {
	// InitTable 初始化数据表
	InitTable() error
	// FindUserByName 根据用户名称查询用户信息
	FindUserByName(string) (*model.User, error)
	// FindUserByID 根据用户ID查找用户信息
	FindUserByID(int64) (*model.User, error)
	// CreateUser 创建用户
	CreateUser(*model.User) (int64, error)
	// DeleteUserByID 根据用户ID删除用户
	DeleteUserByID(int64) error
	// UpdateUser 更新用户信息
	UpdateUser(*model.User) error
	// FindAll 查找所有用户
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

// InitTable 初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.User{}).Error
}

// FindUserByName 根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.Where("user_name = ?", name).Find(user).Error
}

// FindUserByID 根据用户ID查询用户信息
func (u *UserRepository) FindUserByID(userId int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.First(user, userId).Error
}

// CreateUser 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userId int64, err error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

// DeleteUserByID 根据用户ID删除用户
func (u *UserRepository) DeleteUserByID(userId int64) error {
	return u.mysqlDB.Where("id = ?", userId).Delete(&model.User{ID: userId}).Error
}
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDB.Model(user).Update(&user).Error
}
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDB.Find(&userAll).Error
}
