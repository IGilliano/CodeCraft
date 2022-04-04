package main

import (
	"awesomeProject1/CodeCraft/Player"
	"fmt"
)

func Ready() bool {
	var n int64
	fmt.Println("What do you want to do?\n 1: Go to shop\n 2: Ready for fight")
	fmt.Scanf("%d", &n)
	if n == 1 {
		return false
	} else if n == 2 {
		return true
	}
	return false
}

func main() {
	player1 := Player.NewPlayer()
	//player2 := Player.NewPlayer()
	player1.Hire()
	/*
		fmt.Println("Hello and welcome to the CodeCraft! Place, where two of you could contest for... nothing but fun!")
		player1.ChooseName()
		Ready()
		fmt.Println("22")*/
}
