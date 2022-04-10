package Player

import (
	"awesomeProject1/CodeCraft/Units"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	Gold int64
	Army []Units.Unit
	Name string
}

func (p *Player) ChooseName() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("May I know your name?")
	scanner.Scan()
	p.Name = scanner.Text()
	fmt.Println("Greetings,", p.Name, "!")

}

func NewPlayer() Player {
	p := Player{100, make([]Units.Unit, 0), ""}
	p.ChooseName()
	return p

}

func (Player) Ready() bool {
	fmt.Println("What do you want to do?\n 1: Go to shop\n 2: Ready for fight")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	if i == 1 {
		return true
	} else if i == 2 {
		return false
	}
	return false
}
