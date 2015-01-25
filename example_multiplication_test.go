package float64matrix_test

import "fmt"
import "github.com/reiver/go-float64matrix"


// This example shows usage of float64matrix for matrix vector multiplication.
// Note that there are more efficient ways of doing matrix vector multiplication;
// this is show this way for the sake an example.
func Example_multiplication() {

	// Create indexer for a 2x3 matrix.
	// I.e., a matrix with columns of length 2 and rows of length 3.
	//
	// This is the matrix:
	//
	//	| 1 3 5 |
	//	| 2 4 6 |
	//
	slice  := []float64{1, 2, 3, 4, 5, 6}
	matrix := float64matrix.Make(slice, 2, 5)


	// Create a vector.
	//
	// This is the vector:
	//
	//	| 7 |
	//	| 8 |
	//	| 9 |
	//
	vector := []float64{7, 8, 9}


	// Create a vector to store the results.
	//
	// Since the matrix is of size 2×3 and
	// the vector is of length 3, the resulting
	// vector will be of length 2.
	//
	// In other words:
	//
	//	| 1 3 5 |   | 7 |   | ? |
	//	| 2 4 6 | × | 8 | = | ? |
	//                  | 9 |
	//
	result_vector := make([]float64, 2)


	// Do the matrix multiplication.
	result_vector[0] = matrix.I(0,0) * vector[0] +
	                   matrix.I(0,1) * vector[1] +
	                   matrix.I(0,2) * vector[2]

	result_vector[1] = matrix.I(1,0) * vector[0] +
	                   matrix.I(1,1) * vector[1] +
	                   matrix.I(1,2) * vector[2]




	// Output.
	fmt.Printf("| %.0f %.0f %.0f |   | %.0f |   | %3.0f |\n", matrix.I(0,0),
	                                                          matrix.I(0,1),
	                                                          matrix.I(0,2),
	                                                          vector[0],
	                                                          result_vector[0])
	fmt.Printf("| %.0f %.0f %.0f | × | %.0f | = | %3.0f |\n", matrix.I(1,0),
	                                                          matrix.I(1,1),
	                                                          matrix.I(1,2),
	                                                          vector[1],
	                                                          result_vector[1])
	fmt.Printf("            | %.0f |\n",                      vector[2])
	fmt.Println()


	// Output:
	// | 1 3 5 |   | 7 |   |  76 |
	// | 2 4 6 | × | 8 | = | 100 |
	//             | 9 |
}
