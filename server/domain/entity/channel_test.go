package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChannel_HasValidPassword(t *testing.T) {
	c := NewChannel("daniel", "good-password", "daniel@gmail.com")
	valid := c.HasValidPassword("bad-password")
	require.False(t, valid)

	valid = c.HasValidPassword("good-password")
	require.True(t, valid)
}
