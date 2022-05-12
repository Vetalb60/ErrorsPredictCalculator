// Package algorithm
//
//	________predication_error_energy.go________
//
//	The basic formula for finding the prediction error energy
//	for covariance and autocorrelation methods.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package algorithm

func PredictionErrorEnergy(segment []byte,
	predictionCoeffs []float64,
	orderPrediction int32,
	segmentSize int32) float64 {

	energy_error := float64(0)
	segment_float := DataShift(segment)

	for i := orderPrediction; i < segmentSize-orderPrediction; i++ {
		buf := float64(segment_float[i]) - energySum(i, segment_float, predictionCoeffs, orderPrediction)
		energy_error += buf * buf
	}

	return energy_error
}

func energySum(index int32, segment []float64, predictionCoeffs []float64, orderPrediction int32) float64 {
	total := float64(0)

	for k := int32(0); k < orderPrediction; k++ {
		total += predictionCoeffs[k] * segment[index-k]
	}

	return total
}

func DataShift(segment []byte) []float64 {
	new_segment := make([]float64, len(segment))

	for index, elem := range segment {
		new_segment[index] = float64((elem - 0x80) / 0x80)
	}

	return new_segment
}
