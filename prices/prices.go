package prices

import (
	"fmt"
	"price_calculator/conversion"
	"price_calculator/iomanager"
)

// TaxIncludedPriceJob represents a job to calculate tax included prices
type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager // IOManager interface for reading and writing data
	TaxRate           float64             // Tax rate to apply
	InputPrices       []float64           // Input prices to process
	TaxIncludedPrices map[string]string   // Resulting tax included prices
}

// LoadData loads input prices from the IOManager
func (job *TaxIncludedPriceJob) LoadData() error {
	// Read lines from IOManager
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	// Convert strings to floats
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

// NewTaxIncludedPriceJob creates a new instance of TaxIncludedPriceJob
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30}, // Default input prices for demonstration
		TaxRate:     taxRate,
	}
}

// Process calculates tax included prices
func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	// Load input data
	err := job.LoadData()

	if err != nil {
		fmt.Println(err)
		return
	}

	result := make(map[string]string)
	// Calculate tax included prices for each input price
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	// Write result to IOManager
	job.IOManager.WriteResult(job.TaxIncludedPrices)

	doneChan <- true
}
