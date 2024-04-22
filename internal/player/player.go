package player

import (
	"fmt"
	"os/exec"

	"github.com/colorsakura/mojito/internal/music"
)

type Player struct {
	playerList
	State     int
	Direction bool
	Event     chan struct{}
}

type playerList struct {
	Lists  music.List
	cursor int
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Play() {
	go func() {
		command("/home/iFlygo/Music/Awaken.mp3")
	}()
}

func (p *Player) Pause() {}

func (p *Player) LoadList() {}

func (p *Player) ClearList() {}

func command(src string) {
	cmd := exec.Command("mpv", src, "--no-video")
	fmt.Println("play music")
	cmd.Run()
}
