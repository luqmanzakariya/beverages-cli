package views

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView6(t *testing.T) {
	// Save the original os.Stdout
	originalStdout := os.Stdout

	// Create a pipe to capture output
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Redirect os.Stdout to the write end of the pipe
	os.Stdout = writePipe

	// Call the function
	View6()

	// Close the write pipe to flush output and restore os.Stdout
	writePipe.Close()
	os.Stdout = originalStdout

	// Read the output from the read pipe
	var output bytes.Buffer
	_, err = output.ReadFrom(readPipe)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Define the expected output
	expectedOutput := "\033[H\033[2J" +
		"=======================================\n" +
		"Good bye~\n" +
		"=======================================\n"

	// Compare the actual and expected output
	assert.Equal(t, expectedOutput, output.String())
}