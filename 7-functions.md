## Functions

Independent section of code that maps zero or more input parameters 
to zero or more output parameters, a.k.a procedures, subroutines.


### Notes about functions

Names of the parameters don't have to match in the calling function
    
Functions don't have access to anything in the calling function
    
Functions are built up in a "stack", 
each call to return pops the last function off the stack. 

We can also name the return type

   
### Variadic Functions

Are functions with a variable number of arguments.


### Closure

A function inside a function, together with the 
non-local variables it references is known as a closure.


### Recursion

A function that calls itself.

Closure and recursion forms the basis of a 
paradigm known as functional programming.


### Defer

Defer schedules a function call to be run after the function completes.

Often used when resources need to be freed,
for example, closing an open file.


### Panic and Recover

The recover function can be used to recover from a panic.


### Problems

sum is a function which takes a slice of numbers and adds them together. 
What would its function signature look like in Go?

    func sum(args ...int) int
    

Write a function which takes an integer and halves it 
and returns true if it was even or false if it was odd. 
For example half(1) should return (0, false) 
and half(2) should return (1, true).

Write a function with one variadic parameter 
that finds the greatest number in a list of numbers.

Using makeEvenGenerator as an example, 
write a makeOddGenerator function that generates odd numbers.

The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, 
fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).

What are defer, panic and recover? How do you recover from a run-time panic?

    7-functions.go
    
