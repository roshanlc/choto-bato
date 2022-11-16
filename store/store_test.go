package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initalLink := "https://roshanlc.github.io"

	shortURL := "9ARbD34QrLw"

	SaveURL(shortURL, initalLink)

	retrievedURL, _ := RetrieveURL(shortURL)

	assert.Equal(t, initalLink, retrievedURL)
}
