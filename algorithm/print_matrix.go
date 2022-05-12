// Package algorithm
//
//	________print_matrix.go________
//
//	Print matrix.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package algorithm

import "fmt"

// PrintMatrix
//	Function for output of the matrix orderPrediction x orderPrediction
func PrintMatrix(matrix [][]float64, orderPrediction int32) {
	for i := int32(0); i < orderPrediction; i++ {
		fmt.Printf("\n")
		for j := int32(0); j < orderPrediction; j++ {
			fmt.Printf(" %.3f ", matrix[i][j])
		}
	}
	fmt.Println("\n")
}
