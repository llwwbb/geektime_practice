package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// db should not be here
// just for example, not for production
var db *sql.DB

type User struct {
	Id   int64
	Name string
	Age  int
}

func FindUserById(id int64) (*User, error) {
	row := db.QueryRow("select `name`, `age` from `user` where id = ?", id)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	var (
		name string
		age  int
	)
	err = row.Scan(&name, &age)
	// we can return nil directly when meet sql.ErrNoRows
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("FindUserById id = %v error", id))
	}
	return &User{
		Id:   id,
		Name: name,
		Age:  age,
	}, nil
}

func FindUserById2(id int64) (*User, error) {
	row := db.QueryRow("select `name`, `age` from `user` where id = ?", id)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	var (
		name string
		age  int
	)
	err = row.Scan(&name, &age)

	// We can also replace it with a specific error if we need to do some specific operations,
	// such as returning a 404 error code in the http service
	// you don't want to import sql in controller layer, right?
	if err == sql.ErrNoRows {
		// usually call stack is not needed here
		return nil, fmt.Errorf("FindUserById id = %v, %w", id, ErrNotFound)
		// but if you want
		// return nil, errors.Wrapf(ErrNotFound, "FindUserById id = %v", id)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "FindUserById id = %v", id)
	}
	return &User{
		Id:   id,
		Name: name,
		Age:  age,
	}, nil
}
