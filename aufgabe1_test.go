package main

import "testing"

var r rectangle = rectangle{1, 2}
var s square = square{3}

// init wrappers
func area_Rec_Wrapper(v interface{}) int {
	return area_Rec(v.(rectangle))
}
func area_Sq_Wrapper(v interface{}) int {
	return area_Sq(v.(square))
}

// init dict shapes with value and wrapper func
var rDictShape = shape_Value{r, area_Rec_Wrapper}
var sDictShape = shape_Value{s, area_Sq_Wrapper}

func BenchmarkLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumArea_Lookup(r, s)
	}
}

func BenchmarkDict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumArea_Dict(rDictShape, sDictShape)
	}
}
