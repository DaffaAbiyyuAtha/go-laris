package lib

import (
	"fmt"

	"github.com/matthewhartstonge/argon2"
)

var Argon2 = argon2.DefaultConfig()

func Encrypt(password string) (string, error) {
	encoded, err := Argon2.HashEncoded([]byte(password))
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}

	return string(encoded), nil
}

func Verify(plainPassword string, encryptedPassword string) bool {
	ok, err := argon2.VerifyEncoded([]byte(plainPassword), []byte(encryptedPassword))
	return err == nil && ok
}
