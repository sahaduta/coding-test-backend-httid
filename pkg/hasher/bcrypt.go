package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	HashPassword(pwd []byte) (string, error)
	ComparePasswords(hashedPwd []byte, plainPwd []byte) error
}

type BcryptHasher struct{}

func NewBcryptHasher() Hasher {
	return &BcryptHasher{}
}

func (b *BcryptHasher) HashPassword(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func (b *BcryptHasher) ComparePasswords(hashedPwd []byte, plainPwd []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
}
