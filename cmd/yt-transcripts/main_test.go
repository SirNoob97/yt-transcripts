package main

import (
	"os"
	"testing"
)

func TestFlags(t *testing.T) {
	os.Args = []string{"appname", "-v", "-h"}
	v, h := flags()

	if !*v && !*h {
		t.Fatalf("Expected true, true, got %v, %v", *v, *h)
	}
}

func TestFlagsFalseCase(t *testing.T) {
	os.Args = []string{"appname"}
	v, h := flags()

	if *v && *h {
		t.Fatalf("Expected false, false, got %v, %v", *v, *h)
	}
}
