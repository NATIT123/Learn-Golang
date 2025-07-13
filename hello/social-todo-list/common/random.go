package common

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type rune = int32

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string {
	b := make([]rune, n)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range b {
		b[i] = letters[r1.Intn(999999)%len(letters)]
	}
	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return randSequence(length)
}

type bcryptHash struct{}

// NewBcryptHash tạo đối tượng bcryptHash mới
func NewBcryptHash() *bcryptHash {
	return &bcryptHash{}
}

// Hash sử dụng bcrypt để hash mật khẩu
func (h *bcryptHash) Hash(data string) (string, error) {
	// bcrypt tự động sinh salt và mã hóa mật khẩu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *bcryptHash) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
