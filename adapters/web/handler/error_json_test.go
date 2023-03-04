package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrorJson(t *testing.T) {
	t.Run("Should return a json error", func(t *testing.T) {
		got := jsonError("Error")
		expected := []byte("{\"message\":\"Error\"}")

		require.Equal(t, expected, got)
	})
}
