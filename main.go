package main

import (
	"github.com/kabotnik/white-elephant/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "white-elephant",
	Short: "Plays a game of white elephant",
}

func main() {
	cmd.Execute()
}
