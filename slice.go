package float64matrix


// Slice returns the data of the matrix in []float64 slice form.
// Note that the data is returned in column-major order.
func (matrix T) Slice() []float64 {

	return matrix.slice
}
