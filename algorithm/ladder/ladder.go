// Package ladder
//
//	________ladder.go________
//
//	Ladder method.
//	Method of Burg J.R.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package ladder

import (
	"CourseWork/algorithm"
	"CourseWork/decoder/waveform"
)

func Ladder(desc *waveform.WAVFormat, orderPrediction int32) ([]float64, error) {

	energyAllSegments, err := CalculateEnergyForWAV(desc, orderPrediction)
	averageEnergy := make([]float64, orderPrediction)
	if err != nil {
		return nil, err
	}

	for i := int32(0); i < orderPrediction; i++ {
		averageEnergy[i] = average(energyAllSegments[i].energyOnPredict)
	}

	return averageEnergy, nil
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
