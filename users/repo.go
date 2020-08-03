package users

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Register(ctx context.Context, user users) error
	Login(ctx context.Context, user interface{}) (interface{}, error)
}

type UserRepo struct {
	Db *sql.DB
}

func (r *UserRepo) Register(ctx context.Context, user users) error {
	// hashed password before insert to database

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	//Generate jwt
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = user.Email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("token-secret"))

	if err != nil {
		fmt.Println("Error occured inside Register in repo", err)
		return err
	}
	_, err = r.Db.Exec("INSERT INTO users (firstname, lastname, email, password, token) VALUES (?,?,?,?,?)", user.Firstname, user.Lastname, user.Email, string(hashedPass), string(token))

	if err != nil {
		fmt.Println("Error occured inside Register in repo", err)
		return err
	}

	return nil

}

func (r *UserRepo) Login(ctx context.Context, user users) (users, error) {
	userLogin := users{}
	row := r.Db.QueryRow("SELECT email, password FROM users WHERE email=?", user.Email)

	err := row.Scan(&userLogin.Email, &userLogin.Password)
	if err != nil {
		fmt.Println("Error occured inside Login in repo", err)
		return user, err
	}
	fmt.Println("data rep", userLogin)
	err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(user.Password))

	if err != nil {
		fmt.Println("Error occured inside Login in repo", err)
		return user, err
	}

	return userLogin, nil

}
