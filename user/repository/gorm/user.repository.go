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
	return &gormUserRepository{
		Conn: conn,
	}
}

func (u gormUserRepository) GetByID(id int64) (domain.User, error) {
	var user domain.User // su dung var de tao variable co type san
	if err := u.Conn.Where("id = ?", id).Find(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u gormUserRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	if err := u.Conn.Where("username = ?", username).Find(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u gormUserRepository) Store(user *domain.User) error {
	if err := u.Conn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u gormUserRepository) Update(id int64, user *domain.User) (updatedRow int64, err error) {
	result := u.Conn.Where("id = ?", id).Updates(&user)
	if err := result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}
