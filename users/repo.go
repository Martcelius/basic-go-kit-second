package users

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Register(ctx context.Context, user users) error
}

type UserRepo struct {
	Db *sql.DB
}

func (r *UserRepo) Register(ctx context.Context, user users) error {
	// hashed password before insert to database

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	_, err = r.Db.Exec("INSERT INTO users (firstname, lastname, email, password) VALUES (?,?,?,?)", user.Firstname, user.Lastname, user.Email, string(hashedPass))

	if err != nil {
		fmt.Println("Error occured inside Register in repo", err)
		return err
	}

	return nil

}
