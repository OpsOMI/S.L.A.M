package hasherpkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/matthewhartstonge/argon2"
)

// IHasher defines the interface for hashing and encryption operations
type IHasher interface {
	// For Auth
	HashArgon2(password string) (string, error)
	CompareArgon2(hashedPassword, password string) (bool, error)

	EncryptMessage(message string, secretKey []byte) (string, error)
	DecryptMessage(encrypted string, secretKey []byte) (string, error)

	// Generate 2FA
	Generate6DigitCode() (string, error)
}

// Hasher struct implements the IHasher interface
type Hasher struct{}

// New creates a new instance of Hasher
func New() IHasher {
	return &Hasher{}
}

// HashArgon2 hashes a password using Argon2id algorithm
func (h *Hasher) HashArgon2(password string) (string, error) {
	argon := argon2.DefaultConfig()
	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompareArgon2 verifies a password against a hashed string
func (h *Hasher) CompareArgon2(hashedPassword, password string) (bool, error) {
	return argon2.VerifyEncoded([]byte(password), []byte(hashedPassword))
}

// HashMD5 returns the MD5 hash of a given message (not secure for passwords)
func (h *Hasher) HashMD5(msg string) string {
	hash := md5.New()
	hash.Write([]byte(msg))
	return hex.EncodeToString(hash.Sum(nil))
}

// EncryptMessage encrypts a message using AES-GCM with a provided secret key
func (h *Hasher) EncryptMessage(message string, secretKey []byte) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(message), nil)
	return hex.EncodeToString(ciphertext), nil
}

// DecryptMessage decrypts an AES-GCM encrypted string using the provided secret key
func (h *Hasher) DecryptMessage(encrypted string, secretKey []byte) (string, error) {
	data, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// Generate6DigitCode generates a random 6-digit numeric string
func (h *Hasher) Generate6DigitCode() (string, error) {
	max := big.NewInt(1000000) // 0 to 999999
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}
