package main

import "testing"

var rTest1 rectangle = rectangle{1, 2}
var sTest1 square = square{3}

// init dict shapes with value and wrapper func
var rDictShapeTest1 = shape_Value{rTest1, area_Rec_Wrapper}
var sDictShapeTest1 = shape_Value{sTest1, area_Sq_Wrapper}

func BenchmarkSumArea_Lookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumArea_Lookup(rTest1, sTest1)
	}
}

func BenchmarkSumArea_Dict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumArea_Dict(rDictShapeTest1, sDictShapeTest1)
	}
}
