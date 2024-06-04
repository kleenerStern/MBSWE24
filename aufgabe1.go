package main

import (
	"fmt"
	"math"
	"time"
)

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

func main() {
	// Actual Test
	// init shapes
	var r rectangle = rectangle{1, 2}
	var s square = square{3}

	//init dict shapes with value and wrapper func
	rDictShape := shape_Value{r, area_Rec_Wrapper}
	sDictShape := shape_Value{s, area_Sq_Wrapper}

	startingTimeLookup := time.Now()
	time.Sleep(1 * time.Second)
	for i := 0; i < 1000; i++ {
		sumArea_Lookup(r, s)
	}
	executionTimeLookup := time.Since(startingTimeLookup) - 1*time.Second

	startingTimeDict := time.Now()
	time.Sleep(1 * time.Second)
	for i := 0; i < 1000; i++ {
		sumArea_Dict(rDictShape, sDictShape)
	}
	executionTimeDict := time.Since(startingTimeDict) - 1*time.Second

	fmt.Printf("total execution time for 1000 * lookup: %s \n", executionTimeLookup)
	fmt.Printf("mean execution time for lookup: %s \n", executionTimeLookup/1000)
	fmt.Printf("total execution time for 1000 * dict: %s \n", executionTimeDict)
	fmt.Printf("mean execution time for dict: %s \n", executionTimeDict/1000)

	percentageChange := (float64(executionTimeDict) - float64(executionTimeLookup)) / float64(executionTimeLookup) * 100

	fmt.Printf("The execution of dict took %.2f%% ", math.Abs(percentageChange))
	if percentageChange > 0 {
		fmt.Println("more time than the execution of lookup.")
	} else if percentageChange < 0 {
		fmt.Println("less time than the execution of lookup.")
	} else {
		fmt.Println("the same time as the execution of lookup.")
	}
}
