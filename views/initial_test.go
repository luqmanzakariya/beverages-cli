package views

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialMenu(t *testing.T) {
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
	InitialMenu()

	// Close the write pipe and restore os.Stdout
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
		"Welcome to Team 2 Wine & Beverages\n" +
		"=======================================\n" +
		"Please choose:\n" +
		"1. Admin Menu\n" +
		"2. Order\n" +
		"3. Stock\n" +
		"4. Categories\n" +
		"5. Generate Reports\n" +
		"6. Exit\n\n" +
		"Answer: "

	// Compare the actual and expected output
	assert.Equal(t, expectedOutput, output.String())
}

func TestReturnString(t *testing.T) {
	result := ReturnString("Test return the same string")

	if result != "Test return the same string" {
		// error
		t.Error("Result must be 'Test'")
	}
}
