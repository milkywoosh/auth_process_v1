package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"oraluke.com/conn-ora1/util"
)

func TestCreateToken(t *testing.T) {

	secretKeyRand := util.RandomString(10)

	// implement Tokenizer
	JWTSecret := NewJWTTOken(secretKeyRand)

	usernameVal := "john001"
	roleVal := "scmt_admin"
	var durationVal time.Duration = 5 * time.Minute

	before := time.Now()
	signedToken, payload, err := JWTSecret.CreateToken(usernameVal, roleVal, durationVal)
	require.NoError(t, err)
	require.NotEmpty(t, signedToken)

	require.NotEmpty(t, payload)

	require.IsType(t, "", signedToken)
	require.IsType(t, &Payload{}, payload)

	require.Equal(t, payload.Username, usernameVal)
	require.Equal(t, payload.Role, roleVal)

	expectedExpired := before.Add(durationVal)
	require.WithinDuration(t, expectedExpired, payload.ExpiredAt, 1*time.Second)

}
