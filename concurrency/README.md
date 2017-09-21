Concurrency in Go
=================

[I came for the easy concurrency I stayed for the easy composition](https://www.youtube.com/watch?v=woCg2zaIVzQ)

## TLDR

1. It is really hard to do it *correctly*. Try your best to not use it at all.
1. It is really hard to test. Try your best to not use it at all.
1. Test concurrent interaction with the things you write, run those tests with the race detector.

## Concurrency in Go is simple

Doing it correctly is really really hard. It is easy to introduce race conditions, loose control of resources or bring things to a complete standstill.
As often as possible we will try to avoid using it. In fact it is often better to first right something synchronously. Then in the future, extend the functionality with protection against concurrent access, then let the consumer of the type use it concurrently.

When we do use concurrency, what we really want to avoid is leaking the control of concurrent components between types.

What is the meaning of "leaking the control of concurrent components between types"?

This can be occur in many different forms, for example:

1. Channels as either input or return arguments to an exposed function. (return arguments being the worst)
1. Types which expose access to synchronization types i.e. exposing a `sync.Mutex` or `sync.WaitGroup`
1. Types which fire off "unsupervised" goroutines
