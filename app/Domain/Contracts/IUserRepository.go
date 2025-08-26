package Contracts

import "projectUserManagement/app/Domain/Entity"

type IUserRepository interface {
	Create(user *Entity.User) error

	GetByPhone(phone string) (*Entity.User, error)

	GetByID(id uint) (*Entity.User, error)

	GetAll(search string, page, limit int) ([]Entity.User, int64, error)
}
