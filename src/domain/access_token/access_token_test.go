package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "New access token shouldn't be expired")
	assert.EqualValues(t, at.AccessToken, "", "The new access token shouldn't be nil")
	assert.True(t, at.UserId == 0, "", "New access token shouldn't have an associated userid")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Access token should be expired by default!")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "acess token expiring three hours from now, shouldn't be expired!")
}
