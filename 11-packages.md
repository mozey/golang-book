## Packages

### Creating packages

Compile math.go and create a linkable object file

    cd $GOPATH/src/github.com/mozey/golang-book/11-packages/math

    go install

    ls $GOPATH/pkg/darwin_amd64/github.com/mozey/golang-book/11-packages/

    math.a
    

### Documentation

Displays the comment right above the math.Average function

    godoc github.com/mozey/golang-book/11-packages/math Average
    
Docs are also available in web form, 
to browse all packages installed on the system:

    godoc -http=":6060"
    
    http://localhost:6060/pkg/
    
    
### Problems

Why do we use packages?

    To observe DRY

What is the difference between an identifier that starts with a capital letter 
and one which doesn't? (Average vs average)

    Average is visible in other packages and programs, average is private 

What is a package alias? How do you make one?

    Let's us use multiple package with the same basename

    import m "github.com/mozey/golang-book/11-packages/math"

We copied the average function from chapter 7 to our new package. 
Create Min and Max functions which find the minimum 
and maximum values in a slice of float64s.

    11-packages/math/math.go

How would you document the functions you created in #3?
    
    Using comments above the function
    
    
    
    
    
    
    