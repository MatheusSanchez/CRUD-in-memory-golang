package repositories

import (
	"github.com/google/uuid"
)

type UserInMemoryRepository struct {
	db map[uuid.UUID]User
}

func NewUserInMemoryRepository(db map[uuid.UUID]User) UserInMemoryRepository {
	return UserInMemoryRepository{
		db: db,
	}
}

func (repo UserInMemoryRepository) Insert(firstName string, lastName string, bio string) User {

	newId := uuid.New()

	repo.db[newId] = User{
		Id:        newId,
		FirstName: firstName,
		LastName:  lastName,
		Biography: bio,
	}

	return repo.db[newId]
}

func (repo UserInMemoryRepository) FindAll() []User {

	users := []User{}

	for _, user := range repo.db {
		users = append(users, user)
	}

	return users
}

func (repo UserInMemoryRepository) FindById(id uuid.UUID) User {

	user := repo.db[id]

	return user

}

func (repo UserInMemoryRepository) Delete(id uuid.UUID) User {

	user := repo.db[id]
	delete(repo.db, id)

	return user
}

func (repo UserInMemoryRepository) Edit(id uuid.UUID, firstName string, lastName string, bio string) User {

	user, ok := repo.db[id]

	if !ok {
		return User{}
	}

	user.FirstName = firstName
	user.LastName = lastName
	user.Biography = bio

	repo.db[id] = user

	return user
}
