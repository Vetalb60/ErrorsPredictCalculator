package covariance

func predictionCoefficient(
	upperTriangular [][]float64,
	intermidiateY []float64,
	diagonal_elem [][]float64,
	orderPrediction int32) (
	coeffs []float64) {
	coeffs = make([]float64, orderPrediction)

	for i := orderPrediction - 1; i >= 0; i-- {
		if diagonal_elem[i][i] != 0 {
			coeffs[i] = intermidiateY[i]/diagonal_elem[i][i] - coeffSum(i, upperTriangular, coeffs, orderPrediction)
		} else {
			coeffs[i] = 0.0001
		}
	}

	return coeffs
}

func coeffSum(index int32, upperTriangular [][]float64, coeffs []float64, orderPrediction int32) float64 {
	total := float64(0)
	for k := int32(index + 1); k < orderPrediction; k++ {
		total += upperTriangular[index][k] * coeffs[k]
	}

	return total
}
