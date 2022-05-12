// Package autocorrelation
//
//	________autocorrelation.go________
//
//	Autocorrelation method.
//	Method of James Durbin.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package autocorrelation

import (
	"CourseWork/algorithm"
	"CourseWork/decoder/waveform"
)

func Autocorrelation(desc *waveform.WAVFormat, orderPrediction int32) (float64, error) {

	energy, err := CalculateEnergyForWAV(desc, orderPrediction)

	if err != nil {
		return -1, err
	}

	return average(energy), nil
}

func average(array []float64) float64 {
	total := float64(0)

	for _, elem := range array {
		if elem > 1000 {
			continue
		}
		total += elem
	}

	return total / float64(len(array)*algorithm.ENERGY_DEVIDER)
}
