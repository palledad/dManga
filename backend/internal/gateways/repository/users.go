package repository

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

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r Repository) ReadUser(walletAddress string) *User {
	target := &User{WalletAddress: walletAddress}
	r.DB.First(&target)
	return target
}

func (r Repository) CreateUser(user *User) *User {
	r.DB.Create(user)
	return user
}
