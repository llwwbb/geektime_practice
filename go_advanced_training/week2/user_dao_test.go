package dao

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:password@/test")
	if err != nil {
		panic(err)
	}
}

func TestFindUserByIdSuccess(t *testing.T) {
	user, err := FindUserById(1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestFindUserByIdNotFound(t *testing.T) {
	user, err := FindUserById(999)
	assert.NoError(t, err)
	assert.Nil(t, user)
}

func TestFindUserById2NotFound(t *testing.T) {
	user, err := FindUserById2(999)
	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, user)
}
