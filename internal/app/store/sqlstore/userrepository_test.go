package sqlstore_test

import (
	"github.com/dmaok/web-1.1/internal/app/model"
	"github.com/dmaok/web-1.1/internal/app/store"
	"github.com/dmaok/web-1.1/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)
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
