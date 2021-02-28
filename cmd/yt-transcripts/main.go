package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SirNoob97/yt-transcripts/cli"
)

// LDFLAGS
var (
	Version string
	Appname string
)

func main() {
	v, h := flags()

	s := cli.NewSwitch(Appname, Version)

	if *h || len(os.Args) == 1 {
		s.Help()
		return
	}

	if *v {
		s.Info()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error %v\n", err)
		os.Exit(2)
	}
}

func flags() (*bool, *bool) {
	versionFlag, helpFlag := false, false
	flag.BoolVar(&helpFlag, "help", false, "")
	flag.BoolVar(&helpFlag, "h", false, "")
	flag.BoolVar(&versionFlag, "version", false, "")
	flag.BoolVar(&versionFlag, "v", false, "")
	flag.Parse()

	return &versionFlag, &helpFlag
}
