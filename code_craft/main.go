package main

import (
	"awesomeProject1/code_craft/game_manager"
)

func main() {
	gm := game_manager.NewGameManager()
	p := gm.ChoosePlayers()
	for gm.Player1.Ready != true || gm.Player2.Ready != true {
		gm.MainMenu(p)
	}
	gm.Combat()

}
