package user

import (
	"context"
	"database/sql"

	"github.com/hooneun/if-commerce/db"
)

type UserInterface interface {
	GetUserByID(int64, *sql.DB)
}

func GetUserByID(id int64, conn sql.DB) (db.User, error) {
	queries := db.New(&conn)

	user, err := queries.GetUserByID(context.Background(), id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByEmail(email string, conn sql.DB) (db.User, error) {
	queries := db.New(&conn)

	user, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return user, err
	}

	return user, nil

}

func CreateUserByRequest(createUser db.CreateUserParams, conn sql.DB) error {
	queries := db.New(&conn)

	_, err := queries.CreateUser(context.Background(), createUser)

	if err != nil {
		return err
	}

	return nil
}
