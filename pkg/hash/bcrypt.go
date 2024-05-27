package hash

import (
	"assesment/pkg/configs"

	tools "golang.org/x/crypto/bcrypt"
)

func NewBcrypt(config configs.Bcrypt) *Bcrypt {
	return &Bcrypt{
		config: config,
	}
}

type Bcrypt struct {
	config configs.Bcrypt
}

func (h *Bcrypt) Encrypt(password string) string {
	hashedPassword, _ := tools.GenerateFromPassword([]byte(password), h.config.Salt)
	return string(hashedPassword)
}

func (h *Bcrypt) Compare(password string, hashedPassword string) bool {
	err := tools.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
