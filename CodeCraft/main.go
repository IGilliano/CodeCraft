package main

import (
	"awesomeProject1/CodeCraft/Player"
	"fmt"
)

func main() {
	Player1 := Player.NewPlayer()
	//player2 := CodeCraft.NewPlayer()

	Player.Hire(&Player1)
	fmt.Println(Player1)
}
