## Arrays, Slices and Maps

### Arrays

Numbered sequence of elements of a single type with a fixed length.


### Slices

A slice is a segment of an array, slices are index-able and have a length. 

Unlike arrays the slice length is allowed to change.
 
Slices are always associated with some array, 
although they can never be longer than the array, they can be smaller.


### Slice function

append
    
    Creates a new slice by taking an existing slice (the first argument) 
    and appending all the following arguments to it.
    
copy(slice2, slice1)

    Copies the first len(slice2) elements from slice1 into slice2
      
      
### Maps

A map is an unordered collection of key-value pairs,
a.k.a associative array, hash table or dictionary.

The difference between arrays and maps:
 
    The length of a map can change as we add new items to it.
    
    Maps are not sequential, m[1] does not imply m[0] exists.
    
    
### Problems

How do you access the 4th element of an array or slice?

    i := 4 - 1
    a[i]

What is the length of a slice created using: make([]int, 3, 9)?

    3

Given the array:
x := [6]string{"a","b","c","d","e","f"}
what would x[2:5] give you?

    [c d e]

Write a program that finds the smallest number in this list:
x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}

    6-arrays-slices-maps-smallest.go







