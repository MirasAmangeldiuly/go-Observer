package main

func main() {
	weapon := NewWeapon("Меч")
	player1 := NewPlayer("Игрок1")
	player2 := NewPlayer("Игрок2")

	weapon.Register(player1)
	weapon.Register(player2)

	weapon.UpdateAvailability()
}
