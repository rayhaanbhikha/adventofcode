package main

var sequences = [][]int{
	{5, 6, 7, 8, 9},
	{6, 5, 7, 8, 9},
	{7, 5, 6, 8, 9},
	{5, 7, 6, 8, 9},
	{6, 7, 5, 8, 9},
	{7, 6, 5, 8, 9},
	{8, 6, 5, 7, 9},
	{6, 8, 5, 7, 9},
	{5, 8, 6, 7, 9},
	{8, 5, 6, 7, 9},
	{6, 5, 8, 7, 9},
	{5, 6, 8, 7, 9},
	{5, 7, 8, 6, 9},
	{7, 5, 8, 6, 9},
	{8, 5, 7, 6, 9},
	{5, 8, 7, 6, 9},
	{7, 8, 5, 6, 9},
	{8, 7, 5, 6, 9},
	{8, 7, 6, 5, 9},
	{7, 8, 6, 5, 9},
	{6, 8, 7, 5, 9},
	{8, 6, 7, 5, 9},
	{7, 6, 8, 5, 9},
	{6, 7, 8, 5, 9},
	{9, 7, 8, 5, 6},
	{7, 9, 8, 5, 6},
	{8, 9, 7, 5, 6},
	{9, 8, 7, 5, 6},
	{7, 8, 9, 5, 6},
	{8, 7, 9, 5, 6},
	{5, 7, 9, 8, 6},
	{7, 5, 9, 8, 6},
	{9, 5, 7, 8, 6},
	{5, 9, 7, 8, 6},
	{7, 9, 5, 8, 6},
	{9, 7, 5, 8, 6},
	{9, 8, 5, 7, 6},
	{8, 9, 5, 7, 6},
	{5, 9, 8, 7, 6},
	{9, 5, 8, 7, 6},
	{8, 5, 9, 7, 6},
	{5, 8, 9, 7, 6},
	{5, 8, 7, 9, 6},
	{8, 5, 7, 9, 6},
	{7, 5, 8, 9, 6},
	{5, 7, 8, 9, 6},
	{8, 7, 5, 9, 6},
	{7, 8, 5, 9, 6},
	{6, 8, 5, 9, 7},
	{8, 6, 5, 9, 7},
	{5, 6, 8, 9, 7},
	{6, 5, 8, 9, 7},
	{8, 5, 6, 9, 7},
	{5, 8, 6, 9, 7},
	{9, 8, 6, 5, 7},
	{8, 9, 6, 5, 7},
	{6, 9, 8, 5, 7},
	{9, 6, 8, 5, 7},
	{8, 6, 9, 5, 7},
	{6, 8, 9, 5, 7},
	{6, 5, 9, 8, 7},
	{5, 6, 9, 8, 7},
	{9, 6, 5, 8, 7},
	{6, 9, 5, 8, 7},
	{5, 9, 6, 8, 7},
	{9, 5, 6, 8, 7},
	{9, 5, 8, 6, 7},
	{5, 9, 8, 6, 7},
	{8, 9, 5, 6, 7},
	{9, 8, 5, 6, 7},
	{5, 8, 9, 6, 7},
	{8, 5, 9, 6, 7},
	{7, 5, 9, 6, 8},
	{5, 7, 9, 6, 8},
	{9, 7, 5, 6, 8},
	{7, 9, 5, 6, 8},
	{5, 9, 7, 6, 8},
	{9, 5, 7, 6, 8},
	{6, 5, 7, 9, 8},
	{5, 6, 7, 9, 8},
	{7, 6, 5, 9, 8},
	{6, 7, 5, 9, 8},
	{5, 7, 6, 9, 8},
	{7, 5, 6, 9, 8},
	{7, 9, 6, 5, 8},
	{9, 7, 6, 5, 8},
	{6, 7, 9, 5, 8},
	{7, 6, 9, 5, 8},
	{9, 6, 7, 5, 8},
	{6, 9, 7, 5, 8},
	{6, 9, 5, 7, 8},
	{9, 6, 5, 7, 8},
	{5, 6, 9, 7, 8},
	{6, 5, 9, 7, 8},
	{9, 5, 6, 7, 8},
	{5, 9, 6, 7, 8},
	{8, 9, 6, 7, 5},
	{9, 8, 6, 7, 5},
	{6, 8, 9, 7, 5},
	{8, 6, 9, 7, 5},
	{9, 6, 8, 7, 5},
	{6, 9, 8, 7, 5},
	{7, 9, 8, 6, 5},
	{9, 7, 8, 6, 5},
	{8, 7, 9, 6, 5},
	{7, 8, 9, 6, 5},
	{9, 8, 7, 6, 5},
	{8, 9, 7, 6, 5},
	{8, 6, 7, 9, 5},
	{6, 8, 7, 9, 5},
	{7, 8, 6, 9, 5},
	{8, 7, 6, 9, 5},
	{6, 7, 8, 9, 5},
	{7, 6, 8, 9, 5},
	{7, 6, 9, 8, 5},
	{6, 7, 9, 8, 5},
	{9, 7, 6, 8, 5},
	{7, 9, 6, 8, 5},
	{6, 9, 7, 8, 5},
	{9, 6, 7, 8, 5},
}
