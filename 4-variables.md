### Variables

A variable is a storage location, with a specific type and an associated name.


### How to Name a Variable

Lower camel case is preferred, for example

    dogsName
    
    sprayTanCanister

    
### Scope

Go is lexically scoped using blocks,
this means that the variable exists within the nearest curly braces { }
(a block) including any nested curly braces (blocks), but not outside of them.


### Constants

Variables whose values cannot be changed later, created using "const" keyword.


### Problems

What are two ways to create a new variable?

    var and :=
    
What is the value of x after running: 
x := 5; x += 1?

    6
    
What is scope and how do you determine the scope of a variable in Go?

    Scope determines variable visibility,
    variables are visible inside the nearest block 
    and inside any nested blocks.
    
What is the difference between var and const?

    Variables declared with var can be changed.
    
Using the example program as a starting point, 
write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)

    4-variables-fah-to-cel.go
    
Write another program that converts from feet into meters. (1 ft = 0.3048 m)

    4-variables-ft-to-m.go
    

    

    
    
    






