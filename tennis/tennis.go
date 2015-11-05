package tennis

import (
	"fmt"
	"errors"
)

func NewScoreboard(nameA, nameB string) Scoreboard {
	return &scoreboard{
		playerA: playerScore{name: nameA},
		playerB: playerScore{name: nameB},
	}
}

type Player interface {
	Name() string
	Score() int
}
type Scoreboard interface {
	Score(playerName string) error
	Scores() (int, int)
	Finished() bool
	Winner() (Player, error)
	String() string
}

func (p playerScore) Name() string {
	return p.name
}

func (p playerScore) Score() int {
	return p.score
}

func (s scoreboard) Winner() (Player, error) {
	if !s.finished {
		return nil, errors.New("no winner yet, we are not finished!")
	}
	return nil, errors.New("Not implemented yet")
}

func (s scoreboard) Finished() bool {
	return s.finished
}

func (s scoreboard) Scores() (int, int) {
	return s.playerA.score, s.playerB.score
}

func (s *scoreboard) Score(playerName string) error {
	if s.playerA.name == playerName {
		s.playerA.score += 15
		return nil
	}
	if s.playerB.name == playerName {
		s.playerB.score += 15
		return nil
	}
	return errors.New(fmt.Sprintf("no such player: %v", playerName))
}

type scoreboard struct {
	playerA  playerScore
	playerB  playerScore
	finished bool
}

func (s scoreboard) String() string {
	return fmt.Sprintf("%v %v - %v %v\n", s.playerA.name, s.playerA.score, s.playerB.score, s.playerB.name)
}

type playerScore struct {
	name  string
	score int
}
