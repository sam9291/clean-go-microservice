package user

import "github.com/google/uuid"

type Profile struct {
	id        uuid.UUID
	email     string
	firstName string
	lastName  string
}
