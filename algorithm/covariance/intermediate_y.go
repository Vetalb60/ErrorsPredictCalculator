package covariance

func intermediateY(downTriangular [][]float64, sumFuncZero []float64, orderPrediction int32) []float64 {
	intermediate_y := make([]float64, orderPrediction)

	for i := int32(0); i < orderPrediction; i++ {
		intermediate_y[i] = float64(sumFuncZero[i]) - sumIntermediateY(i, downTriangular, intermediate_y)
	}

	return intermediate_y
}

func sumIntermediateY(i int32, downTriangular [][]float64, intermediate_y []float64) (sum float64) {
	for k := int32(0); k < i; k++ {
		sum += float64(downTriangular[i][k]) * intermediate_y[k]
	}

	return sum
}
