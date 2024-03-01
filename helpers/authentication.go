package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// GenerateToken menghasilkan token JWT untuk pengguna yang diberikan
func GenerateToken(userID uint) (string, error) {
	// Buat token dengan klaim subjek dan kedaluwarsa
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Tanda tangan token dan dapatkan string token yang lengkap
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

// ComparePasswords membandingkan kata sandi yang diberikan dengan hash kata sandi yang tersimpan
func ComparePasswords(hashedPassword []byte, plainPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword))
}

// GeneratePasswordHash menghasilkan hash dari sebuah kata sandi menggunakan bcrypt
func GeneratePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
