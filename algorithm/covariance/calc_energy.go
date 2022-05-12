package covariance

import (
	"CourseWork/algorithm"
	"CourseWork/decoder/waveform"
)

type sumFuncMatr struct {
	sumFunc     [][]float64
	sumFuncZero []float64
}

func CalculateEnergyForWAV(desc *waveform.WAVFormat, orderPrediction int32) ([]float64, error) {
	var err error
	count := int(0)
	segmentSize := int32(desc.Meta_.ByteRate_ / algorithm.SEGMENT_FREQ)
	energyForSegment := make([]float64, int32(desc.Meta_.FileSize_)/segmentSize)

	sumFuncMatrobj := make([]sumFuncMatr, int32(desc.Meta_.FileSize_)/segmentSize)
	for index := segmentSize; index < int32(desc.Meta_.FileSize_)-segmentSize-orderPrediction; index += segmentSize {
		sumFuncMatrobj[count].sumFunc = make([][]float64, orderPrediction)
		sumFuncMatrobj[count].sumFuncZero = make([]float64, orderPrediction)
		segment := desc.Data_[index-orderPrediction : index+segmentSize+orderPrediction]
		for i := int32(0); i < orderPrediction; i++ {
			sumFuncMatrobj[count].sumFunc[i] = make([]float64, orderPrediction)

			sumFuncMatrobj[count].sumFuncZero[i], err = SumFuncMatrixElem(segment, i, 0, segmentSize, orderPrediction)

			if err != nil {
				return nil, err
			}

			for k := int32(0); k < orderPrediction; k++ {
				sumFuncMatrobj[count].sumFunc[i][k], err = SumFuncMatrixElem(segment, i, k, segmentSize, orderPrediction)

				if err != nil {
					return nil, err
				}
			}
		}

		down, up, diag := splitIntoComponents(sumFuncMatrobj[count].sumFunc, orderPrediction)
		int_y := intermediateY(down, sumFuncMatrobj[count].sumFuncZero, orderPrediction)
		coeffs := predictionCoefficient(up, int_y, diag, orderPrediction)

		energyForSegment[count] = algorithm.PredictionErrorEnergy(segment, coeffs, orderPrediction, segmentSize)

		count++
	}

	return energyForSegment, nil

}
