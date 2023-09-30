package main

import "fmt"

type Player struct {
	name string
}

func NewPlayer(name string) *Player {
	return &Player{name: name}
}

func (p *Player) Update(weaponName string) {
	fmt.Printf("Игрок %s может получить новое оружие %s\n", p.name, weaponName)
}
