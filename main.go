package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name string
	Pos  int
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GetPos() int {
	return p.Pos
}

func (p *Player) Move() {
	rand.Seed(time.Now().UnixNano())
	p.Pos += rand.Intn(5)
}

type Board struct {
	Player   *Player
	Computer *Player
}

func (b *Board) PrtBoard() {
	PPos := b.Player.GetPos()
	CPos := b.Computer.GetPos()

	fmt.Print("\033[H\033[2J")

	fmt.Print(b.Player.GetName(), " : ")
	for i := 0; i <= 30; i++ {
		if i == PPos {
			fmt.Print("♟️ ")
		} else {
			fmt.Print("⬛")
		}
	}
	fmt.Println(" 🎈")

	fmt.Print(b.Computer.GetName(), " : ")
	for i := 0; i <= 30; i++ {
		if i == CPos {
			fmt.Print("♟️ ")
		} else {
			fmt.Print("⬛")
		}
	}
	fmt.Println(" 🎈")
}

func (b *Board) Evaluate() bool {
	if b.Player.GetPos() >= 30 {

		fmt.Print("\033[H\033[2J")

		fmt.Println("Player Win!")
		return false
	} else if b.Computer.GetPos() >= 30 {

		fmt.Print("\033[H\033[2J")

		fmt.Println("Computer Win!")
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Print("Press Enter to start the game....")
	fmt.Scanln()
	fmt.Print("Plase enter your name: ")
	var name string
	fmt.Scanln(&name)
	player := Player{Name: name}
	computer := Player{Name: "computer"}

	board := Board{Player: &player, Computer: &computer}

	board.PrtBoard()

	for board.Evaluate() {
		fmt.Print("Press Enter to Next turn...")
		fmt.Scanln()
		player.Move()
		board.PrtBoard()

		if !board.Evaluate() {
			break
		}
		fmt.Print("Press Enter to Next turn...")
		fmt.Scanln()

		computer.Move()
		board.PrtBoard()
	}
}
