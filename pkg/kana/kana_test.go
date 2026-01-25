package kana_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pdarulewski/goyoon/pkg/kana"
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
			"あ": kana.Mora{
				Character: "あ",
				Romaji:    "a",
				Type:      kana.Basic,
			},
			"い": kana.Mora{
				Character: "い",
				Romaji:    "i",
				Type:      kana.Basic,
			},
		}, dict)
	})

	t.Run("FromRomajiToKana", func(t *testing.T) {
		t.Parallel()

		dict, err := input.ToDictionary(kana.FromRomajiToKana)
		require.NoError(t, err)

		assert.Equal(t, kana.Dictionary{
			"a": kana.Mora{
				Character: "あ",
				Romaji:    "a",
				Type:      kana.Basic,
			},
			"i": kana.Mora{
				Character: "い",
				Romaji:    "i",
				Type:      kana.Basic,
			},
		}, dict)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		dict, err := input.ToDictionary(kana.Direction(-1))
		require.ErrorIs(t, err, kana.ErrKana)

		assert.Empty(t, dict)
	})
}
