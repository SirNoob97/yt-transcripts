package cli

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"
	"unicode/utf8"

	"github.com/SirNoob97/yt-transcripts/transcript"
	"github.com/pasdam/mockit/matchers/argument"
	"github.com/pasdam/mockit/mockit"
)

func Test_NewSwitch(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	s := NewSwitch(appname, version)

	if len(s.comands) == 0 {
		t.Fatalf("Expected a non-empty commands map")
	}

	if s.appname != appname && s.version != version {
		t.Fatalf("Expected appname %s, got %s\n Expected version %s got %s", appname, s.appname, version, s.version)
	}
}

func Test_Switch_WithAValidCommand(t *testing.T) {
	save := func() func(string) error {
		return func(save string) error {
			return nil
		}
	}
	const appname, version = "TESTING", "0.0.0"
	s := NewSwitch(appname, version)
	s.comands["save"] = save
	os.Args[1] = "save"

	err := s.Switch()
	if err != nil {
		t.Fatalf("Expected nil error, got %s\n", err)
	}
}

func Test_Switch_WithAnInvalidCommand(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	s := NewSwitch(appname, version)
	os.Args[1] = appname

	err := s.Switch()
	if err == nil {
		t.Fatal("Expected and invalid command error, got nil\n")
	}
}

func Test_parseCmd(t *testing.T) {
	const appname, command, testFlag, arg = "TESTING", "COMMAND", "-i", "TEST"
	os.Args = []string{appname, command, testFlag, arg}
	cmd := flag.NewFlagSet(command, flag.ExitOnError)
	cmd.String("i", "", "")
	s := Switch{}

	err := s.parseCmd(cmd)
	if err != nil {
		t.Fatalf("Expected nil error, got %s\n", err)
	}
}

func Test_parseCmd_WithInvalidFlag(t *testing.T) {
	const appname, command, testFlag, arg = "TESTING", "COMMAND", "-e", "-1"
	os.Args = []string{appname, command, testFlag, arg}
	cmd := flag.NewFlagSet(command, flag.ContinueOnError)
	cmd.String("i", "", "")
	s := Switch{}

	err := s.parseCmd(cmd)
	if err == nil {
		t.Fatal("Expected and invalid flag err got nil\n")
	}
}

func Test_checkCommandArgs(t *testing.T) {
	const appname, command, testFlag, arg = "TESTING", "COMMAND", "-e", "-1"
	os.Args = []string{appname, command, testFlag, arg}
	s := Switch{}
	err := s.checkCommandArgs(1)
	if err != nil {
		t.Fatalf("Expected nil error, got %s", err)
	}
}

func Test_checkCommandArgs_WithHFlag(t *testing.T) {
	const appname, command, testFlag = "TESTING", "COMMAND", "-h"
	os.Args = []string{appname, command, testFlag}
	s := Switch{}
	err := s.checkCommandArgs(1)
	if err != nil {
		t.Fatalf("Expected nil error, got %s", err)
	}
}

func Test_checkCommandArgs_WithHelpFlag(t *testing.T) {
	const appname, command, testFlag = "TESTING", "COMMAND", "--help"
	os.Args = []string{appname, command, testFlag}
	s := Switch{}
	err := s.checkCommandArgs(1)
	if err != nil {
		t.Fatalf("Expected nil error, got %s", err)
	}
}

func Test_checkCommandArgs_With0Args(t *testing.T) {
	const appname, command, testFlag = "TESTING", "COMMAND", "-flag"
	os.Args = []string{appname, command, testFlag}
	s := Switch{}
	err := s.checkCommandArgs(2)
	if err == nil {
		t.Fatal("Expected and error message, got nil")
	}
}

func Test_Help(t *testing.T) {
	stderr := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("PIPE error: %s", err)
	}
	os.Stderr = w

	s := Switch{}
	s.Help()

	w.Close()

	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("READ error: %s", err)
	}
	os.Stdout = stderr

	if count := utf8.RuneCountInString(string(out)); count == 0 {
		t.Fatal("Expected an error message, got nothing")
	}
}

func Test_Info(t *testing.T) {
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("PIPE error: %s", err)
	}
	os.Stdout = w

	const appname, version = "TESTING", "0.0.0"
	s := Switch{appname: appname, version: version}
	s.Info()

	w.Close()

	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("READ error: %s", err)
	}
	os.Stdout = stdout

	if count := utf8.RuneCountInString(string(out)); count == 0 {
		t.Fatal("Expected an the appname and the version, got nothing")
	}
}

func Test_saveFlags(t *testing.T) {
	const appname, command = "TESTING", "COMMAND"
	cmd := flag.NewFlagSet(command, flag.ExitOnError)
	s := Switch{}

	i, l, o := s.saveFlags(cmd)

	if *i != "" {
		t.Fatalf("Expected an empty string as i value, got %v", *i)
	}
	if *l != "" {
		t.Fatalf("Expected an empty string as i value, got %v", *l)
	}
	if *o != "" {
		t.Fatalf("Expected an empty string as i value, got %v", *o)
	}
}

func Test_fetchFlags(t *testing.T) {
	const appname, command = "TESTING", "COMMAND"
	cmd := flag.NewFlagSet(command, flag.ExitOnError)
	s := Switch{}

	i, l := s.fetchFlags(cmd)

	if *i != "" {
		t.Fatalf("Expected an empty string as i value, got %v", *i)
	}
	if *l != "" {
		t.Fatalf("Expected an empty string as i value, got %v", *l)
	}
}

func Test_save(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	const command, flagI, flagL, flagO, arg = "save", "-i", "-l", "-o", "ARG"
	os.Args = []string{appname, command, flagI, arg, flagL, arg, flagO, arg}

	s := NewSwitch(appname, version)
	f := mockit.MockFunc(t, s.client.Fetch)
	f.With(arg, arg).Return(appname, nil)

	tr := mockit.MockFunc(t, transcript.FetchTranscript)
	tr.With(argument.Any, argument.Any, argument.Any).Return(transcript.Transcript{})

	fi := mockit.MockFunc(t, os.OpenFile)
	fi.With(argument.Any, argument.Any, argument.Any)

	sa := mockit.MockFunc(t, s.client.Save)
	sa.With(argument.Any, argument.Any, argument.Any).Return(nil)

	err := s.Switch()
	if err != nil {
		t.Fatalf("Expected nil error got, %v\n", err)
	}
}
