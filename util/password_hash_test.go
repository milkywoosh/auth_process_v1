package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {

	myPw := "marine001"
	hashedPw, err := HashPassword(myPw)

	require.NoError(t, err)
	require.IsType(t, "", hashedPw)

	err = ComparePassword(hashedPw, myPw)
	if err != nil {
		t.Fatalf("password is wrong: %v", err)
	}

	require.NoError(t, err, "check if hashed password equal to myPw")

}
