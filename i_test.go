package float64maxtrix


import "testing"


func TestI(t *testing.T) {

	tests := []struct{
		slice []float64
		column_length int
		row_length    int
		subtests []struct{
			i int
			j int
			expected float64
		}
	}{
		{
			slice: []float64{1,2,3,4,5,6,7,8,9},
			column_length: 3,
			row_length:    3,
			subtests: []struct{
				i int
				j int
				expected float64
			}{
				{
					i: 0 , j: 0,
					expected: 1,
				},
				{
					i: 1 , j: 0,
					expected: 2,
				},
				{
					i: 2 , j: 0,
					expected: 3,
				},
				{
					i: 0 , j: 1,
					expected: 4,
				},
				{
					i: 1 , j: 1,
					expected: 5,
				},
				{
					i: 2 , j: 1,
					expected: 6,
				},
				{
					i: 0 , j: 2,
					expected: 7,
				},
				{
					i: 1 , j: 2,
					expected: 8,
				},
				{
					i: 2 , j: 2,
					expected: 9,
				},
			},
		},
		{
			slice: []float64{1,2,3,4,5,6,7,8,9,10},
			column_length: 5,
			row_length:    2,
			subtests: []struct{
				i int
				j int
				expected float64
			}{
				{
					i: 0 , j: 0,
					expected: 1,
				},
				{
					i: 1 , j: 0,
					expected: 2,
				},
				{
					i: 2 , j: 0,
					expected: 3,
				},
				{
					i: 3 , j: 0,
					expected: 4,
				},
				{
					i: 4 , j: 0,
					expected: 5,
				},
				{
					i: 0 , j: 1,
					expected: 6,
				},
				{
					i: 1 , j: 1,
					expected: 7,
				},
				{
					i: 2 , j: 1,
					expected: 8,
				},
				{
					i: 3 , j: 1,
					expected: 9,
				},
				{
					i: 4 , j: 1,
					expected: 10,
				},
			},
		},
		{
			slice: []float64{1,2,3,4,5,6,7,8,9,10},
			column_length: 2,
			row_length:    5,
			subtests: []struct{
				i int
				j int
				expected float64
			}{
				{
					i: 0 , j: 0,
					expected: 1,
				},
				{
					i: 1 , j: 0,
					expected: 2,
				},
				{
					i: 0 , j: 1,
					expected: 3,
				},
				{
					i: 1 , j: 1,
					expected: 4,
				},
				{
					i: 0 , j: 2,
					expected: 5,
				},
				{
					i: 1 , j: 2,
					expected: 6,
				},
				{
					i: 0 , j: 3,
					expected: 7,
				},
				{
					i: 1 , j: 3,
					expected: 8,
				},
				{
					i: 0 , j: 4,
					expected: 9,
				},
				{
					i: 1 , j: 4,
					expected: 10,
				},
			},
		},
	}

	for test_number,test := range tests {

		matrix := Make(test.slice, test.column_length, test.row_length)

		for subtest_number,subtest := range test.subtests {
			actual := matrix.I(subtest.i, subtest.j)

			if subtest.expected != actual {
				t.Errorf("For test #%d subtest #%d expected %f but actually got %f. For i,j = %d,%d and matrix = %v.", test_number, subtest_number, subtest.expected, actual, subtest.i, subtest.j, matrix.slice)
			}
		}
	}
}
