package main

import (
    "D:\Desktop\myproject\football"
    "D:\Desktop\myproject\players"
    "fmt"
)

func main() {
    player1 := players.NewPlayer("Messi")
    player2 := players.NewPlayer("Ronaldo")

    team1 := players.NewTeam("Barcelona")
    team2 := players.NewTeam("Real Madrid")

    team1.AddPlayer(player1)
    team2.AddPlayer(player2)

    match := football.NewMatch(*team1, *team2)
    match.Start()
}
