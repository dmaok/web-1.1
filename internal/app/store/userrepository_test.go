package store_test

import (
	"github.com/dmaok/web-1.1/internal/app/model"
	"github.com/dmaok/web-1.1/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "test@mail.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "test@mail.com"
	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	if _, err := s.User().Create(&model.User{
		Email: email,
	}); err != nil {
		t.Fatal(err)
	}

	u, err := s.User().FindByEmail("test@mail.com")

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
