package security

import "golang.org/x/crypto/bcrypt"

func GenerateHashedPassword(textPass string) (string, error) {
	pass := []byte(textPass)

	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hash string,
	password string,
) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
