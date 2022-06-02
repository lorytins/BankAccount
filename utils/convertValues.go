package utils

func ConvertToCents(value float64) int{
	return int(value * 100)
}

func ConvertToReal(value int) float64{
	return float64(value) / float64(100)
}

func CalcPercent(value int, percent float64) float64{
	percentageValue := (float64(value) * percent) / 100
	return percentageValue
}