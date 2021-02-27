package client

import (
	"fmt"
	"os"
)

func saveHelp() {
	help := `
Options:
  -i, --id          Video ID
  -l, --language    Language code in which you want to store the transcript
  -o, --output      Filename in which the data will be stored
`
	printMsg(help)
}

func fetchHelp() {
	help := `
Options:
  -i, --id          Video ID
  -l, --language    Language code in which you want to search for the transcript
`
	printMsg(help)
}

func listHelp() {
	help := `
Options:
  -i, --id          Video ID
`
	printMsg(help)
}

func printMsg(help string) {
	fmt.Fprintf(os.Stderr, "Usage of: %s %s [OPTIONS]\n%s", os.Args[0], os.Args[1], help)
}

