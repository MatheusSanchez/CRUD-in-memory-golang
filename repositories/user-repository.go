package repositories

import "github.com/google/uuid"

type UserRepository interface {
	Insert(firstName string, lastName string, bio string) User
	Edit(id uuid.UUID, firstName string, lastName string, bio string) User
	FindAll() []User
	FindById(id uuid.UUID) User
	Delete(id uuid.UUID) User
}

type User struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Biography string    `json:"bio"`
}
