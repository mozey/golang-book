## Structs and Interfaces

### Structs

A type which contains named fields.

### Embedded Types

Is useful for create relationships between structs.

### Interfaces

Are created like structs but using the keyword interface. 

Instead of defining fields, we define a "method set", 
a list of methods that a type must have in order to "implement" the interface.

### Problems

What's the difference between a method and a function?

    Methods are defined on a type

Why would you use an embedded anonymous field instead of a normal named field?

    To create relationships between types

Add a new method to the Shape interface called perimeter which calculates 
the perimeter of a shape. Implement the method for Circle and Rectangle.

    9-structs-interfaces.go
    
    