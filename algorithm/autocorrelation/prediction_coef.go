package autocorrelation

func PredictionCoefficient(autoCorrValues []float64, orderPrediction int32) []float64 {
	coeffs := make([][]float64, orderPrediction)

	averageSqr := make([]float64, orderPrediction)
	pseudoK := make([]float64, orderPrediction)

	for i := int32(0); i < orderPrediction; i++ {
		coeffs[i] = make([]float64, orderPrediction)

		if i != 0 {
			averageSqr[i] = (1 - (pseudoK[i-1] * pseudoK[i-1])) * averageSqr[i-1]
		} else {
			averageSqr[i] = autoCorrValues[i]
		}

		if averageSqr[i] == 0 {
			pseudoK[i] = 0.001
		} else {
			pseudoK[i] = (autoCorrValues[i+1] - sumPseudoK(i, coeffs, autoCorrValues)) / averageSqr[i]
		}
		for j := i; j > 0; j-- {
			if j == i {
				coeffs[i][j] = pseudoK[i]
			} else {
				coeffs[i][j] = coeffs[i-1][j] - pseudoK[i]*coeffs[i-1][i-j-1]
			}
		}
	}

	return coeffs[orderPrediction-1]
}

func sumPseudoK(index int32, coeffs [][]float64, autoCorrValues []float64) float64 {
	total := float64(0)

	for j := int32(0); j < index; j++ {
		total += coeffs[index-1][j] * autoCorrValues[index-j]
	}

	return total
}
