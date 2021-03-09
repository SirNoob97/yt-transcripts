package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SirNoob97/yt-transcripts/cli"
	"github.com/SirNoob97/yt-transcripts/client"
	"github.com/SirNoob97/yt-transcripts/transcript"
)

// LDFLAGS
var (
	Version string
	Appname string
)

var (
	helpLong     = flag.Bool("help", false, "")
	helpShort    = flag.Bool("h", false, "")
	versionLong  = flag.Bool("version", false, "")
	versionShort = flag.Bool("v", false, "")
)

func main() {
	flag.Parse()
	hc := client.NewHTTPClient()
	t := transcript.NewTrasncript(hc)
	c := cli.NewFetcherClient(t)
	s := cli.NewSwitch(Appname, Version, c)

	if *helpLong || *helpShort || len(os.Args) == 1 {
		s.Help()
		return
	}

	if *versionLong || *versionShort {
		s.Info()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error %v\n", err)
		os.Exit(1)
	}
}
