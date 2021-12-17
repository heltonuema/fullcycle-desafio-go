package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIfPrintFullCycleRocks(t *testing.T) {
	//Expected message to be printed
	expectedMsg := "FullCycle rocks\n"

	//Change Stdout
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	//Run our code
	main()

	//Rescue Stdout
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	//Check results
	if string(out) != expectedMsg {
		t.Errorf("Expected %s, got %s.", expectedMsg, out)
	}
}
