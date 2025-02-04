package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"regexp"

	"github.com/google/uuid"
)

var (
	ErrInvalidEmail      = errors.New("invalid email address")
	ErrUserShortPassword = errors.New("password must be at least 8 characters long")
)

type (
	UserID   uint
	UserUUID = uuid.UUID
	Email    string
)

var NilUUID = UserUUID{}

func (e *Email) IsValid() bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	r := regexp.MustCompile(emailRegex)
	return r.Match([]byte(*e))
}

func (e *Email) String() string {
	return string(*e)
}

func NewPassword(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

type User struct {
	ID       UserID
	UUID     UserUUID
	Email    Email
	Password string
}

func (u *User) PasswordIsCorrect(pass string) bool {
	return u.Password == NewPassword(pass)
}

func (u *User) Validate() error {
	if !u.Email.IsValid() {
		return ErrInvalidEmail
	}
	if len(u.Password) < 8 {
		return ErrUserShortPassword
	}
	return nil
}

type UserFilter struct {
	ID    UserID
	UUID  UserUUID
	Email Email
}

func (f *UserFilter) IsValid() bool {
	return f.ID > 0 || f.UUID != UserUUID{} || f.Email.IsValid()
}
