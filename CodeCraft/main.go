package main

import (
	"awesomeProject1/CodeCraft/Player"
	"fmt"
)

func main() {
	//player2 := Player.NewPlayer()

	fmt.Println("Hello and welcome to the CodeCraft! Place, where two of you could contest for... nothing but fun!")
	player1 := Player.NewPlayer()
	shop := Player.NewShop()
	r := player1.Ready()
	for r == true {
		player1.Hire(shop)
	}
}
