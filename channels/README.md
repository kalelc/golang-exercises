# Channels

If goroutines are the activities of a concurrent Go program, channels are the connections between them. A channel is a communication mechanism that lets one goroutine send values to another goroutine. Each channel is a conduit for values of a particular type, called the channel’s _element type_.

```
ch := make(chan int) // ch has type 'chan int'
```

As with maps, a channel is a reference to the data structure created by make. When we copy a channel or pass one as an argument to a function, we are copying a reference, so caller and callee refer to the same data structure. As with other reference types, the zero value of a channel is nil.

Two channels of the same type may be compared using ==. The comparison is true if both are references to the same channel data structure. A channel may also be compared to nil.

A channel has two principal operations, _send_ and _receive_, collectively known as communications. A send statement transmits a value from one goroutine, through the chan- nel, to another goroutine executing a corresponding receive expression. Both operations are written using the <- operator. In a send statement, the <- separates the channel and value operands. In a receive expression, <- precedes the channel operand. A receive expression whose result is not used is a valid statement.

```
ch <- x // a send statement
x = <-ch // a receive expression in an assignment statement <-ch // a receive statement; result is discarded
```

Channels support a third operation, _close_, which sets a flag indicating that no more values will ever be sent on this channel; subsequent attempts to send will panic.

closed channel yield the values that have been sent until no more values are left; any receive operations thereafter complete immediately and yield the zero value of the channel’s element type.

To close a channel, we call the built-in close function: _close(ch)_.

A channel created with a simple call to make is called an unbuffered channel, but make accepts an optional second argument, an integer called the channel’s capacity. If the capacity is non- zero, make creates a buffered channel.

```
ch = make(chan int)    // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

### Unbuffered Channels

A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel, at which point the value is transmitted and both goroutines may continue. Conversely, if the receive operation was attempted first, the receiving goroutine is blocked until another goroutine performs a send on the same channel.

Communication over an unbuffered channel causes the sending and receiving goroutines to synchronize. Because of this, unbuffered channels are sometimes called synchronous channels. When a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the sending goroutine.

### Deadlock

One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.

Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.

### Buffered Channels

A send operation in a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.

Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.

```
ch := make(chan type, capacity) 
```

### Select Operation

The select statement is used to choose from multiple send/receive channel operations. The select statement blocks until one of the send/receive operation is ready. If multiple operations are ready, one of them is chosen at random. The syntax is similar to switch except that each of the case statement will be a channel operation.

#### Default case
The default case in a select statement is executed when none of the other case is ready. This is generally used to prevent the select statement from blocking.