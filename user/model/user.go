package model

import (
	"errors"
	"time"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`

	CreatedAt time.Time `json:"created_at"` //these fields would be managed automatically
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	ErrInvalidName = errors.New("user has and invalid name")
	ErrInvalidAge  = errors.New("user has invalid age")
)

func NewUser(first, last string, age int) (*User, error) {
	if err := checkName(first, last); err != nil {
		return nil, err
	}
	if err := checkAge(age); err != nil {
		return nil, err
	}
	return &User{FirstName: first, LastName: last, Age: age}, nil
}

func checkName(first, last string) error {
	if first == "" || last == "" {
		return ErrInvalidName
	}
	return nil
}

func checkAge(age int) error {
	if age < 0 || age > 160 {
		return ErrInvalidAge
	}
	return nil
}

func Validateuser(u *User) error {
	if err := checkName(u.FirstName, u.LastName); err != nil {
		return err
	}
	if err := checkAge(u.Age); err != nil {
		return err
	}
	return nil
}
