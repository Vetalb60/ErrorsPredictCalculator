package ladder

func calcReflectionCoeff(forward_errors []float64, reverse_forward []float64) float64 {
	coeff := 2 * reflectionCoeffSum(forward_errors, reverse_forward) / (reflectionCoeffSqr(forward_errors))

	return coeff
}

func reflectionCoeffSum(forward_errors []float64, reverse_forward []float64) float64 {
	total := float64(0)

	for i := int(1); i < len(forward_errors)-1; i++ {
		total += forward_errors[i] * reverse_forward[i]
	}

	return total
}

func reflectionCoeffSqr(forward_errors []float64) float64 {
	total := float64(0)

	for i := int(1); i < len(forward_errors)-1; i++ {
		total += forward_errors[i] * forward_errors[i]
	}

	if total == 0 {
		return 0.0001
	}

	return total
}
