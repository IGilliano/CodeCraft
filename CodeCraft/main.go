package main

import (
	"awesomeProject1/CodeCraft/Player"
	"fmt"
)

func main() {
	//player2 := Player.NewPlayer()
	fmt.Println("Hello and welcome to the CodeCraft! Place, where two of you could contest for... nothing but fun!")
	player1 := Player.NewPlayer()
	p1 := &player1
	shop := Player.NewShop()
	fmt.Println("What do you want to do?\n 1: Go to shop\n 2: Ready for fight")
	r := player1.IsReady()
	for r == true {
		r = shop.SelectUnit(p1)
	}
}
