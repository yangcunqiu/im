package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"strings"
)

func EncodeBySHA256(plaintext string, salt string) string {
	h := sha256.New()
	h.Write([]byte(plaintext))
	return hex.EncodeToString(h.Sum([]byte(salt)))
}

func VailPasswordBySHA256(plaintext string, sourceCiphertext string, salt string) (ok bool) {
	h := sha256.New()
	h.Write([]byte(plaintext))
	ciphertext := hex.EncodeToString(h.Sum([]byte(salt)))
	return ciphertext == sourceCiphertext
}

func GetUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
