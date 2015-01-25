package float64matrix


// Set sets the element of the matrix at i,j = column_index,row_index.
func (matrix T) Set(value float64, column_index int, row_index int) {

	i := row_index * matrix.column_length + column_index

	matrix.slice[i] = value
}
