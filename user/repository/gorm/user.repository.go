package gorm

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mercee-be/domain"
)

// https://tour.golang.org/methods/1
type gormUserRepository struct {
	Conn *gorm.DB
}

func NewGormUserRepository(conn *gorm.DB) domain.UserRepository {
	return gormUserRepository{
		Conn: conn,
	}
}

func (u gormUserRepository) GetByID(id int64) (domain.User, error) {
	var user domain.User // su dung var de tao variable co type san
	if err := u.Conn.Where("id = ?", id).Find(&user).Error; err != nil {
		return domain.User{}, nil
	}
	return user, nil
}
