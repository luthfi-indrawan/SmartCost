package helper

import "golang.org/x/crypto/bcrypt"

func (h *Helper) CompareStringWithHash(plain, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}

func (h *Helper) HashingPassword(plain string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plain), 12)
}