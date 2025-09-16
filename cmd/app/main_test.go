package main

/*
import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	t.Parallel()

	// Redirect stdout to a buffer
	old := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	// Call the main function
	main()

	// Read the output from the buffer
	writer.Close()

	out, _ := io.ReadAll(reader)
	os.Stdout = old

	expectedOutput := "Hello, world!"

	// Check if the expected output is contained within out
	if !strings.Contains(string(out), expectedOutput) {
		t.Errorf("Output does not contain the term [%s]:\n%s", expectedOutput, string(out))
	}
}

*/
