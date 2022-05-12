package covariance

import "CourseWork/algorithm"

func SumFuncMatrixElem(segment []byte, i int32, k int32, segment_size int32, orderPrediction int32) (float64, error) {
	total := float64(0)

	segment_float := algorithm.DataShift(segment)

	for m := orderPrediction; m < segment_size-orderPrediction-1; m++ {
		total += segment_float[m-i] * segment_float[m-k]
	}

	return total, nil
}
