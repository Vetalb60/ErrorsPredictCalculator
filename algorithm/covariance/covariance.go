// Package covariance
//
//	________covariance.go________
//
//	Covariance method.
//	Method of Cholesky.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package covariance

import (
	"CourseWork/decoder/waveform"
)

func Covariance(desc *waveform.WAVFormat, orderPrediction int32) (float64, error) {

	energy, err := CalculateEnergyForWAV(desc, orderPrediction)

	if err != nil {
		return -1, err
	}

	return average(energy), nil
}

func average(array []float64) float64 {
	total := float64(0)
	count := int(0)

	for _, elem := range array {
		if elem != 0 {
			total += elem
			count++
		}
	}

	return total / (float64(count) * 1000)
}
