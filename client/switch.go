package client

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
	client   TranscriptClient
	comands  map[string]func() func(string) error
}

// NewSwitch ...
func NewSwitch() Switch {
	tClient := NewClient()
	s := Switch{
		client:   tClient,
	}

	s.comands = map[string]func() func(string) error{
		"save":  s.save,
		"list":  s.list,
		"fetch": s.fetch,
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
		return errors.New("Could not parse '"+cmd.Name()+"' command flags")
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
	var help string
	for name := range s.comands {
		help += name + "\t --help\n"
	}

	fmt.Printf("Usage off: %s:\n <command> [<arguments>]\n%s", os.Args[0], help)
}

func (s Switch) saveFlags(f *flag.FlagSet) (*string, *string, *string) {
	i, l, o := "", "", ""

	f.StringVar(&i, "id", "", "Video ID")
	f.StringVar(&i, "i", "", "Video ID")
	f.StringVar(&l, "language", "", "Language code in which you want to store the transcript")
	f.StringVar(&l, "l", "", "Language code in which you want to store the transcript")
	f.StringVar(&o, "output", "", "Filename in which the data will be stored")
	f.StringVar(&o, "o", "", "Filename in which the data will be stored")

	return &i, &l, &o
}

func (s Switch) save() func(string) error {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		i, l, o := s.saveFlags(createCmd)

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
		ids := ""

		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		editCmd.StringVar(&ids, "id", "", "Video ID")

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(editCmd); err != nil {
			return err
		}

		res, err := s.client.List(ids)
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

	f.StringVar(&i, "id", "", "Video ID")
	f.StringVar(&i, "i", "", "Video ID")
	f.StringVar(&l, "language", "", "Language code in which you want to search for the transcript")
	f.StringVar(&l, "l", "", "Language code in which you want to search for the transcript")

	return &i, &l
}

func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		i, l := s.fetchFlags(fetchCmd)

		if err := s.checkArgs(1); err != nil {
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
