package user

import (
	"github.com/google/uuid"
)

type Repository interface {
	CreateUser(user User) (uuid.UUID, error)
	GetUsers() ([]User, error)
}
