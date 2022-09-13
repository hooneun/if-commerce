package auth

import (
	"database/sql"

	"github.com/hooneun/if-commerce/db"
	"github.com/hooneun/if-commerce/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthInterface interface {
	Signup(db.CreateUserParams, *sql.DB)
	Login(LoginRequest, *sql.DB)
}

type LoginRequest struct {
	Email    string
	Password string
}

func Signup(createUser db.CreateUserParams, conn sql.DB) error {
	password, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	createUser.Password = string(password)

	err = user.CreateUserByRequest(createUser, conn)
	if err != nil {
		return err
	}

	return nil
}

func Login(login LoginRequest, conn sql.DB) (db.User, error) {
	user, err := user.GetUserByEmail(login.Email, conn)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}
