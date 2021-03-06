package stub

import (
	"errors"

	"github.com/mzki/go-chat/entity"
)

type UserRepository struct{}

var ErrNotFound = errors.New("not found")

func (repo UserRepository) Get(name string, password string) (entity.User, error) {
	if repo.Exist(name, password) {
		return entity.User{ID: 0, Name: name, Password: password}, nil
	}
	return entity.User{}, ErrNotFound
}

func (repo UserRepository) Set(name string, password string) (uint64, error) {
	return 0, nil
}

func (repo UserRepository) Exist(name string, password string) bool {
	return name == "user" && password == "password"
}

func (repo UserRepository) Find(id uint64) (entity.User, error) {
	return entity.User{}, ErrNotFound
}
