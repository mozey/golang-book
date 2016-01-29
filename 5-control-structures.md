## Control Structures

### For

Repeat a list of statements (a block) multiple times.

Go only has one type of loop that can be used in different ways.
 
 
### If

Similar to a for statement in that it has a condition followed by a block. 
If statements also have an optional else part. If the condition evaluates to 
true then the block after the condition is run, otherwise either the block 
is skipped or if the else block is present that block is run.


### Switch

A switch statement starts with the keyword switch followed by an expression 
and then a series of cases. The value of the expression is compared to the 
expression following each case keyword. If they are equivalent then the 
statement(s) following the : is executed.


### Problems

What does the following program print:

    i := 10
    if i > 10 {
      fmt.Println("Big")
    } else {
      fmt.Println("Small")
    }
    
    => Small
    
    
Write a program that prints out all the numbers 
evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)

    for i := 1; i <= 10; i++ {
        if i % 3 == 0 {
            fmt.Println(i)
        }
    }


Write a program that prints the numbers from 1 to 100. 
But for multiples of three print "Fizz" instead of the number 
and for the multiples of five print "Buzz". For numbers which are 
multiples of both three and five print "FizzBuzz".

    5-control-structures-fizzbuzz.go
        
