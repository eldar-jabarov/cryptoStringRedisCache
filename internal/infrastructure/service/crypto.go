package service

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

type Crypter interface {
	MD5(string) (string, error)
	SHA256(string) (string, error)
}

type Crypto struct{}

func NewCrypter() Crypter {
	return &Crypto{}
}

func (c Crypto) MD5(s string) (string, error) {
	return fmt.Sprintf("%x", md5.Sum([]byte(s))), nil
}

func (c Crypto) SHA256(s string) (string, error) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s))), nil
}
