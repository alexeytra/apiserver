package store

import "http-rest-api/v1/internal/app/model"

type UserRepository struct {
	store *Store
}

func (u *UserRepository) Create(m *model.User) (*model.User, error) {
	return nil, nil
}

func (u *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}