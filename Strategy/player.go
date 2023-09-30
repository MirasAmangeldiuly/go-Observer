package players

import "fmt"

type Player struct {
	name string
}

func NewPlayer(name string) Player {
	return Player{name}
}

func (p Player) Kick() {
	fmt.Printf("%s kicked the ball.\n", p.name)
}
