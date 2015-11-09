package main

import (
	"testing"
)


func TestPlayer1ScoreTheFirstPoint(t *testing.T) {
	expected := game{player1: 15, player2: 0}
	game := game{}
	game.ScorePlayer1()

	testPlayerScores(t, game, expected)
}

func assert(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("error expected '%v' but was '%v'", expected, actual)
	}
}

func testPlayerScores(t *testing.T, expected, actual game) {
	if expected.player1 != actual.player1 {
		t.Errorf("error expected player1 '%v' but was '%v'", expected, actual)
	}
	if expected.player2 != actual.player2 {
		t.Errorf("error expected player2 '%v' but was '%v'", expected, actual)
	}
}

func TestPlayer1ScoresFrom15(t *testing.T) {
	expected := game{player1: 30}

	game := game{player1:15}
	game.ScorePlayer1()

	testPlayerScores(t, expected, game)
}

func TestPlayer1ScoresFrom30(t *testing.T) {
	expected := game{player1: 40}

	game := game{player1:30}
	game.ScorePlayer1()

	testPlayerScores(t, expected, game)
}

func TestPlayer2ScoresFrom30(t *testing.T) {
	expected := game{player2: 40}

	game := game{player2:30}
	game.ScorePlayer2()

	testPlayerScores(t, expected, game)
}

func TestPlayer2WinsTheGameIfScoresFrom40(t *testing.T) {
	game := game{player2: 40}
	expected := "player2"

	game.ScorePlayer2()

	actual := game.winner
	assert(t, expected, actual)
}

func TestGivenPlayer1AndPlayer2HasScore40_WhenPlayer1Score_ThenNoWinner(t *testing.T) {
	game := game{player2: 40, player1: 40}
	expected := ""

	game.ScorePlayer1()

	actual := game.winner
	assert(t, expected, actual)
}

func TestGivenPlayer1AndPlayer2HasScore40_WhenPlayer1Scores_ThenPlayer1GetsAdv(t *testing.T) {
	setup := game{player1: 40, player2: 40}
	expected := game{player1: Adv, player2: 40}

	setup.ScorePlayer1()

	testPlayerScores(t, expected, setup)
}

func TestGivenPlayer1AndPlayer2HasScore40_WhenPlayer2Scores_ThenPlayer2GetsAdv(t *testing.T) {
	setup := game{player1: 40, player2: 40}
	expected := game{player1: 40, player2: Adv}

	setup.ScorePlayer2()

	testPlayerScores(t, expected, setup)
}

func TestGivenPlayer1AndPlayer2HasScore40_WhenPlayer2ScoresTwice_ThenPlayer2Wins(t *testing.T) {
	setup := game{player1: 40, player2: 40}
	expected := game{player1: 40, player2: Wins}

	setup.ScorePlayer2()
	setup.ScorePlayer2()

	testPlayerScores(t, expected, setup)
}

func TestGivenPlayer1AndPlayer2HasScore40_WhenPlayer2Score_ThenNoWinner(t *testing.T) {
	game := game{player2: 40, player1: 40}
	expected := ""

	game.ScorePlayer2()

	actual := game.winner
	assert(t, expected, actual)
}

func TestPlayer2DoesntWinTheGameIfScoresNotFrom40(t *testing.T) {
	game := game{player2: 30}
	expected := ""

	game.ScorePlayer2()

	actual := game.winner
	assert(t, expected, actual)
}


func TestPlayer1WinsTheGameIfScoresFrom40(t *testing.T) {
	game := game{player1: 40}
	expected := "player1"

	game.ScorePlayer1()

	actual := game.winner
	assert(t, expected, actual)
}

func TestPlayer1Does_NotWinTheGameIf_ScoresFrom30(t *testing.T) {
	game := game{player1: 30}
	expected := ""

	game.ScorePlayer1()

	actual := game.winner
	assert(t, expected, actual)
}

func TestGivenPlayer1HasAdvantage_WhenPlayer2Scores_ThenPlayer1AndPlayer2Has40(t *testing.T) {
	setup := game{player1: Adv, player2: 40}
	expected := game{player1: 40, player2: 40}

	setup.ScorePlayer2()

	testPlayerScores(t, expected, setup)
}

func TestGivenPlayer2HasAdvantage_WhenPlayer1Scores_ThenPlayer1AndPlayer2Has40(t *testing.T) {
	setup := game{player1: 40, player2: Adv}
	expected := game{player1: 40, player2: 40}

	setup.ScorePlayer1()

	testPlayerScores(t, expected, setup)
}

func TestStartMatch(t *testing.T) {
	expected := game{player1:0, player2:0}
	result := newGame()

	if expected != result {
		t.Errorf("error expected '%v' but was '%v'", expected, result)
	}
}

