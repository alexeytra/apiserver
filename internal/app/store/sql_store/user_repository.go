package sql_store

import (
	"http-rest-api/v1/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (u *UserRepository) Create(m *model.User) error {
	if err := m.Validate(); err != nil {
		return err
	}

	if err := m.BeforeCreate(); err != nil {
		return err
	}
	return u.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		m.Email,
		m.EncryptedPassword,
	).Scan(&m.ID)
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := ur.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1", email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}
