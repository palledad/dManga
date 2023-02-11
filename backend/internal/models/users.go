package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	WalletAddress string         `gorm:"primarykey"`
	Name          sql.NullString `gorm:"unique"`
	IconUrl       sql.NullString
	Biography     sql.NullString
}

type UserModel struct{}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (m UserModel) ReadUser(db *gorm.DB, walletAddress string) (*User, error) {
	target := &User{WalletAddress: walletAddress}
	if result := db.First(&target); result.Error != nil {
		return nil, result.Error
	}
	return target, nil
}

func (m UserModel) CreateUser(db *gorm.DB, user *User) error {
	if result := db.Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (m UserModel) DeleteUser(db *gorm.DB, walletAddress string) (*User, error) {
	target := &User{WalletAddress: walletAddress}
	if result := db.First(&target); result.Error != nil {
		return nil, result.Error
	}

	if result := db.Delete(&target); result.Error != nil {
		return target, result.Error
	}

	return target, nil
}
