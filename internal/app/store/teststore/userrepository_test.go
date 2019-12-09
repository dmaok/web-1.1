package teststore_test

import (
	"github.com/dmaok/web-1.1/internal/app/model"
	"github.com/dmaok/web-1.1/internal/app/store"
	"github.com/dmaok/web-1.1/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "test@mail.com"
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(t)
	testUser.Email = email

	if err := s.User().Create(testUser); err != nil {
		t.Fatal(err)
	}

	u, err := s.User().FindByEmail("test@mail.com")

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	id := 300
	_, err := s.User().Find(id)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(t)

	if err := s.User().Create(testUser); err != nil {
		t.Fatal(err)
	}

	u, err := s.User().Find(testUser.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
