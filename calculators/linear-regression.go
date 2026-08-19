package calculators

import "errors"

func LinearRegressionLine(data []int) (float64, float64, error) {
	var a, b, n, sumX, sumY, sumXY, sumXX float64
	// var a, b, n float64
	n = float64(len(data))

	if n <= 1 {
		return 0, 0, errors.New("Not enough values to calculate Linear Regression Line")
	}

	sumX, sumY, sumXY, sumXX, _ = sums(data)
	// for i, num := range data {
	// 	sumX += float64(i)
	// 	sumY += float64(num)
	// 	sumXY += float64(i) * float64(num)
	// 	sumXX += float64(i) * float64(i)
	// }

	a = ((n * sumXY) - (sumX * sumY)) / ((n * sumXX) - (sumX * sumX))
	b = (sumY - (a * sumX)) / n
	return a, b, nil
}
