package music

import (
	"time"
)

type Music struct {
	Name     string
	Path     string // Music path
	Lyric    string // Lyric path
	Time     time.Duration
	DateTime time.Time
}

type List []Music

func parseMusic(p string) (Music, error) {
	var music Music

	return music, nil
}

func ParseMusicLists(p string) (*List, error) {

	return nil, nil
}
