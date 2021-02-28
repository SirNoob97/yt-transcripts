package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// TranscriptClient ...
type TranscriptClient interface {
	Save(id, language, fileName string) error
	List(id string) ([]string, error)
	Fetch(id, language string) (string, error)
}

// Switch ...
type Switch struct {
	client  TranscriptClient
	comands map[string]func() func(string) error
	version string
	appname string
}

// NewSwitch ...
func NewSwitch(version, appname string) Switch {
	tClient := NewClient()
	s := Switch{
		client: tClient,
	}

	s.comands = map[string]func() func(string) error{
		"save":    s.save,
		"list":    s.list,
		"fetch":   s.fetch,
		"version": s.info,
	}
	return s
}

// Switch ...
func (s Switch) Switch() error {
	cmdName := os.Args[1]
	cmd, ok := s.comands[cmdName]
	if !ok {
		return fmt.Errorf("Invalid Command '%s'", cmdName)
	}

	return cmd()(cmdName)
}

func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return errors.New("Could not parse '" + cmd.Name() + "' command flags")
	}
	return nil
}

func (s Switch) checkArgs(minArgs int) error {
	if (len(os.Args) == 3 && os.Args[2] == "--help") || (len(os.Args) == 3 && os.Args[2] == "-h") {
		return nil
	}

	if len(os.Args)-2 < minArgs {
		fmt.Printf("Incorrect use of %s\n%s %s --help\n", os.Args[1], os.Args[0], os.Args[1])
		return fmt.Errorf("%s expects at least %d arg(s), %d provided", os.Args[1], minArgs, len(os.Args)-2)
	}

	return nil
}

// Help ...
func (s Switch) Help() {
	help := `
Commands:
  save    Save the transcript to the specified file path.
  list    List available video transcripts.
  fetch   Fetch the transcript.

Options:
  --help, -h   Display command help message.
`
	fmt.Printf("Usage off: %s: [COMMAND] [OPTIONS]\n%s", os.Args[0], help)
}

func (s Switch) saveFlags(f *flag.FlagSet) (*string, *string, *string) {
	i, l, o := "", "", ""

	f.StringVar(&i, "id", "", "")
	f.StringVar(&i, "i", "", "")
	f.StringVar(&l, "language", "", "")
	f.StringVar(&l, "l", "", "")
	f.StringVar(&o, "output", "", "")
	f.StringVar(&o, "o", "", "")

	return &i, &l, &o
}

func (s Switch) save() func(string) error {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		i, l, o := s.saveFlags(createCmd)
		createCmd.Usage = saveHelp

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(createCmd); err != nil {
			return err
		}

		err := s.client.Save(*i, *l, *o)
		if err != nil {
			return errors.New("Could not save transcript")
		}

		fmt.Println("Transcript save successfully")
		return nil
	}
}

func (s Switch) list() func(string) error {
	return func(cmd string) error {
		id := ""

		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		editCmd.StringVar(&id, "i", "", "")
		editCmd.StringVar(&id, "id", "", "")
		editCmd.Usage = listHelp

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(editCmd); err != nil {
			return err
		}

		res, err := s.client.List(id)
		if err != nil {
			return errors.New("Could not find transcripts")
		}

		for _, r := range res {
			fmt.Println(r)
		}
		return nil
	}
}

func (s Switch) fetchFlags(f *flag.FlagSet) (*string, *string) {
	i, l := "", ""

	f.StringVar(&i, "id", "", "")
	f.StringVar(&i, "i", "", "")
	f.StringVar(&l, "language", "", "")
	f.StringVar(&l, "l", "", "")

	return &i, &l
}

func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		i, l := s.fetchFlags(fetchCmd)
		fetchCmd.Usage = fetchHelp

		if err := s.checkArgs(2); err != nil {
			return err
		}

		if err := s.parseCmd(fetchCmd); err != nil {
			return err
		}

		res, err := s.client.Fetch(*i, *l)
		if err != nil {
			return errors.New("Could not fetch transcript")
		}

		fmt.Println(res)
		return nil
	}
}

func (s Switch) info() func(string) error {
	return func(cmd string) error {
		if err := s.checkArgs(0); err != nil {
			return err
		}

		fmt.Printf("%s %s", s.appname, s.version)
		return nil
	}
}
