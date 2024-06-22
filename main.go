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
	// init organizations
	var a association = association{2000, 18}
	var c company = company{350}

	//init dict persons with value and wrapper func
	aDictPersons := organization_Value{a, persons_association_wrapper}
	cDictPersons := organization_Value{c, persons_company_wrapper}

	fmt.Println("SumPersons_Lookup:")
	fmt.Println("The Sum of the persons in a association{2000,18} and a company{350} is: ", sumPersons_Lookup(a, c))
	fmt.Println("SumArea_Lookup with type bound:")
	fmt.Println("The Sum of the persons in a association{2000,18} and a company{350} is: ", sumPersons_Dict(aDictPersons, cDictPersons))
}

func aufgabe3() {
	// Actual Test
	// init shapes
	var r rectangle = rectangle{1, 2}
	var s square = square{3}

	//init dict shapes with value and wrapper func
	rDictShape := shape_Value{r, area_Rec_Wrapper}
	sDictShape := shape_Value{s, area_Sq_Wrapper}

	fmt.Println("SumArea_Lookup with type assertion:")
	fmt.Println("The Sum of the areas of rectangle{1,2} and square{3} is: ", sumArea_Lookup_Variant(r, s))
	fmt.Println("SumArea_Lookup with type bound:")
	fmt.Println("The Sum of the areas of rectangle{1,2} and square{3} is: ", sumArea_Dict_Variant(rDictShape, sDictShape))
}

func aufgabe4() {
	// with method overloading
	i1 := integer{1}
	i2 := integer{2}
	i3 := integer{3}
	i4 := integer{4}

	n1 := node[integer]{i1, nil}
	n2 := node[integer]{i2, nil}
	n3 := node[integer]{i3, nil}
	n4 := node[integer]{i4, nil}

	n1.next = &n2
	n2.next = &n3
	n3.next = &n4

	// default variant with method overloading and interfaces
	fmt.Println("Shownode with method overrides and interfaces:")
	fmt.Println(showNode[integer](&n1))

	// lookup variant without interfaces / method overloading
	fmt.Println("Shownode lookup without method overloading and interfaces:")
	fmt.Println(showNodeLookup[integer](&n1))

	// dictionary variant without interfaces / method overloading
	d1 := show_value{integer{1}, show_integer_wrapper}
	d2 := show_value{integer{2}, show_integer_wrapper}
	d3 := show_value{integer{3}, show_integer_wrapper}
	d4 := show_value{integer{4}, show_integer_wrapper}

	nd1 := node[show_value]{d1, nil}
	nd2 := node[show_value]{d2, nil}
	nd3 := node[show_value]{d3, nil}
	nd4 := node[show_value]{d4, nil}

	nd1.next = &nd2
	nd2.next = &nd3
	nd3.next = &nd4

	fmt.Println("Shownode with dictionary translation:")
	fmt.Println(showNodeDict(&nd1))
}

func main() {
	fmt.Println("--------------------------Aufgabe 1----------------------------------")
	aufgabe1()
	fmt.Println("--------------------------Aufgabe 2----------------------------------")
	aufgabe2()
	fmt.Println("--------------------------Aufgabe 3----------------------------------")
	aufgabe3()
	fmt.Println("--------------------------Aufgabe 4----------------------------------")
	aufgabe4()
}
