package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SirNoob97/yt-transcripts/cli"
)

func main() {
	helpFlag := flag.Bool("help", false, "Display help message")
	flag.Parse()
	s := client.NewSwitch()

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
