package float64matrix


// I returns the element of the matrix at i,j = column_index,row_index.
func (matrix T) I(column_index int, row_index int) float64 {

	i := row_index * matrix.column_length + column_index

	return matrix.slice[i]
}
