package tennis

import (
	"testing"
)

var game Scoreboard
const (
	eivind = "Eivind"
	thomas = "Thomas"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestScoringA(t *testing.T) {
	game := NewScoreboard(eivind, thomas)

	err := game.Score(eivind)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}

	ei, th := game.Scores()
	if th != 0 {
		t.Errorf("%v should have score %v, but was %v", thomas, 0, th)
	}

	if ei != 15 {
		t.Errorf("%v should have score %v, but was %v", eivind, 15, ei)
	}
}

func TestScoringB(t *testing.T) {
	game := NewScoreboard(eivind, thomas)

	err := game.Score(thomas)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}

	ei, th := game.Scores()
	if th != 15 {
		t.Errorf("%v should have score %v, but was %v", thomas, 15, th)
	}

	if ei != 0 {
		t.Errorf("%v should have score %v, but was %v", eivind, 0, ei)
	}
}
