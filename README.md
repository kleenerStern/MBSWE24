# Modellbasierte Softwareentwicklung Projekt SS24

## Aufgabe 1
To compare the run-time performance of the sumArea function using run-time method look up and dictionary translation, we used the standard time library and for loops to track execution times of both variants.  
But since this method of performance testing did not yield persistent comparable data, we also implemented GO Benchmark tests, found in "aufgabe1_test.go".  

The results were the following:  
When executing both operations (RT / DT) 1000000 times, RT only took 0.6223ns/execution and DT took 5.803ns/exec.
So in our test case, run time method look up was 832.4% faster than dictionary translation.

## Aufgabe 2-4
Tasks 2-4 did not yield any significant new insights, so we are not sure if there is anything to document here, or if this was intended in general.