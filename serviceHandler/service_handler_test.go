package serviceHandler

import (
	"context"
	"math/rand"
	"testing"
	"time"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserMetaService"
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
	handler := New()
	ctx := context.TODO()

	// Testing Add
	userID := generateRandStr(10)
	identifier := proto.Identifier{UserID: userID}
	handler.Add(ctx, &identifier)

	// Testing Get
	data, _ := handler.Get(ctx, &identifier)
	assert.NotZero(data.AccountStatus)
	assert.NotZero(data.AccountCreationUTC)

	// Testing Activate
	handler.Activate(ctx, &identifier)
	data, _ = handler.Get(ctx, &identifier)
	assert.Equal(accountStatuses.Str("verified"), data.AccountStatus)

	// Testing Delete
	handler.Delete(ctx, &identifier)
	data, _ = handler.Get(ctx, &identifier)
	assert.Equal(accountStatuses.Str("deleted"), data.AccountStatus)
}
