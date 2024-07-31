package repository

import (
	"go-rest-api-udemy/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	// UserRepositoryで持たせるメソッドを設定
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	// UserRepositoryで持たせるプロパティ(属性)を設定
	db *gorm.DB
}

// UserRepositoryのコンストラクタ（インスタンス生成時に渡す引数設定&実行）
// gorm.DBを引数で注入する形にしている（DIパターン）
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// GetUserByEmail関数（UserRepositoryのインスタンスメソッドの役割）
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser関数（UserRepositoryのインスタンスメソッドの役割）
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
