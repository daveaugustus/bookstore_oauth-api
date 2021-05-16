package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstent(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "Brand new access token should not be expired")

	assert.EqualValues(t, "", at.AccessToken, "newAccessToken shouldn't have a defined access token id")
	assert.True(t, at.UserID == 0, "new access token shouldn't  have an associated used id")
	assert.EqualValues(t, 0, at.ClientId, "new access token shouldn't  have an associated ClientId")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "empty access token should be expired")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.False(t, at.IsExpired(), "access token created three hours from now should NOT be expired")
}
