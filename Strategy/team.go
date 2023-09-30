package players

import "fmt"

type Team struct {
	name    string
	players []Player
	score   int
	hasBall bool
}

func NewTeam(name string) *Team {
	return &Team{name, nil, 0, false}
}

func (t *Team) AddPlayer(player Player) {
	t.players = append(t.players, player)
}

func (t *Team) Pass() {
	fmt.Printf("%s passed the ball.\n", t.name)
	t.hasBall = true
}

func (t *Team) Score() {
	if t.hasBall {
		t.score++
		fmt.Printf("%s scored!\n", t.name)
		t.hasBall = false
	}
}
