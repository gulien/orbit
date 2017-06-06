package main

import (
	"github.com/gulien/orbit/commands"
	"github.com/gulien/orbit/notifier"
)

func main() {
	notifier.Start()

	if err := commands.RootCmd.Execute(); err != nil {
		notifier.Error(err)
	}
}
