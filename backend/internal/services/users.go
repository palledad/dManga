package services

import (
	"github.com/palledad/dManga/backend/internal/models"
	"gorm.io/gorm"
)

type UsersService struct {
	userModel *models.UserModel
	db        *gorm.DB
}

func NewUsersService(db *gorm.DB, userModel *models.UserModel) *UsersService {
	return &UsersService{
		userModel: userModel,
		db:        db,
	}
}

func (s UsersService) CreateUser(user *models.User) (*models.User, error) {
	if err := s.userModel.CreateUser(s.db, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s UsersService) ReadUser(walletAddress string) (*models.User, error) {
	user, err := s.userModel.ReadUser(s.db, walletAddress)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UsersService) DeleteUser(walletAddress string) (*models.User, error) {
	user, err := s.userModel.DeleteUser(s.db, walletAddress)
	if err != nil {
		return nil, err
	}
	return user, nil
}
