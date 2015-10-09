package tennis
import (
	"testing"
)

func TestScoring(t *testing.T) {
	eivind := "Eivind"
	thomas := "Thomas"
	s := NewScoreboard(eivind, thomas)

	s.Score(eivind)

	ei, th := s.Scores()
	if th != 0 {
		t.Errorf("%v should have score %v, but was %v", thomas, 0, th)
	}

	if ei != 15 {
		t.Errorf("%v should have score %v, but was %v", eivind, 15, ei)
	}
}
