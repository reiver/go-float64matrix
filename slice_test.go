package float64matrix


import "testing"


func TestSlice(t *testing.T) {

	tests := []struct{
		column_length int
		row_length    int
		expected []float64
		actions []struct{
			i int
			j int
			value float64
		}
	}{
		{
			column_length: 3,
			row_length:    3,
			expected: []float64{1,2,3,4,5,6,7,8,9},
			actions: []struct{
				i int  
				j int
				value float64
			}{
				{
					i: 0, j: 0,
					value: 1,
				},
				{
					i: 1, j: 0,
					value: 2,
				},
				{
					i: 2, j: 0,
					value: 3,
				},
				{
					i: 0, j: 1,
					value: 4,
				},
				{
					i: 1, j: 1,
					value: 5,
				},
				{
					i: 2, j: 1,
					value: 6,
				},
				{
					i: 0, j: 2,
					value: 7,
				},
				{
					i: 1, j: 2,
					value: 8,
				},
				{
					i: 2, j: 2,
					value: 9,
				},
			},
		},
		{
			column_length: 5,
			row_length:    2,
			expected: []float64{1,2,3,4,5,6,7,8,9,10},
			actions: []struct{
				i int  
				j int
				value float64
			}{
				{
					i: 0, j: 0,
					value: 1,
				},
				{
					i: 1, j: 0,
					value: 2,
				},
				{
					i: 2, j: 0,
					value: 3,
				},
				{
					i: 3, j: 0,
					value: 4,
				},
				{
					i: 4, j: 0,
					value: 5,
				},
				{
					i: 0, j: 1,
					value: 6,
				},
				{
					i: 1, j: 1,
					value: 7,
				},
				{
					i: 2, j: 1,
					value: 8,
				},
				{
					i: 3, j: 1,
					value: 9,
				},
				{
					i: 4, j: 1,
					value: 10,
				},
			},
		},
		{
			column_length: 2,
			row_length:    5,
			expected: []float64{1,2,3,4,5,6,7,8,9,10},
			actions: []struct{
				i int  
				j int
				value float64
			}{
				{
					i: 0, j: 0,
					value: 1,
				},
				{
					i: 1, j: 0,
					value: 2,
				},
				{
					i: 0, j: 1,
					value: 3,
				},
				{
					i: 1, j: 1,
					value: 4,
				},
				{
					i: 0, j: 2,
					value: 5,
				},
				{
					i: 1, j: 2,
					value: 6,
				},
				{
					i: 0, j: 3,
					value: 7,
				},
				{
					i: 1, j: 3,
					value: 8,
				},
				{
					i: 0, j: 4,
					value: 9,
				},
				{
					i: 1, j: 4,
					value: 10,
				},
			},
		},
	}


	for test_number, test := range tests {

		actual := make([]float64, len(test.expected))

		actual_matrix := Make(actual, test.column_length, test.row_length)


		for _,action := range test.actions {
			actual_matrix.Set(action.value, action.i, action.j)
		}


		for ii,actual_value := range actual {

			if expected_value := test.expected[ii] ; expected_value != actual_value {
				t.Errorf("For test #%d for array index #%d expected value %f but instead got %f. For expected slice %v and actual slice %v.", test_number, ii, expected_value, actual_value, test.expected, actual)
			}
		}
	}
}
