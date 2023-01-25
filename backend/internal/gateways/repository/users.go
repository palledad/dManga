package repository

import "gorm.io/gorm"

type User struct {
	WalletAddress string `gorm:"primarykey"`
	Name          string `gorm:"unique"`
	IconUrl       string
	Biography     string
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
