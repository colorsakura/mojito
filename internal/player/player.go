package player

import "github.com/colorsakura/mojito/internal/music"

type Player struct {
	State int // player state
	playerList
	Direction   bool //playerlist direction
	EnableLyric bool
}

type playerList struct {
	Lists  music.List
	cursor int
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Play() {}

func (p *Player) Pause() {}

func (p *Player) LoadList() {}

func (p *Player) ClearList() {}
