package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)



func Test_printSomething(t *testing.T) {

	stdOut := os.Stdout

	// Create a pipe to read from stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("Hello", &wg)
	
	wg.Wait()

	// Close the write end of the pipe so that the reader knows
	// when to stop reading
	w.Close()

	// Read the output

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello") {
		t.Errorf("printSomething() = %v, want %v", output, "Hello")
	}

}