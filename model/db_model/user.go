package db_model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"my-backend/global"
)

type User struct {
	global.BaseModel
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	Password string `json:"-"`
	Status   string `json:"status" gorm:"default:inactive"`
	RoleName string `json:"roleName"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 10
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"

	// DefaultPassword 管理员添加用户时的默认密码
	DefaultPassword = "123456"
)

// GetAll 查全部用户
func (u *User) GetAll() (objs []User, err error) {
	db := global.DB
	err = db.Find(&objs).Error
	return
}

// Get 获取用户
func (u *User) Get() error {
	result := global.DB.First(&u)
	return result.Error
}

// Insert 创建用户
func (u *User) Insert() error {
	err := u.SetPassword(u.Password)
	if err != nil {
		return err
	}
	return global.DB.Create(&u).Error
}

// UpdateWithoutZero 更改用户，不会更新零值
func (u *User) UpdateWithoutZero() error {
	return global.DB.Updates(&u).Error
}

// SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword() (bool, error) {
	var existedUser User
	err := global.DB.First(&existedUser, "phone = ?", u.Phone).Error
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(u.Password))
	return err == nil, err
}

// CheckUser 校验用户
func (u *User) CheckUser() (bool, error) {
	// 从数据库中查找用户
	var user User
	err := global.DB.First(&user, "phone = ?", u.Phone).Error
	if err != nil {
		return false, err
	}
	fmt.Printf("======user: %+v\n", user)
	// 校验密码，第一个参数数据库中的加密后的密码、第二个参数为字节切片转换过的用户输入密码值
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	ok := err == nil //&& user.Status == Active
	if ok {
		err := global.DB.First(&u, "phone = ?", u.Phone).Error
		if err != nil {
			return false, err
		}
		return true, err
	} else {
		return false, err
	}
}
