package hash

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateHash(link string, id string) string {
	data := link + id

	hash := sha256.Sum256([]byte(data))

	encodedHash := base64.URLEncoding.EncodeToString(hash[:])

	return encodedHash[:8]
}
