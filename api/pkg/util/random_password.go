package util

import (
	"crypto/rand"
	"fmt"
)

func MakeRandomPassword() (string, error) {

	bytes := make([]byte, 8)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", bytes), nil

}
