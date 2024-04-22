package app

import (
	"fmt"
	"os"

	"github.com/colorsakura/mojito/internal/config"
	"github.com/colorsakura/mojito/internal/tui"
)

// Init app
func Init(c *config.Config) error {
	return nil
}

// Run app
func Run() {
	tui := tui.NewTui()
	if _, err := tui.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
