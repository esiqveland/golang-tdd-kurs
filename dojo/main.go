package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)

	game := New()
	c := game.Scores()

	go printGame(4, c)

	game.ScorePlayer1()
	game.ScorePlayer1()
	game.ScorePlayer2()
	game.ScorePlayer2()
	game.ScorePlayer1()

	wg.Wait()
}

func printGame(i int, c chan game) {
	count := 0
	for g := range c {
		fmt.Printf("[logstate] %v\n", g)
		count = count + 1
		if count == i {
			wg.Done()
		}
	}
}
func New() TennisScorer {
	return &game{
		player1:  0,
		player2:  0,
		producer: make(chan game),
	}
}

type game struct {
	player1  int
	player2  int
	winner   string
	producer chan game
}

const Wins = -1
const Adv = -2

type TennisScorer interface {
	ScorePlayer1()
	ScorePlayer2()
	Winner() string
	Scores() chan game
	String() string
}

func (g *game) String() string {
	return fmt.Sprintf("%v %v", g.player1, g.player2)
}
func (g *game) Winner() string {
	return g.winner
}
func (g *game) logState() {
	newState := *g
	g.producer <- newState
}

func (game *game) Scores() chan game {
	return game.producer
}

func (game *game) checkWinner() {
	if game.player1 == Wins {
		game.winner = "player1"
	}
	if game.player2 == Wins {
		game.winner = "player2"
	}
}

func (game *game) ScorePlayer1() {
	game.player1, game.player2 = ScorePlayer(game.player1, game.player2)
	game.checkWinner()
	game.logState()
}

func (game *game) ScorePlayer2() {
	game.player2, game.player1 = ScorePlayer(game.player2, game.player1)
	game.checkWinner()
	game.logState()
}

func ScorePlayer(player int, otherPlayer int) (int, int) {
	switch player {
	case 0:
		player = 15
	case 15:
		player = 30
	case 30:
		player = 40
	case 40:
		if otherPlayer == 40 {
			player = Adv
		} else if otherPlayer == Adv {
			otherPlayer = 40
		} else {
			player = Wins
		}
	case Adv:
		player = Wins
	}

	return player, otherPlayer
}

func newGame() game {
	return game{}
}
