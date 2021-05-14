package access_token

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("Brand new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("newAccessToken shouldn't have a defined access token id")
	}

	if at.UserID != 0 {
		t.Error("new access token shouldn't  have an associated used id")
	}

	if at.ClientId != 0 {
		t.Error("new access token shouldn't  have an associated ClientId ")
	}

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("empty access token should be expired")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token created three hours from now should NOT be expired ")
	}
}
