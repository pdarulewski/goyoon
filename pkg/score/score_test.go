package score_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pdarulewski/goyoon/pkg/score"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		actual := score.New("test")

		const loops = 6

		correct := false
		entries := make(score.Log, loops)

		for i := range loops {
			entry := score.LogEntry{
				ID:         i + 1,
				Speed:      time.Duration((loops - i)) * time.Second,
				WasCorrect: correct,
			}

			actual.Update(entry)
			entries[i] = entry

			correct = !correct
		}

		assert.Equal(t, "test", actual.ItemID)
		assert.Equal(t, entries, actual.Log)
		assert.Equal(t, loops, actual.Samples)
		assert.Equal(t, loops/2, actual.Errors)

		assert.InEpsilon(t, 0.5, actual.MeanAccuracy, 0)
		assert.InEpsilon(t, 0.7376, actual.AccuracyConfidence, 0.0001)

		assert.Equal(t, (5+3+1)*time.Second/3, actual.MeanSpeed)
		assert.InEpsilon(t, 4.392*float64(time.Second), actual.SpeedConfidence, 0)
	})
}
