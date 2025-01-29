package models

import (
	"errors"
	"strings"

	"github.com/lib/pq"
)

type userSignupFormData interface {
	GetName() string
	GetEmail() string
	GetHashedPassword() (string, error)
}

func (m *UserModel) Insert(f userSignupFormData) error {
	hashedPassword, err := f.GetHashedPassword()
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (name, email, hashed_password, created)
VALUES ($1, $2, $3, NOW())`

	_, err = m.Exec(stmt,
		f.GetName(),
		f.GetEmail(),
		hashedPassword,
	)
	if err != nil {
		var postgresErr *pq.Error

		if errors.As(err, &postgresErr) {
			if postgresErr.Code == "23505" && strings.Contains(postgresErr.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}
