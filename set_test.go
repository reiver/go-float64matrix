package float64maxtrix


import "testing"

import "math/rand"
import "time"


func TestSet(t *testing.T) {

	rand.Seed( time.Now().UTC().UnixNano() )

	tests := []struct{
		slice []float64
		column_length int
		row_length    int
	}{
		{
			slice: []float64{1,2,3,4,5,6,7,8,9},
			column_length: 3,
			row_length:    3,
		},
		{
			slice: []float64{1,2,3,4,5,6,7,8,9,10},
			column_length: 5,
			row_length:    2,
		},
		{
			slice: []float64{1,2,3,4,5,6,7,8,9,10},
			column_length: 2,
			row_length:    5,
		},
	}

	for test_number,test := range tests {

		matrix := Make(test.slice, test.column_length, test.row_length)

		for i:=0; i<test.column_length; i++ {
			for j:=0; j<test.row_length; j++ {

				expected := rand.Float64()

				matrix.Set(expected, i,j)

				actual := matrix.I(i,j)

				if expected != actual {
					t.Errorf("For test #%d with i,j = %d,%d expected %f but actually got %f. For matrix = %v.", test_number, i, j, expected, actual, matrix.slice)
				}
			}
		}
	}
}
