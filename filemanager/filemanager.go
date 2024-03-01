package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

// FileManager represents a file manager
type FileManager struct {
	InputFilePath  string // Path to input file
	OutputFilePath string // Path to output file
}

// ReadLines reads lines from a file
func (fm FileManager) ReadLines() ([]string, error) {
	// Open input file
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer file.Close()

	// Create a scanner to read lines
	scanner := bufio.NewScanner(file)

	var lines []string

	// Read lines and append to slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		return nil, errors.New("failed to read line in file")
	}

	return lines, nil
}

// WriteResult writes data to a file
func (fm FileManager) WriteResult(data interface{}) error {
	// Create output file
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	// Encode data and write to file
	if err := encoder.Encode(data); err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

// New creates a new instance of FileManager
func New(inputPath string, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
