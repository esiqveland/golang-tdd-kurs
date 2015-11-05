package main

import "fmt"

// START OMIT
type Message struct {
	ID int
}

func (m *Message) Inc() {
	m.ID = m.ID + 1
}
func main() {
	m := Message{ID: 1}
	m.Inc()
	m.Inc()
	fmt.Printf("%+v", m) // Hvorfor?
}
// END OMIT

func tempFunc() {
	fmt.Print("")
}