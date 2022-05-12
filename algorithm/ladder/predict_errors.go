package ladder

func calcForwardError(current_order int32, ladder_error []LadderErrors, reflection_coeffs []float64) []float64 {
	array := make([]float64, len(ladder_error[current_order-1].forward_error))

	for i := int(0); i < len(ladder_error[current_order-1].forward_error)-1; i++ {
		array[i] = ladder_error[current_order-1].forward_error[i] -
			reflection_coeffs[current_order]*ladder_error[current_order-1].reverse_error[i+1]
	}

	return array
}

func calcReverseError(current_order int32, ladder_error []LadderErrors, reflection_coeffs []float64) []float64 {
	array := make([]float64, len(ladder_error[current_order-1].forward_error))

	for i := int(0); i < len(ladder_error[current_order-1].forward_error)-1; i++ {
		array[i] = ladder_error[current_order-1].reverse_error[i] -
			reflection_coeffs[current_order]*ladder_error[current_order-1].forward_error[i+1]
	}

	return array
}
