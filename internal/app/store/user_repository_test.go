package store_test

import (
	"http-rest-api/v1/internal/app/model"
	"http-rest-api/v1/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "test@mail.com",
		EncryptedPassword: "52525",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}