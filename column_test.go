package float64matrix


import "testing"


func TestColumn(t *testing.T) {

	tests := []struct{
		slice []float64
		column_length int
		row_length    int
		subtests []struct{
			column_number int
			expected []float64
		}
	}{
		{
			slice: []float64{1,2,3,4,5,6,7,8,9},
			column_length: 3,
			row_length:    3,
			subtests: []struct{
				column_number int
				expected []float64
			}{
				{
					column_number: 0,
					expected: []float64{1,2,3},
				},
				{
					column_number: 1,
					expected: []float64{4,5,6},
				},
				{
					column_number: 2,
					expected: []float64{7,8,9},
				},
			},
		},
		{
			slice: []float64{1,2,3,4,5,6,7,8,9,10},
			column_length: 5,
			row_length:    2,
			subtests: []struct{
				column_number int
				expected []float64
			}{
				{
					column_number: 0,
					expected: []float64{1,2,3,4,5},
				},
				{
					column_number: 1,
					expected: []float64{6,7,8,9,10},
				},
			},
		},
		{
			slice: []float64{1,2,3,4,5,6,7,8,9,10},
			column_length: 2,
			row_length:    5,
			subtests: []struct{
				column_number int
				expected []float64
			}{
				{
					column_number: 0,
					expected: []float64{1,2},
				},
				{
					column_number: 1,
					expected: []float64{3,4},
				},
				{
					column_number: 2,
					expected: []float64{5,6},
				},
				{
					column_number: 3,
					expected: []float64{7,8},
				},
				{
					column_number: 4,
					expected: []float64{9,10},
				},
			},
		},
	}

	for test_number,test := range tests {

		matrix := Make(test.slice, test.column_length, test.row_length)

		for subtest_number,subtest := range test.subtests {
			actual := matrix.Column(subtest.column_number)


			len_expected := len(subtest.expected)
			len_actual   := len(actual)

			if len_expected != len_actual {
				t.Errorf("For test #%d subtest #%d expected length %d but actually got %d. For column #%d and matrix = %v.", test_number, subtest_number, len_expected, len_actual, subtest.column_number, matrix.slice)
			}


			for i,expected_value := range subtest.expected {

				if actual_value := actual[i] ; expected_value != actual_value {
					t.Errorf("For test #%d subtest #%d for index %d expected value %f but actually got %f. For column #%d and matrix = %v.", test_number, subtest_number, i, expected_value, actual_value, subtest.column_number, matrix.slice)
				}

			}
		}
	}
}
