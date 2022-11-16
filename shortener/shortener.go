package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

func sha256Hash(data []byte) []byte {
	hash := sha256.New()

	hash.Write(data)

	return hash.Sum(nil)
}

func GenerateShortURL(data []byte) (string, error) {

	// Generate sha256 hash of provided data
	hash := sha256Hash(data)

	// Convert hash into a big.Int's Uint64
	big := new(big.Int).SetBytes(hash).Uint64()

	// Generate base58 encoding based on Flickr
	encoded, err := base58.FlickrEncoding.Encode([]byte(fmt.Sprintf("%d", big)))
	if err != nil {
		return "", nil
	}
	return string(encoded), nil

}
