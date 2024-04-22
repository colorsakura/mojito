package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/colorsakura/mojito/internal/app"
	"github.com/colorsakura/mojito/internal/config"
	"github.com/colorsakura/mojito/internal/util/version"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	showVersion bool
	debugMode   bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file of mojito")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "version of mojito")
	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "debug mode for mojito")
}

var rootCmd = &cobra.Command{
	Use:   "mojito",
	Short: "Mojito is a tui music player.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Println(version.Full())
			return nil
		}
		if debugMode {
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				fmt.Println("fatal:", err)
				os.Exit(1)
			}
			defer f.Close()
		}
		// TODO: implement load cfg function
		var (
			cfg *config.Config
			err error
		)
		if cfgFile != "" {
			cfg, err = config.LoadConfig(cfgFile)
			if err != nil {
				return err
			}
		} else {
			cfg = config.DefaultConfig()
		}

		if err := app.Init(cfg); err != nil {
			return err
		}

		app.Run()

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
