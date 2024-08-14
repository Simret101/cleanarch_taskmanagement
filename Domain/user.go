package Domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type Credentials struct {
	Username string
	Password string
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
type UserRepository interface {
	GetUserByUsername(username string) (*User, error)

	CreateUser(user *User) error
}
