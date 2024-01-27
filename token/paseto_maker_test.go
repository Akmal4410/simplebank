package token

import (
	"testing"
	"time"

	"github.com/akmal4410/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	userName := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(userName, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userName, payload.UserName)
	require.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiresAt.Time, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	userName := util.RandomOwner()
	duration := -time.Minute

	token, err := maker.CreateToken(userName, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrorExpiredToken.Error())
	require.Nil(t, payload)
}
