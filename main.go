package main

import (
	"github.com/gulien/orbit/commands"
	"os"

	jww "github.com/spf13/jwalterweatherman"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		jww.ERROR.Println(err)
		os.Exit(1)
	}
}
