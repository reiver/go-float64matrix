package float64maxtrix_test

import "fmt"
import "github.com/reiver/go-float64maxtrix"


// This example shows basic usage of float64maxtrix.
func Example_basic() {

	slice := []float64{1,2,3,4,5,6,7,8}

	// Note that we creating a 5×2 matrix here.
	// I.e, a matrix with column length 5 and
	// row length 2.
	//
	// The matrix looks like this:
	//
	//	| 1 3 5 7 | 
	//	| 2 4 6 8 |
	//
	// Also note that "slice" and "matrix" are
	// sharing the same memory. So if you change
	// one you change the other.
	//
	matrix := float64maxtrix.Make(slice, 2, 4)


	// Output.
	fmt.Printf("slice = %v\n", slice)
	fmt.Println()
	fmt.Printf("multi-dimensional slice = | %.0f %.0f %.0f %.0f |\n", matrix.I(0,0), matrix.I(0,1), matrix.I(0,2), matrix.I(0,3))
	fmt.Printf("                          | %.0f %.0f %.0f %.0f |\n", matrix.I(1,0), matrix.I(1,1), matrix.I(1,2), matrix.I(1,3))

	// Output:
	// slice = [1 2 3 4 5 6 7 8]
	//
	// multi-dimensional slice = | 1 3 5 7 |
	//                           | 2 4 6 8 |
}

