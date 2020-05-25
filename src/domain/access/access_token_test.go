package access

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccesTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
}

func TestAccesTokenConstantsIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.False(t, at.IsExpired(), "empty access token should be expired")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")
}
