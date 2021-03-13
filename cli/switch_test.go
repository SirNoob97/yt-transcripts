package cli

import (
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"testing"
	"unicode/utf8"

	mocks "github.com/SirNoob97/yt-transcripts/mocks/cli"
)

func TestNewSwitch(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	m := new(mocks.TranscriptClient)
	s := NewSwitch(appname, version, m)

	if len(s.comands) == 0 {
		t.Fatalf("Expected a non-empty commands map")
	}

	if s.appname != appname && s.version != version {
		t.Fatalf("Expected appname %s, got %s\n Expected version %s got %s", appname, s.appname, version, s.version)
	}
}

func TestSwitch_WithAValidCommand(t *testing.T) {
	save := func() func(string) error {
		return func(save string) error {
			return nil
		}
	}
	const appname, version = "TESTING", "0.0.0"
	m := new(mocks.TranscriptClient)
	s := NewSwitch(appname, version, m)
	s.comands["save"] = save
	os.Args[1] = "save"

	err := s.Switch()
	if err != nil {
		t.Fatalf("Expected nil error, got %s\n", err)
	}
}

func TestSwitch_WithAnInvalidCommand(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	m := new(mocks.TranscriptClient)
	s := NewSwitch(appname, version, m)
	os.Args[1] = appname

	err := s.Switch()
	if err == nil {
		t.Fatal("Expected and invalid command error, got nil\n")
	}
}

func TestParseCmd(t *testing.T) {
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

func TestParseCmd_WithInvalidFlag(t *testing.T) {
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

func TestCheckCommandArgs(t *testing.T) {
	const appname, command, testFlag, arg = "TESTING", "COMMAND", "-e", "-1"
	os.Args = []string{appname, command, testFlag, arg}
	s := Switch{}
	err := s.checkCommandArgs(1)
	if err != nil {
		t.Fatalf("Expected nil error, got %s", err)
	}
}

func TestCheckCommandArgs_WithHFlag(t *testing.T) {
	const appname, command, testFlag = "TESTING", "COMMAND", "-h"
	os.Args = []string{appname, command, testFlag}
	s := Switch{}
	err := s.checkCommandArgs(1)
	if err != nil {
		t.Fatalf("Expected nil error, got %s", err)
	}
}

func TestCheckCommandArgs_WithHelpFlag(t *testing.T) {
	const appname, command, testFlag = "TESTING", "COMMAND", "--help"
	os.Args = []string{appname, command, testFlag}
	s := Switch{}
	err := s.checkCommandArgs(1)
	if err != nil {
		t.Fatalf("Expected nil error, got %s", err)
	}
}

func TestCheckCommandArgs_With0Args(t *testing.T) {
	const appname, command, testFlag = "TESTING", "COMMAND", "-flag"
	os.Args = []string{appname, command, testFlag}
	s := Switch{}
	err := s.checkCommandArgs(2)
	if err == nil {
		t.Fatal("Expected and error message, got nil")
	}
}

func TestHelp(t *testing.T) {
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

func TestInfo(t *testing.T) {
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

func TestSaveFlags(t *testing.T) {
	const command = "COMMAND"
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

func TestFetchFlags(t *testing.T) {
	const command = "COMMAND"
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

func TestSave(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	const command, flagI, flagL, flagO, arg = "save", "-i", "-l", "-o", "ARG"
	os.Args = []string{appname, command, flagI, arg, flagL, arg, flagO, arg}

	m := new(mocks.TranscriptClient)

	s := NewSwitch(appname, version, m)

	m.On("Save", arg, arg, arg).Return(nil)

	err := s.save()(command)
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
}

func TestSave_FailCase(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	const command, flagI, flagL, flagO, arg = "save", "-i", "-l", "-o", "ARG"
	const errorMsg = "ERROR"
	os.Args = []string{appname, command, flagI, arg, flagL, arg, flagO, arg}

	m := new(mocks.TranscriptClient)

	s := NewSwitch(appname, version, m)

	m.On("Save", arg, arg, arg).Return(errors.New(errorMsg))

	err := s.save()(command)
	if err == nil {
		t.Fatalf("Expected %s as error message, got nil", errorMsg)
	}
}

func TestFetch(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	const command, flagI, flagL, arg = "fetch", "-i", "-l", "ARG"
	const success = "SUCCESS"
	os.Args = []string{appname, command, flagI, arg, flagL, arg}

	m := new(mocks.TranscriptClient)

	s := NewSwitch(appname, version, m)

	m.On("Fetch", arg, arg).Return(success, nil)

	err := s.fetch()(command)
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
}

func TestFetch_FailCase(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	const command, flagI, flagL, arg = "fetch", "-i", "-l", "ARG"
	const empty, errorMsg = "", "ERROR"
	os.Args = []string{appname, command, flagI, arg, flagL, arg}

	m := new(mocks.TranscriptClient)

	s := NewSwitch(appname, version, m)

	m.On("Fetch", arg, arg).Return(empty, errors.New(errorMsg))

	err := s.fetch()(command)
	if err == nil {
		t.Fatalf("Expected %s as error message, got nil", errorMsg)
	}
}

func TestList(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	const command, flagI, arg = "list", "-i", "ARG"
	var success = []string{"SUCCESS"}
	os.Args = []string{appname, command, flagI, arg}

	m := new(mocks.TranscriptClient)

	s := NewSwitch(appname, version, m)

	m.On("List", arg).Return(success, nil)

	err := s.list()(command)
	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
}

func TestList_FailCase(t *testing.T) {
	const appname, version = "TESTING", "0.0.0"
	const command, flagI, arg = "list", "-i", "ARG"
	var errorMsg = errors.New("ERROR")
	os.Args = []string{appname, command, flagI, arg}

	m := new(mocks.TranscriptClient)

	s := NewSwitch(appname, version, m)

	m.On("List", arg).Return([]string{}, errorMsg)

	err := s.list()(command)
	if err == nil {
		t.Fatalf("Expected %s as error message, got nil", errorMsg)
	}
}
