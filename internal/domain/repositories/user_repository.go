package repositories

import (
	"github.com/faruqii/goproto/internal/domain/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entities.User) (err error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(user *entities.User) (err error) {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
