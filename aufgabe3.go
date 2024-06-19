package main

// lookup
func sumArea_Lookup_Variant(x, y interface{}) int {
	// we assume the backend supports type assertions
	z := y.(square)

	return area_Lookup(x) + area_Lookup(y) + z.length
}

// dictionary
func sumArea_Dict_Variant(x, y shape_Value) int {
	// we assume the backend supports type assertions
	z := y.val.(square)

	return x.area(x.val) + y.area(y.val) + z.length
}
