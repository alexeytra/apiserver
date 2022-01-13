package sql_store_test

import (
	"http-rest-api/v1/internal/app/model"
	"http-rest-api/v1/internal/app/store/sql_store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sql_store.TestDB(t, databaseURL)
	defer teardown("users")
	s := sql_store.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sql_store.TestDB(t, databaseURL)
	defer teardown("users")
	s := sql_store.New(db)
	email := "user1@mail.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
