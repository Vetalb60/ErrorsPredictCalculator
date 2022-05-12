package ladder

func calcErrorEnergy(forward_error []float64, reverse_error []float64) float64 {
	total := float64(0)

	for i := int(0); i < len(forward_error); i++ {
		total += (forward_error[i] * forward_error[i]) + (reverse_error[i] * reverse_error[i])
	}

	return total
}
