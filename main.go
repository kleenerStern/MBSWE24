package main

import (
	"fmt"
	"math"
	"time"
)

func printResults(execTimeDict time.Duration, execTimeLookup time.Duration) {
	fmt.Printf("total execution time for 1000 * lookup: %s \n", execTimeLookup)
	fmt.Printf("mean execution time for lookup: %s \n", execTimeLookup/1000)
	fmt.Printf("total execution time for 1000 * dict: %s \n", execTimeDict)
	fmt.Printf("mean execution time for dict: %s \n", execTimeDict/1000)

	percentageChange := (float64(execTimeDict) - float64(execTimeLookup)) / float64(execTimeLookup) * 100

	fmt.Printf("The execution of dict took %.2f%% ", math.Abs(percentageChange))
	if percentageChange > 0 {
		fmt.Println("more time than the execution of lookup.")
	} else if percentageChange < 0 {
		fmt.Println("less time than the execution of lookup.")
	} else {
		fmt.Println("the same time as the execution of lookup.")
	}
}

func aufgabe1() {
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

	printResults(executionTimeDict, executionTimeLookup)
}

func aufgabe2() {
	// Actual Test
	// init shapes
	var r rectangle = rectangle{1, 2}
	var s square = square{3}

	//init dict shapes with value and wrapper func
	rDictShapeExt := shapeExt_Value{r, area_Rec_Wrapper, perimeter_Rec_Wrapper}
	sDictShapeExt := shapeExt_Value{s, area_Sq_Wrapper, perimeter_Sq_Wrapper}

	startingTimeLookup := time.Now()
	time.Sleep(1 * time.Second)
	for i := 0; i < 1000; i++ {
		sumPerimeter_Lookup(r, s)
	}
	executionTimeLookup := time.Since(startingTimeLookup) - 1*time.Second

	startingTimeDict := time.Now()
	time.Sleep(1 * time.Second)
	for i := 0; i < 1000; i++ {
		sumPerimeter_Dict(rDictShapeExt, sDictShapeExt)
	}
	executionTimeDict := time.Since(startingTimeDict) - 1*time.Second

	printResults(executionTimeDict, executionTimeLookup)
}

func main() {
	fmt.Println("--------------------------Aufgabe 1----------------------------------")
	aufgabe1()
	fmt.Println("--------------------------Aufgabe 2----------------------------------")
	aufgabe2()
}
