package meta

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randStr(length int) string {
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
	ctx := context.TODO()
	m := New()

	// Testing Create.
	userID := randStr(10)
	assert.NoError(m.Create(ctx, userID))

	// Testing Get.
	data, err := m.Get(ctx, userID)
	assert.NotZero(data.AccountStatus)
	assert.NotZero(data.AccountCreationUTC)
	assert.NoError(err)

	// Testing ChangeStatus.
	assert.NoError(m.SetStatus(ctx, userID, "something"))
	data, _ = m.Get(ctx, userID)
	assert.Equal("something", data.AccountStatus)
}
