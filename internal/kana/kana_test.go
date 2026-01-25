package kana_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pdarulewski/phrame/internal/kana"
)

func TestToMap(t *testing.T) {
	t.Parallel()

	input := kana.Kana{
		{"あ", "a", kana.Basic},
		{"い", "i", kana.Basic},
	}

	t.Run("FromKanaToRomaji", func(t *testing.T) {
		t.Parallel()

		dict, err := input.ToDictionary(kana.FromKanaToRomaji)
		require.NoError(t, err)

		assert.Equal(t, kana.Dictionary{
			"あ": "a",
			"い": "i",
		}, dict)
	})

	t.Run("FromRomajiToKana", func(t *testing.T) {
		t.Parallel()

		dict, err := input.ToDictionary(kana.FromRomajiToKana)
		require.NoError(t, err)

		assert.Equal(t, kana.Dictionary{
			"a": "あ",
			"i": "い",
		}, dict)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		dict, err := input.ToDictionary(kana.Direction(-1))
		require.ErrorIs(t, err, kana.ErrKana)

		assert.Empty(t, dict)
	})
}
