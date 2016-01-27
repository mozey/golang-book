### Types vs tokens

Suppose you have a dog named Max.

    Max is the token (a particular instance or member) 
    dog is the type (the general concept)


### Sets in mathematics

For example:
 
    ℝ => the set of all real numbers
    ℕ => the set of all natural numbers

Sets are similar to types in programming languages since all the 
values of a particular type share certain properties


### Statically typed

Go is a statically typed programming language. 
This means that variables always have a specific type and that type cannot change.


### Base types

The types listed below are the simplest types included with Go and 
form the foundation from which all later types are built.


### Numbers

#### Integers

Are numbers without a decimal component.

Go has the following unsigned (uint) and signed integer types.
Unsigned integers only contain positive numbers (or zero).

    uint8, uint16, uint32, uint64, int8, int16, int32 and int64
    
There are two alias types: 

    byte => uint8 => 8 bits 
    rune => int32
    
There are also 3 machine dependent integer types, their size 
depends on the type of architecture you are using:

    uint, int and uintptr
    
Generally if you are working with integers you should just use the int type.


#### Floating Point Numbers

Are numbers that contain a decimal component (real numbers).
 
Go has two floating point types: 

    float32 and float64
    
Two additional types for representing complex numbers 
(numbers with imaginary parts)

    complex64 and complex128
    
Floating point numbers are inexact
    
    1.01 - 0.99 = 0.020000000000000018
    
In addition to numbers there are several other values which can be represented

    NaN // Not a number, for things like 0/0)
    +∞  // positive infinity
    −∞  // negative infinity
    
Generally we should stick with float64 when working with floating point numbers


### Strings

Sequence of characters with a definite length used to represent text.

Made up of individual bytes, usually one for each character.
Other languages like Chinese may use more than one byte per character.

The [ascii table](https://en.wikipedia.org/wiki/ASCII#ASCII_printable_code_chart)
and list of [unicode characters](https://en.wikipedia.org/wiki/List_of_Unicode_characters)

The fmt "printing verbs" are explained [here](https://golang.org/pkg/fmt/)


### Booleans

1 bit integer type to represent true and false.


### Problems

How are integers stored on a computer?

    In binary
    
We know that (in base 10) the largest 1 digit number is 9 and the largest 2 
digit number is 99. Given that in binary the largest 2 digit number is 11 (3), 
the largest 3 digit number is 111 (7) and the largest 4 digit number is 1111 
(15) what's the largest 8 digit number? (hint: 101-1 = 9 and 102-1 = 99)

    255
    
Although overpowered for the task you can use Go as a calculator. 
Write a program that computes 32132 × 42452 and prints it to the terminal. 
(Use the * operator for multiplication)

    1364067664

What is a string? How do you find its length?

    A sequence of characters with definite length used to represent text.
    
    Made up of individual bytes, usually one for each character.

What's the value of the expression (true && false) || 
(false && true) || !(false && false)?

    true
    
   
    
