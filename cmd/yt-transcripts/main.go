package main

import (
	"flag"
	"fmt"
	"os"

	cli "github.com/SirNoob97/yt-transcripts/cli"
)

// LDFLAGS
var (
	Version string
	AppName string
)

func main() {
	helpFlag := flag.Bool("help", false, "Display help message")
	flag.Parse()
	s := cli.NewSwitch(Version, AppName)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error %v\n", err)
		os.Exit(2)
	}
}
