package hasher

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
)

// CompareToSha256Hash returns true if Sha256 hash of the test equals to given hash
func CompareToSha256Hash(test string, hash string) (bool, error) {
	sha256Sum, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	inHash32 := sha256.Sum256([]byte(test))
	return subtle.ConstantTimeCompare(inHash32[:], sha256Sum) == 1, nil
}
