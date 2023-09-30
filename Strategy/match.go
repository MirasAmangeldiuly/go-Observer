package football

import (
	"D:\Desktop\myproject\players"
	"fmt"
)

type Match struct {
	team1 players.Team
	team2 players.Team
}

func NewMatch(team1, team2 *players.Team) *Match {
	return &Match{*team1, *team2}
}

func (m *Match) Start() {
	fmt.Println("Match started!")
	m.team1.Pass()
	m.team2.Pass()
	m.team1.Score()
	m.team2.Score()
}
