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
	if player1.Ready() == true {
		player1.Hire(shop)
		fmt.Println("Wanna hire more?")
	} else {
		fmt.Println("Okay")
	}
}
