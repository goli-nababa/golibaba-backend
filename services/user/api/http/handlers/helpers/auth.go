package helpers

import "crypto/rand"

func GenerateOTP() (string, error) {
	n := 6
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < n; i++ {
		bytes[i] = (bytes[i] % 10) + '0'
	}

	return string(bytes), nil
}
