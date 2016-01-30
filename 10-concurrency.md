## Concurrency

Large programs are often made up of many smaller sub-programs.

It would be ideal for programs like these to be able 
to run their smaller components at the same time.


## Goroutines

Function that is capable of running concurrently with other functions.
 
 
## Channels

Provide a way for two goroutines to communicate and synchronize their execution.

<- (left arrow operator)

    Used to send and receive messages on the channel
    
    
## Channel Direction

We can specify a direction on a channel type 
thus restricting it to either sending or receiving:
 
    func pinger(c chan<- string) // Sender
    
    func printer(c <-chan string) // Receiver
 
Sending on a receive only channel or vice versa results in a compiler error:
 
    ... send to receive-only type ...
    
An unrestricted channel is called bi-directional, 
they can be passed into functions that take send or receive only channels.
    
    
## Select

Works like a switch but for channels.

Select picks the first channel that is ready and receives from / send to it,
if more than one of the channels are ready then it randomly picks one.

It blocks until a channel is available for send / receive.


## Buffered Channels

Pass a second parameter (buffer size) to make when creating a channel,

    c := make(chan int, 1)

a buffered channel is asynchronous; sending or receiving a message 
will not wait unless the channel is already full.


## Problems

How do you specify the direction of a channel type?

    By appending chan<- (send only) 
    or prefixing <-chan (receive only) 
    the left arrow operator.

Write your own Sleep function using time.After.

    10-concurrency-sleep.go

What is a buffered channel? How would you create one with a capacity of 20?

    An asynchronous channel

    c := make(chan int, 20)
    
 
