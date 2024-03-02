package main

import (
	"fmt"
	"price_calculator/filemanager"
	"price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100)) // Adjust file paths as needed
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)                            // Pass FileManager to NewTaxIncludedPriceJob
		go priceJob.Process(doneChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
