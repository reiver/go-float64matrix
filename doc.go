/*
Package float64maxtrix provides minimal help to interpret a `[]float64` as a matrix.

So that you can do things such as:

	slice := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	
	// 2×4 matrix.
	//
	// I.e.,:
	//
	//	| 1.0  3.0  5.0  7.0 |
	//	| 2.0  4.0  6.0  8.0 |
	//
	matrix := float64maxtrix.Make(slice, 2, 4)
	
	
	// Get the element of the matrix at i,j.
	i,j := 0,3
	x := matrix.I(i,j)
	// x == 7.0
	
	
	// Set the element of the matrix at ii,jj, to 1000.0.
	ii,jj := 1,2
	value := 1000.0
	matrix.Set(value, ii,jj)
	// Matrix is now
	//
	//	| 1.0  3.0     5.0  7.0 |
	//	| 2.0  4.0  1000.0  8.0 |
	//
	
	
	// Get column 2.
	//
	// I.e.,
	//
	//	|    5.0 |
	//	| 1000.0 |
	//
	column2 := matrix.Column(2)

This has applications to most things that use linear algebra.

Motivation

Go does not provide native support for multi-dimensional indices on slices and arrays. So things
such as the following are not possible:

	// NOTE: This will NOT work in Go!
	data[1,2] = 14.4

One alternative that looks cosmetically similar (but that you should avoid if you care about certain
kinds of performance) is using a 'slice of slices' or an 'array of arrays'. For example:

	data[1][2] = 14.4

Although this and the previous code sample (cosmetically) look similar, the performance difference
between the two can be very significant!

The former code (if Go actually supported it) could be used to to create (what are called)
'cache-oblivious algorithms' (or 'cache-transcendent algorithms') that can be optimized to take
advantage of CPU caches, without knowing the cache size or the length or the cache line.

The former code (if Go actually supported it) could do this because the memory space for the
data structure is a single unified contiguous block of memory.

However, the latter code cannot do this! 'Slice of slices' and 'array of arrays'
are (potentially) made up of multiple discontiguous blocks of memory.

These type of poor performance characteristics (of 'slice of slices' and 'array of arrays') can be
unacceptable in certain problem spaces.

Simulating Multi-Dimensional Indices

This package provides are way to simulate multi-dimensional indexing on slices and arrays in Go,
so that a programmer can create (what are called) 'cache-oblivious algorithms' (or 'cache-transcendent algorithms')
that can be optimized to take advantage of CPU caches, without knowing the cache size or the length or the cache line

*/
package float64maxtrix
