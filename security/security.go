package security

import "golang.org/x/crypto/bcrypt"

var _pwdHash []byte

// Init inits the security
func Init(passwordHash string) {
	_pwdHash = []byte(passwordHash)
}

// HashPassword is a wrapper over bcrypt.GenerateFromPassword
func HashPassword(password []byte) ([]byte, error) {
	saltedPassword := append(password, _pwdHash...)
	return bcrypt.GenerateFromPassword(saltedPassword, 10)
}

// ComparePassword is a wrapper over bcrypt.CompareHashAndPassword
func ComparePassword(hashedPassword, password []byte) error {
	saltedPassword := append(password, _pwdHash...)
	return bcrypt.CompareHashAndPassword(hashedPassword, saltedPassword)
}

// IsPasswordHashed checks if password is hashed
func IsPasswordHashed(hashedPassword []byte) bool {
	_, err := bcrypt.Cost(hashedPassword)
	return err == nil
}
