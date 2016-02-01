## Testing

The `go test` command will look for any tests in 
any of the files in the current folder and run them.

Test functions must start with the word "Test"
and take on argument of type `*testing.T`

A common test pattern is to create a struct that represents the inputs and 
outputs of the function, populate a variable with values and loop over it.


### Problems

Writing a good suite of tests is not always easy, but the process of writings 
tests often reveals more about a problem then you may at first realize. 
For example, with our Average function what happens if you pass in an 
empty list ([]float64{})? 
How could we modify the function to return 0 in this case?

	if (len(xs) == 0) {
		return 0
	}

Write a series of tests for the Min and Max 
functions you wrote in the previous chapter.

    ...
    
    