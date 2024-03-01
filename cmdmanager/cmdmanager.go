package cmdmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

// CMDManager represents a command line interface manager
type CMDManager struct{}

// ReadLines reads lines from command line input
func (cmdm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER. Enter '0' to finish.")

	var prices []string

	// Read prices from command line until "0" or end-of-file is entered
	for {
		var price string
		fmt.Println("Price:")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

// WriteResult writes data to a JSON file
func (cmdm CMDManager) WriteResult(data interface{}) error {
	// Open output file
	file, err := os.Create("result.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode data to JSON and write to file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// New creates a new instance of CMDManager
func New() CMDManager {
	return CMDManager{}
}
