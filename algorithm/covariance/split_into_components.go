package covariance

func splitIntoComponents(
	sumFuncMatrix [][]float64,
	orderPrediction int32) (
	downTriangular [][]float64,
	upperTriangular [][]float64,
	diagonal [][]float64) {

	downTriangular = make([][]float64, orderPrediction)
	diagonalElems := make([]float64, orderPrediction)

	for i := int32(0); i < orderPrediction; i++ {
		downTriangular[i] = make([]float64, orderPrediction)
	}

	diagonalElems[0] = sumFuncMatrix[0][0]
	downTriangular[0][0] = 1
	for i := int32(1); i < orderPrediction; i++ {
		for j := int32(0); j < orderPrediction; j++ {
			if i != j {
				if diagonalElems[j] == 0 {
					downTriangular[i][j] = (sumFuncMatrix[i][j] - downTglSum(i, j, downTriangular, diagonalElems)) / 0.0001
				} else {
					downTriangular[i][j] = (sumFuncMatrix[i][j] - downTglSum(i, j, downTriangular, diagonalElems)) / diagonalElems[j]
				}
			} else {
				downTriangular[i][i] = 1
				break
			}
		}
		diagonalElems[i] = sumFuncMatrix[i][i] - diagonalSum(i, downTriangular, diagonalElems)
	}

	diagonal = fillDiagonalMatrix(diagonalElems, orderPrediction)

	upperTriangular = transposeMatrix(downTriangular, orderPrediction)

	return downTriangular, upperTriangular, diagonal
}

func fillDiagonalMatrix(diagonalElems []float64, orderPrediction int32) (diagonal [][]float64) {
	diagonal = make([][]float64, orderPrediction)

	for i := int32(0); i < orderPrediction; i++ {
		diagonal[i] = make([]float64, orderPrediction)
		for j := int32(0); j < orderPrediction; j++ {
			if i == j {
				diagonal[i][j] = diagonalElems[i]
			} else {
				diagonal[i][j] = 0
			}
		}
	}

	return diagonal
}

func diagonalSum(index int32, downTriangular [][]float64, diagonalElems []float64) (sum float64) {
	for k := int32(0); k < index; k++ {
		sum += downTriangular[index][k] * downTriangular[index][k] * diagonalElems[k]
	}
	return sum
}

func downTglSum(i int32, j int32, downTriangular [][]float64, diagonalElems []float64) (sum float64) {
	for k := int(j - 1); k >= 0; k-- {
		sum += downTriangular[i][k] * downTriangular[j][k] * diagonalElems[k]
	}
	return sum
}

func transposeMatrix(matrix [][]float64, orderPrediction int32) (trans [][]float64) {
	trans = make([][]float64, orderPrediction)
	for i := int32(0); i < orderPrediction; i++ {
		trans[i] = make([]float64, orderPrediction)
		for j := int32(0); j < orderPrediction; j++ {
			trans[i][j] = matrix[j][i]
		}
	}

	return trans
}
