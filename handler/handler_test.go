package handler

import (
	"context"
	"math/rand"
	"testing"
	"time"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserMetaService"
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

func TestHandler(t *testing.T) {
	assert := assert.New(t)
	ctx := context.TODO()
	h := New()

	// Testing Add.
	userID := randStr(10)
	identifier := proto.Identifier{UserID: userID}
	_, err := h.Add(ctx, &identifier)
	assert.NoError(err)

	// Testing Get.
	data, err := h.Get(ctx, &identifier)
	assert.NotZero(data.AccountStatus)
	assert.NotZero(data.AccountCreationUTC)
	assert.NoError(err)

	// Testing Activate.
	_, err = h.Activate(ctx, &identifier)
	assert.NoError(err)
	data, _ = h.Get(ctx, &identifier)
	assert.Equal(verified, data.AccountStatus)

	// Testing Delete.
	_, err = h.Delete(ctx, &identifier)
	assert.NoError(err)
	data, _ = h.Get(ctx, &identifier)
	assert.Equal(deleted, data.AccountStatus)
}
