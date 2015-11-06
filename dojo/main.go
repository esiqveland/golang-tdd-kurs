package main
import "fmt"

func main() {
	addOnePointToThisScore(0, 0)
}

const Win = -1

type scoreboard struct {
	player1 int
	player2 int
}

func (s *scoreboard) ScorePlayer1() (p1, p2 int) {

}

func (s *scoreboard) ScorePlayer2() (p1, p2 int) {

}
func addOnePointToThisScore(score1, score2 int) (p1, p2 int) {
	updatedScore := score1 + 15
	if score1 == 30 {
		return 40, score2
	}
	if score1 == 40 {
		return Win, score2
	}
	fmt.Printf("Will now add 1 point to this score '%v'. The new score is '%v'", score1, updatedScore)
	return updatedScore, score2
}

