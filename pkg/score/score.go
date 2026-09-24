// score measures an item's metrics using an exponential moving average.
package score

import "time"

// Alpha is a parameter for exponential moving average.
const (
	Alpha       float64       = 0.2
	TargetSpeed time.Duration = 900 * time.Millisecond
)

// LogEntry is an append-only log entry.
type LogEntry struct {
	ID         int
	Speed      time.Duration
	WasCorrect bool
}

// Log is an implementation of an append-only log.
type Log []LogEntry

// Score stores metrics for a single item.
type Score struct {
	ItemID string

	Log Log

	Samples int
	Errors  int

	MeanSpeed    time.Duration
	MeanAccuracy float64

	SpeedConfidence    float64
	AccuracyConfidence float64
}

func New(id string) *Score {
	return &Score{
		ItemID:             id,
		Log:                Log{},
		Samples:            0,
		Errors:             0,
		MeanSpeed:          0,
		MeanAccuracy:       0,
		SpeedConfidence:    0,
		AccuracyConfidence: 0,
	}
}

// GetConfidence returns the final confidence metrics which is used to unlocking next items.
func (s *Score) GetConfidence() float64 {
	if s.Samples == 0 {
		return 0
	}

	speed := float64(TargetSpeed) / s.SpeedConfidence
	if speed > 1 {
		speed = 1
	}

	return speed * s.AccuracyConfidence
}

// Update updates the score based on a new LogEntry.
func (s *Score) Update(entry LogEntry) {
	s.Log = append(s.Log, entry)
	s.Samples++

	hit := 1.0

	if !entry.WasCorrect {
		s.Errors++
		hit = 0.0
	}

	if s.Samples == 1 {
		s.MeanAccuracy = hit
		s.AccuracyConfidence = 1.0

		s.MeanSpeed = entry.Speed
		s.SpeedConfidence = float64(entry.Speed)

		return
	}

	s.MeanAccuracy = (s.MeanAccuracy*float64(s.Samples-1) + hit) / float64(s.Samples)
	s.AccuracyConfidence = Alpha*hit + (1-Alpha)*s.AccuracyConfidence

	if !entry.WasCorrect {
		return
	}

	s.MeanSpeed = time.Duration(
		(float64(s.MeanSpeed)*float64(s.Samples-s.Errors-1) + float64(entry.Speed)) / float64(
			s.Samples-s.Errors,
		),
	)
	s.SpeedConfidence = Alpha*float64(entry.Speed) + (1-Alpha)*s.SpeedConfidence
}
