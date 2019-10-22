# Goroutines and Channels

Goroutines and Channels support CSP (communicating sequential processes) a model of concurrency in which values are passed between independent activities (goroutines) but variables are for the most part confined to a single activity.

## Goroutines

Each concurrently executing activity is called a _goroutines_. Goroutines is similar to thread, the differences between these are quantitative and not qualitative.

When a program starts, its only goroutine is the one that calls the _main_ function, new goroutines are created using the _go_ statement.


#### Ordinary Function

```
f() //call f(); wait for it to return
```

#### Goroutine Function

```
go f() //create a new goroutine that calls f(f); don't wait
```

## Channels

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