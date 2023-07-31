package account

import (
	"errors"
	"stagingmanager/domain/user"
	"stagingmanager/domain/user/repositories/memory"
	"stagingmanager/domain/user/repositories/sqlserver"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	ErrInvalidServiceConfiguration = errors.New("Invalid service configuration")
)

type Service interface {
	CreateUsers(email string, firstName string, lastName string) error
	GetUsers() ([]user.User, error)
}

type AccountConfiguration func(s *AccountService) error

type AccountService struct {
	repository user.Repository
}

func NewService(configurations ...AccountConfiguration) (Service, error) {
	s := &AccountService{
		repository: nil,
	}

	for _, config := range configurations {
		err := config(s)
		if err != nil {
			return nil, err
		}
	}

	if s.repository == nil {
		return nil, ErrInvalidServiceConfiguration
	}

	return s, nil
}

func WithMemoryRepository() AccountConfiguration {
	return func(s *AccountService) error {
		s.repository = memory.NewMemoryRepository()
		return nil
	}
}

func WithSqlServerRepository(connectionString string) AccountConfiguration {
	return func(s *AccountService) error {
		s.repository = sqlserver.NewSqlServerRepository(connectionString)
		return nil
	}
}

// CreateUsers implements Service.
func (s AccountService) CreateUsers(email string, firstName string, lastName string) error {
	user, err := user.NewUser(email, firstName, lastName)

	if err != nil {
		return err
	}

	_, err = s.repository.CreateUser(user)

	return err
}

// GetUsers implements Service.
func (s AccountService) GetUsers() ([]user.User, error) {
	return s.repository.GetUsers()
}
