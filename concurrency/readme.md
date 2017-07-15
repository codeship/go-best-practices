Concurrency in Go
=================

[I came for the easy concurrency I stayed for the easy composition](https://www.youtube.com/watch?v=woCg2zaIVzQ)

## TLDR

1. It is really hard to do it *correctly*. Try your best to not use it at all.
2. It is really hard to test. Try your best to not use it at all.
3. Test concurrent interaction with the things you write, run those tests with the race detector.

## Concurrency in Go is simple

Doing it correctly is really really hard. It is easy to introduce race conditions, loose control of resources or bring things to a complete standstill.
