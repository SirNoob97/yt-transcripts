package main

import (
	"flag"
	"os"
	"testing"
)

func TestShortFlagsFalse(t *testing.T) {
	os.Args = []string{"appname"}
	flag.Parse()

	if *versionShort && *helpShort {
		t.Fatalf("Expected false, false, got %v, %v", *versionShort, *helpShort)
	}
}

func TestLongFlagsFalse(t *testing.T) {
	os.Args = []string{"appname"}
	flag.Parse()

	if *versionLong && *helpLong {
		t.Fatalf("Expected false, false, got %v, %v", *versionLong, *helpLong)
	}
}

func TestShortFlagsTrue(t *testing.T) {
	os.Args = []string{"appname", "-v", "-h"}
	flag.Parse()

	if !*versionShort && !*helpShort {
		t.Fatalf("Expected true, true, got %v, %v", *versionShort, *helpShort)
	}
}

func TestLongFlagsTrue(t *testing.T) {
	os.Args = []string{"appname", "-version", "-help"}
	flag.Parse()

	if !*versionLong && !*helpLong {
		t.Fatalf("Expected true, true, got %v, %v", *versionLong, *helpLong)
	}
}
