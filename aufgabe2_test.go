package main

import "testing"

var rTest2 rectangle = rectangle{1, 2}
var sTest2 square = square{3}

// init dict shapes with value and wrapper func
var rDictShapeExtTest2 = shapeExt_Value{rTest2, area_Rec_Wrapper, perimeter_Rec_Wrapper}
var sDictShapeExtTest2 = shapeExt_Value{sTest2, area_Sq_Wrapper, perimeter_Sq_Wrapper}

func BenchmarkSumPerimeter_Lookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumPerimeter_Lookup(rTest2, sTest2)
	}
}

func BenchmarkDict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumPerimeter_Dict(rDictShapeExtTest2, sDictShapeExtTest2)
	}
}
