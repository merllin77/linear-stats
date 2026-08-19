package calculators

func sums(data []int) (float64, float64, float64, float64, float64) {
	var sumX, sumY, sumXY, sumXX, sumYY float64

	for i, num := range data {
		sumX += float64(i)
		sumY += float64(num)
		sumXY += float64(i) * float64(num)
		sumXX += float64(i) * float64(i)
		sumYY += float64(num) * float64(num)
	}
	return sumX, sumY, sumXY, sumXX, sumYY
}
