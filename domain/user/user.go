package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmailInvalid     = errors.New("Invalid email")
	ErrFirstNameInvalid = errors.New("Invalid first name")
	ErrLastNameInvalid  = errors.New("Invalid last name")
)

type User struct {
	user *Profile
}

func (u User) GetID() uuid.UUID {
	return u.user.id
}
func (u User) GetFirstName() string {
	return u.user.firstName
}
func (u User) GetLastName() string {
	return u.user.lastName
}
func (u User) GetEmail() string {
	return u.user.email
}

func NewUser(email, firstName, lastName string) (User, error) {
	return NewUserWithId(uuid.New(), email, firstName, lastName)
}

func NewUserWithId(id uuid.UUID, email, firstName, lastName string) (User, error) {
	if email == "" {
		return User{}, ErrEmailInvalid
	}
	if firstName == "" {
		return User{}, ErrFirstNameInvalid
	}
	if lastName == "" {
		return User{}, ErrLastNameInvalid
	}

	return User{
		user: &Profile{
			id:        id,
			email:     email,
			firstName: firstName,
			lastName:  lastName,
		},
	}, nil
}
