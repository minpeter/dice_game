package Player

import (
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

func (p *Player) GetPost() int {
	return p.Pos
}

func (p *Player) Move() {
	rand.Seed(time.Now().UnixNano())
	p.Pos += rand.Intn(5)
}
