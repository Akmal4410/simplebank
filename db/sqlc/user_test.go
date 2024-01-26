package db

import (
	"context"
	"testing"
	"time"

	"github.com/akmal4410/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPasswod, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPasswod)

	arg := CreateUserParams{
		UserName:       util.RandomOwner(),
		FirstName:      util.RandomOwner(),
		HashedPassword: hashedPasswod,
		Email:          util.RandomEmail(),
	}

	user, err := testQuires.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	getUser, err := testQuires.GetUser(context.Background(), user.UserName)

	require.NoError(t, err)
	require.NotEmpty(t, getUser)

	require.Equal(t, user.UserName, getUser.UserName)
	require.Equal(t, user.FirstName, getUser.FirstName)
	require.Equal(t, user.HashedPassword, getUser.HashedPassword)
	require.Equal(t, user.Email, getUser.Email)

	require.WithinDuration(t, user.PasswordChangedAt, getUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, getUser.CreatedAt, time.Second)
}
