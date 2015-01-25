package float64matrix


// Column returns the "column_number" column as a []float64 slice.
// So, for example, if "column_number" is 2, then this method returns
// the 2nd column as a []float64 slice.
func (matrix T) Column(column_number int) []float64 {

	column_length := matrix.column_length

	begin := column_length * column_number
	end   := begin + matrix.column_length

	return matrix.slice[begin:end]
}
