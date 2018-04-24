Concurrency in Go
=================

[I came for the easy concurrency I stayed for the easy composition](https://www.youtube.com/watch?v=woCg2zaIVzQ)

## TLDR

1. It is really hard to do it *correctly*. Try your best to not use it at all.
1. It is really hard to test. Try your best to not use it at all.
1. Test concurrent interaction with the things you write, run those tests with the race detector.

## Concurrency in Go is simple

Doing it correctly is really really hard. It is easy to introduce race conditions, lose control of resources or bring things to a complete standstill.
As often as possible we will try to avoid using it. Often our problems can be solved in the following order:

1. Start by defining your code synchronously.
2. Refactor the type to be defensive against concurrent access.
3. Then let consumers of the type use it concurrently.

## Ownership and Responsibility

### **Don't** loose contact with your goroutines

How can this be avoided? `make(chan struct{})`, `sync.WaitGroup`, `context.Context` and `select` are your friends.

It is likely that your type should be:

  1. Able to interrupt the goroutines it spawned when necessary.
  2. Concerned with waiting until all the goroutines it produced are finished.

#### Interruption

This can be achieved by:

  1. Sharing an unbuffered empty struct channel (`make(chan struct{})`) which is closed by the goroutine producer to signal a shutdown.
  2. A cancellable [`context.Context`](https://golang.org/pkg/context/#WithCancel).
  3. Ensuring your goroutines use `select` to check-in on their signal from time to time, without blocking on them indefinitely.

#### Waiting for goroutines to finish

The simplest way to achieve this is with a `sync.WaitGroup`.
**Before** you run that `go routine()`, make sure you `wg.Add(1)`.
**After** you run that `go routine()`, but before it `return`s, make sure you `wg.Done()`. This is where `defer` is your friend.

### **Don't** use wait groups to count more than one type of goroutine

Goroutine type in this scenario relates to the type of the function being called as a goroutine.
This function could be a member of another type, be a named function in the package or it could be anonymous.
The important take-away is that you shouldn't share the WaitGroup among different functions called as goroutines.
Keep it simple and add another WaitGroup if you find yourself calling `go` before a different function and name the WaitGroup's appropriately.

```go
type Parent struct {
  wgFoo sync.WaitGroup
  wgBar sync.WaitGroup
}

func (p *Parent) foo() {
  defer p.wgFoo.Done()
}

func (p *Parent) bar() {
  defer p.wgBar.Done()
}

func (p *Parent) Go() {
  p.wgFoo.Add(1)
  go p.foo()

  p.wgBar.Add(1)
  go.bar()
}
```

Though sharing a WaitGroup may be a correct solution, it adds to the cognitive complexity of a problem when the next engineer comes to grok it.

### **Don't** let a channel consumer say when it is done

> A send on a closed channel will cause a panic.

First and foremost this mandates that the code is modeled as channel __consumers__ and __producers__.
This is a good practice in and of itself. It is a clear separations of concerns.

Go gives you the ability to, at compile time, define the direction of a channel `recvOnly <-chan Thing := make(chan Thing)`.
This is rarely useful when defining a variable, however, it is super useful when defining the receive arguments of a function.
For example:

```go
func consume(things <-chan Thing) {
  // will do work until close
  for thing := range things {
    // do work
  }
}
```

This enforces (at compile time!) that the consumer goroutine **cannot** send on that channel.
This includes the ability to **close** that channel.

The aids in enforcing another tenant of safe channel management. Only close a channel, once all producers have stopped producing. Remember a send on a closed channel will cause a panic.
It is important that you maintain responsibility over producers.

**The piece of code which closes a channel must first guarantee that nothing else will produce on it.**

If all sends on that channel have happened synchronously before the call to close, then you will be safe as long as you don't accidentally try and send again. If production on that channel is relinquished to other goroutines, then you need to be able to synchronize with the exit of these producing routines.

If we did the work to ensure we are counting our routines and waiting for them to exit, then we can be sure that a close won't cause a panic elsewhere.

```go
func doConcurrently() {
  var (
    things   = make(chan Thing)
    finished = make(chan struct{})
    wg       sync.WaitGroup
  )

  go func() {
    // will consume until close
    consume(things)
    // signal consumption has finished
    close(finished)
  }()

  for i := 0; i < noOfThingsWeWantToDo; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()

      things <- Thing{}
    }()
  }
  
  // wait until all producers have stopped
  wg.Wait()

  // then you can close
  close(things)

  // wait until finished consuming
  <-finished
}
```

## Summary

1. Ensure consumers can only consume. `recvOnly <-chan Thing` are your friends.
2. Track completion of goroutines. `sync.WaitGroup` is your friend.
3. Close only when producing routines can be verified as no longer able to send on the channel being closed.
