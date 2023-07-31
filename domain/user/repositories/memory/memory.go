package memory

import (
	"stagingmanager/domain/user"

	"github.com/google/uuid"
)

type memoryRepository struct {
	users []user.User
}

// CreateUser implements user.Repository.
func (r *memoryRepository) CreateUser(user2 user.User) (uuid.UUID, error) {
	r.users = append(r.users, user2)
	return user2.GetID(), nil
}

// GetUsers implements user.Repository.
func (r *memoryRepository) GetUsers() ([]user.User, error) {
	return r.users, nil
}

func NewMemoryRepository() user.Repository {
	return &memoryRepository{
		users: make([]user.User, 0),
	}
}
