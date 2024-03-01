package conversion

import (
	"errors"
	"strconv"
)

// StringsToFloats converts string slice to float slice
func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	// Iterate over strings and convert to floats
	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float.")
		}
		floats = append(floats, floatVal)
	}

	return floats, nil
}
