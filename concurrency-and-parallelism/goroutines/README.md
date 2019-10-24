# Goroutines

Goroutines and Channels support CSP (communicating sequential processes) a model of concurrency in which values are passed between independent activities (goroutines) but variables are for the most part confined to a single activity.

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
