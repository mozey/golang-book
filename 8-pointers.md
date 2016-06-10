## Pointers

Pointers reference a location in memory where 
a value is stored rather than the value itself.
 
 
### The * and & operators

&lowast; (Pointer dereference operator)    

- Find or modify the value of a pointer

- Dereferencing a pointer gives us access to the value the pointer points to
        
&amp;

- Find the address of a variable

- &x returns a *int (pointer to int) if x is an int


### New

Another way to get a pointer.

Takes a type as an argument, 
allocates enough memory to fit a value of that type 
and returns a pointer to it.

New does not require any additional garbage collection from the programmer, 
memory is cleaned up automatically when nothing refers to it anymore.

Pointer are rarely used with built in types,
they are most useful when paired with structs.


### Problems

How do you get the memory address of a variable?

    &variable

How do you assign a value to a pointer?

    *pointer = "the value"

How do you create a new pointer?

    new(Type)

What is the value of x after running this program:

    func square(x *float64) {
      *x = *x * *x
    }
    func main() {
      x := 1.5
      square(&x)
    }
    
    // ANSWER: x == 2.25
    
Write a program that can swap two integers 
(x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

    8-pointers.go
    

