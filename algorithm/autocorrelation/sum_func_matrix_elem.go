package autocorrelation

import "CourseWork/algorithm"

func SumFuncMatrixElem(segment []byte, i int32, k int32, segment_size int32) (float64, error) {
	total := float64(0)

	segment_float := algorithm.DataShift(segment)

	for m := int32(0); m < segment_size-k-1; m++ {
		total += segment_float[m+i] * segment_float[m+k]
	}

	return total, nil
}
