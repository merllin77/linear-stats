package calculators

import (
	"errors"
	"math"
)

func PearsonCorrelationCoefficient(data []int) (float64, error) {
	var n, r, sumX, sumY, sumXY, sumXX, sumYY float64
	n = float64(len(data))

	if n <= 1 {
		return 0, errors.New("Not enough values to calculate Pearson Correlation Coefficient")
	}

	sumX, sumY, sumXY, sumXX, sumYY = sums(data)
	zeroResult := (n * sumYY) - (sumY * sumY)

	// Check if all values of data file are equal
	if zeroResult == 0 {
		return 0, errors.New("All values in data file are equal. Cannot calculate Pearson Correlation Coefficient")
	}

	r = ((n * sumXY) - (sumX * sumY)) / math.Sqrt(((n*sumXX)-(sumX*sumX))*((n*sumYY)-(sumY*sumY)))
	return r, nil
}
