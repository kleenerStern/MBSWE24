package main

type shapeExt interface {
	shape
	perimeter()
}

func (r rectangle) perimeter() int {
	return 2 * (r.length + r.width)
}

func (s square) perimeter() int {
	return 4 * s.length
}

// Introducing unique function names for overloaded methods
func perimeter_Rec(r rectangle) int {
	return r.length * r.width
}

func perimeter_Sq(s square) int {
	return s.length * s.length
}

// Run-time method lookup
func perimeter_Lookup(x interface{}) int {
	var y int

	switch v := x.(type) {
	case square:
		y = perimeter_Sq(v)
	case rectangle:
		y = perimeter_Rec(v)
	}
	return y
}

func sumPerimeter_Lookup(x, y interface{}) int {
	return perimeter_Lookup(x) + perimeter_Lookup(y)
}

// Dictionary translation
type shapeExt_Value struct {
	val       interface{}
	area      func(interface{}) int
	perimeter func(interface{}) int
}

func sumPerimeter_Dict(x, y shapeExt_Value) int {
	return x.perimeter(x.val) + y.perimeter(y.val)
}

// wrapper functions
func perimeter_Rec_Wrapper(v interface{}) int {
	return perimeter_Rec(v.(rectangle))
}
func perimeter_Sq_Wrapper(v interface{}) int {
	return perimeter_Sq(v.(square))
}
