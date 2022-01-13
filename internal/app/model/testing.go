package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email: "user1@example.com",
		Password: "password",
	}
}