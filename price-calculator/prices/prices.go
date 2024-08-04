package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate             float64           `json:"tax_rate"`
	InputPrices         []float64         `json:"input_prices"`
	TaxIncludedPrices   map[string]string `json:"tax_included_prices"`
	iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.loadData()
	if err != nil {
		return err
	}
	result := map[string]string{}
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	job.WriteResult(job)
	return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) loadData() error {
	lines, err := job.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil
}
