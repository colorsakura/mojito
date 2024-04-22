package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type General struct {
	Download string `toml:"download"` // Download location
}

// Theme struct for tui color
type Theme struct {
	Foreground string `toml:"foreground"`
	Background string `toml:"background"`
}

type Window struct {
	Width  int `toml:"width"`
	Height int `toml:"height"`
}

type Music struct {
	Library string `toml:"library"`
}

type Player struct {
	Lyric *Lyric `toml:"lyric"`
}

type Lyric struct {
	Font   string `toml:"font"`
	Color  string `toml:"color"`
	Enable bool   `toml:"enable"`
}

type Config struct {
	General *General
	Window  *Window `toml:"window"`
	Theme   *Theme  `toml:"theme"`
	Music   *Music  `toml:"music"`
	Player  *Player `toml:"player"`
	Lyric   *Lyric  `toml:"lyric"`
}

func (c *Config) String() string {
	// TODO: implement Config struct string function
	return "config content"
}

func DefaultConfig() *Config {
	return &Config{
		General: &General{},
		Window:  &Window{},
		Theme: &Theme{
			Foreground: "#d8d8d8",
			Background: "#181818",
		},
		Music: &Music{
			Library: "~/Music",
		},
		Player: &Player{},
		Lyric:  &Lyric{},
	}
}

func UnmarshalConfig(buf []byte) (*Config, error) {
	cfg := DefaultConfig()
	if err := toml.Unmarshal(buf, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadConfig(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c, err := UnmarshalConfig(b)
	if err != nil {
		return nil, err
	}

	return c, nil
}
