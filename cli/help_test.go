package cli

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_PrintMsg(t *testing.T) {
	stderr := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("PIPE error: %s", err)
	}
	os.Stderr = w

	const testing = "TESTING"
	os.Args[0] = testing
	os.Args[1] = testing
	printMsg(testing)

	w.Close()

	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("READ error: %s", err)
	}
	os.Stdout = stderr

	count := strings.Count(string(out), testing)
	if count == 0 {
		t.Errorf("Expected 3 occurencies, got %d, of the value: %s", count, testing)
	}
}
