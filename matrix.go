package float64maxtrix


// A T is the matrix data type.
type T struct {
	slice []float64
	column_length int
	row_length    int
}


// Make conceptually turns a []float64 slice into a matrix.
// Note that this uses the same backing array as the slice.
func Make(slice []float64, column_length int, row_length int) T {
	return T{
		slice: slice,
		column_length: column_length,
		row_length: row_length,
	}
}
