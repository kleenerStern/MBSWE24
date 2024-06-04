package main

type rectangle struct {
	length int
	width  int
}

type square struct {
	length int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (s square) area() int {
	return s.length * s.length
}

func (r *rectangle) scale(x int) {
	r.length = r.length * x
	r.width = r.width * x
}

func (s *square) scale(x int) {
	s.length = s.length * x
}

type shape interface {
	area() int
}

func sumArea(x, y shape) int {
	return x.area() + y.area()
}

// Introducing unique function names for overloaded methods
func area_Rec(r rectangle) int {
	return r.length * r.width
}

func area_Sq(s square) int {
	return s.length * s.length
}

// Run-time method lookup
func area_Lookup(x interface{}) int {
	var y int

	switch v := x.(type) {
	case square:
		y = area_Sq(v)
	case rectangle:
		y = area_Rec(v)
	}
	return y
}

func sumArea_Lookup(x, y interface{}) int {
	return area_Lookup(x) + area_Lookup(y)
}

// Dictionary translation
type shape_Value struct {
	val  interface{}
	area func(interface{}) int
}

func sumArea_Dict(x, y shape_Value) int {
	return x.area(x.val) + y.area(y.val)
}

// wrapper functions
func area_Rec_Wrapper(v interface{}) int {
	return area_Rec(v.(rectangle))
}
func area_Sq_Wrapper(v interface{}) int {
	return area_Sq(v.(square))
}
