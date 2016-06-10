### Problems

What is whitespace?

    Newlines, spaces and tabs
    
    Go (mostly) ignores whitespace
    
Some characters are not whitespace, 
[and your text editor might not display them]
(http://stackoverflow.com/questions/2526881/text-editor-capable-of-viewing-invisibles)

What is a comment? What are the two ways of writing a comment?

    Comments are ignored by the Go compiler 
    
    Comments document the code for the benefit of humans
     
    // This is a comment
    
    /* This is also
    a comment */
    
Our program began with package main. 
What would the files in the fmt package begin with?
    
    package fmt
    
We used the Println function defined in the fmt package. 
If we wanted to use the Exit function from the os package what would we need to do?

    import "os"

Modify the program we wrote so that instead of printing Hello World it prints 
"Hello, my name is" followed by your name.

    fmt.Println("Hello, my name is " + "Chris")
    
    