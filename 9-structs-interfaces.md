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

Note that dereferencing is not required when 
[accessing fields in the method](http://stackoverflow.com/a/30786320/639133)

    Methods are defined on a type.

Why would you use an embedded anonymous field instead of a normal named field?
[Tutorial](http://golangtutorials.blogspot.co.za/2011/06/anonymous-fields-in-structs-like-object.html)

    To create relationships between types. It allows us to access members of the 
    anonymous field as if they were members of the encompassing class, e.g.
    
    type Kitchen struct {
        numOfPlates int 
    }
    type House struct {
        Kitchen //anonymous field
        numOfRooms int 
    }
    h := House{Kitchen{10}, 3}
    fmt.Println(h.numOfPlates) // 10
    fmt.Println(h.Kitchen) // {10}

Add a new method to the Shape interface called perimeter which calculates 
the perimeter of a shape. Implement the method for Circle and Rectangle.

    9-structs-interfaces.go
    
    