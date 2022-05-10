package main

import (
	"awesomeProject1/CodeCraft/Player"
	"awesomeProject1/CodeCraft/game_manager"
)

func main() {
	gm := game_manager.NewGameManager()
	p1 := Player.NewPlayer()
	p := &p1
	gm.Menu(p)
}
