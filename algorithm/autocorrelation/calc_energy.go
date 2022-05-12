package autocorrelation

import (
	"CourseWork/algorithm"
	"CourseWork/decoder/waveform"
)

type autoCorrMatr struct {
	AutoCorrValues []float64
}

func CalculateEnergyForWAV(desc *waveform.WAVFormat, orderPrediction int32) ([]float64, error) {
	var err error
	count := int(0)
	segmentSize := int32(desc.Meta_.ByteRate_ / algorithm.SEGMENT_FREQ)
	energyForSegment := make([]float64, int32(desc.Meta_.FileSize_)/segmentSize)
	autoCorrMatrobj := make([]autoCorrMatr, int32(desc.Meta_.FileSize_)/segmentSize)

	for index := segmentSize; index < int32(desc.Meta_.FileSize_)-segmentSize-orderPrediction; index += segmentSize {
		autoCorrMatrobj[count].AutoCorrValues = make([]float64, orderPrediction+1)
		segment := desc.Data_[index : index+segmentSize]
		for i := int32(0); i < orderPrediction+1; i++ {
			autoCorrMatrobj[count].AutoCorrValues[i], err = SumFuncMatrixElem(segment, 0, i, segmentSize)
			if err != nil {
				return nil, err
			}

		}
		coeffs := PredictionCoefficient(autoCorrMatrobj[count].AutoCorrValues, orderPrediction)
		energyForSegment[count] = algorithm.PredictionErrorEnergy(segment, coeffs, orderPrediction, segmentSize)
		count++
	}

	return energyForSegment, nil
}
