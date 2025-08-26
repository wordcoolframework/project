package Repository

import (
	"gorm.io/gorm"
	"projectUserManagement/app/Domain/Contracts"
	"projectUserManagement/app/Domain/Entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Contracts.IUserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *Entity.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetByID(id uint) (*Entity.User, error) {
	var user Entity.User

	err := ur.db.First(&user, id).Error

	return &user, err
}

func (ur *UserRepository) GetByPhone(phone string) (*Entity.User, error) {
	var user Entity.User
	err := ur.db.Where("phone=?", phone).First(&user).Error
	return &user, err
}

func (ur *UserRepository) GetAll(search string, page, limit int) ([]Entity.User, int64, error) {
	var users []Entity.User
	var count int64

	query := ur.db.Model(&Entity.User{})
	if search != "" {
		query = query.Where("phone LIKE ?", "%"+search+"%")
	}

	query.Count(&count)

	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, count, err
}
