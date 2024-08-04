package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(lines []string) ([]float64, error) {
	prices := make([]float64, len(lines))
	for stringIndex, stringVal := range lines {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}
		prices[stringIndex] = floatPrice
	}
	return prices, nil
}
