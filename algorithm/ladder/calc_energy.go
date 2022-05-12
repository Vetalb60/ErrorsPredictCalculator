package ladder

import (
	"CourseWork/algorithm"
	"CourseWork/decoder/waveform"
)

type LadderErrors struct {
	forward_error []float64
	reverse_error []float64
}

type EnergyOfSegment struct {
	energyOnPredict []float64
}

func CalculateEnergyForWAV(desc *waveform.WAVFormat, orderPrediction int32) ([]EnergyOfSegment, error) {
	count := int32(0)
	segmentSize := int32(desc.Meta_.ByteRate_ / algorithm.SEGMENT_FREQ)
	lad_err := make([]LadderErrors, orderPrediction)
	ref_coeffs := make([]float64, orderPrediction)
	energyForSegment := make([]EnergyOfSegment, orderPrediction)

	for i := int32(0); i < orderPrediction; i++ {
		energyForSegment[i].energyOnPredict = make([]float64, int32(desc.Meta_.FileSize_)/segmentSize)
	}

	for index := segmentSize; index < int32(desc.Meta_.FileSize_)-segmentSize-orderPrediction; index += segmentSize {
		segment := desc.Data_[index-1 : index+segmentSize+1]
		lad_err[0].forward_error = ladderDataShift(segment)
		lad_err[0].reverse_error = ladderDataShift(segment)
		ref_coeffs[0] = calcReflectionCoeff(lad_err[0].forward_error, lad_err[0].reverse_error)
		energyForSegment[0].energyOnPredict[count] =
			calcErrorEnergy(lad_err[0].forward_error, lad_err[0].reverse_error)

		for i := int32(1); i < orderPrediction; i++ {
			lad_err[i].forward_error = calcForwardError(i, lad_err, ref_coeffs)
			lad_err[i].reverse_error = calcReverseError(i, lad_err, ref_coeffs)

			ref_coeffs[i] = calcReflectionCoeff(lad_err[i].forward_error, lad_err[i-1].reverse_error)

			energyForSegment[i].energyOnPredict[count] =
				calcErrorEnergy(lad_err[i].forward_error, lad_err[i].reverse_error)
		}
		count++
	}

	return energyForSegment, nil
}

func ladderDataShift(segment []byte) []float64 {
	new_segment := make([]float64, len(segment))

	for index, elem := range segment {
		new_segment[index] = float64(int32(elem) - 0x80)
	}

	return new_segment
}
