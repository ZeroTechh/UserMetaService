package metaDB

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// generateRandStr generates a random string
func generateRandStr(length int) string {
	charset := "1234567890abcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TestMetaDB(t *testing.T) {
	assert := assert.New(t)
	metaDB := New()

	// Testing Create
	userID := generateRandStr(10)
	metaDB.Create(userID)

	// Testing Get
	data := metaDB.Get(userID)
	assert.NotZero(data.AccountStatus)
	assert.NotZero(data.AccountCreationUTC)

	// Testing ChangeStatus
	metaDB.ChangeStatus(userID, "something")
	assert.Equal("something", metaDB.Get(userID).AccountStatus)
}
