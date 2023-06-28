package types

import (
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const ENCRYPTPASSCOST = 12

type UserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}
type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Email           string             `json:"email,omitempty"`
	EncodedPassword string             `json:"-" bson:"encoded_password"`
}

func (p *UserParams) Validate() []string {
	valid := []string{}
	if len(p.FirstName) < 2 {
		valid = append(valid, "FirstName length must be at least 2 characters")
	}
	if len(p.LastName) < 2 {
		valid = append(valid, "LastName length must be at least 2 characters")
	}
	if len(p.Password) < 7 {
		valid = append(valid, "Password length must be at least 7 characters")
	}
	if !isEmail(p.Email) {
		valid = append(valid, "invalid email address")
	}
	return valid
}

func isEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func NewUserFromParams(params UserParams) (*User, error) {
	encPW, err := bcrypt.GenerateFromPassword([]byte(params.Password), ENCRYPTPASSCOST)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:       params.FirstName,
		LastName:        params.LastName,
		Email:           params.Email,
		EncodedPassword: string(encPW),
	}, nil
}
