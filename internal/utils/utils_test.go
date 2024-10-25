package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSOrX(t *testing.T) {
	require.Equal(t, "a", SOrX("a", "b"))
	require.Equal(t, "a", SOrX("a", ""))
	require.Equal(t, "b", SOrX("", "b"))
}
