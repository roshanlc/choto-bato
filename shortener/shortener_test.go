package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURL(t *testing.T) {

	initalLink := "https://roshanlc.github.io"
	shortLink, _ := GenerateShortURL([]byte(initalLink))
	assert.Equal(t, shortLink, "9ARbD34QrLw")
}
