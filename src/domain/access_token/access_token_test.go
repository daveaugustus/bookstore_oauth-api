package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstents(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "Brand new access token shouldn't be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token shouldn't have defined access token id")
	assert.True(t, at.UserId == 0, "new access token shouldn't have en associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.False(t, at.IsExpired(), "Access token expiring three hours from now shouldn't be expired!")

}
