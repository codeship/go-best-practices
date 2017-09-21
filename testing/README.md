Testing in Go
=============

## The Minimum Standard

*80%* Coverage (preferably 100% happy path)

## TLDR

1. `testing` package is your friend.
1. `https://github.com/stretchr/testify` assertions and requirements are more than welcome.
1. table driven tests are often worthwhile.
1. tests will dictate how the production code is structured.
1. the minimum standard of coverage will be upheld.
1. we have decided against the use of mocking libraries

## Unit Testing

Most practical ideas have been stolen from Mitchell Hashimoto's very comprehensive talk on advanced testing in Go [video](1).

### Tests must be written before code reaches master

Because tests can't be written later down the line. The way in which tests are written dictates the way in which the go we are testing is written.

Any code that doesn't have suitable testing is subject to being thrown away and written from scratch with tests. In fact it is urged that is done as soon as possible.

### Use the standard libraries

The `testing` package is your friend. It comes equipped with nearly everything you need to raise failures, create sub tests and perform suite setup and tear down.

### We like testify

Simple comparisons are good enough to test with. However, it can get tedious and inconsistent to write our own failure messages. `assert` and `require` reduce the noise in a test and provide nicely formatted default failure messages. Plus it works very well with the standard libraries.

### Table driven tests

1. Use subtests to give each case of the table its own context
1. Use name field in your table case to be more descriptive
1. Use comments in the struct definition to further elaborate a fields intent

#### examples

1. [fibonacci](./examples/table-driven)

## Mocking

Currently our stance is simple. We don't use any third party libraries for mocking. Instead we favor hand-rolled mocks, dummies, noops, fakes and spies.

Let me give an example:

Say that we have an interface which when supplied with a "container ID", it is intended to stop that container from running. The interface for this functionality might look like this:

```go
type ContainerStopper interface {
    Stop(containerID string) error
}
```

Instead of turning to a mocking library, all sorts of hand-rolled tricks can be employed to swap this behavior out.

```go
type spyContainerStopper struct {
    id     string
    err    error
}

func (s spyContainerStopper) Stop(id string) error {
    s.id = id
    return n.err
}

type containerStopperFunc func(string) error

func (c containerStopperFunc) func(s string) error { return c(s) }
```

The first example above is the simple spy-like implementation which captures the input `id` on calls to `Stop(...)`. It also returns the error configured on the struct in the `err` field.

The second allows for simple anonymous functions to be used in a test case as an implementation. Note that this is often only useful for smaller interfaces, often single function (but your interfaces should be small anyway). The following demonstrates an anonymous function which achieves the same ends as the struct implementation.

```go
func TestSomething(t *testing.T) {
    var (
        capturedID string
        spy        ContainerStopper = containerStopperFunc(func(id string) error {
            capturedID = id
            return errors.New("something went wrong")
        })
    )

    //...
}
```

If you are dealing with a big pesky interface, it is recommended that you use something like [impl](https://github.com/josharian/impl) to generate a skeleton and speed things up.

## Acceptance Testing

> TODO: Here we will reach out in to existing Go projects and fill out this section with what others are doing well.

[1]:[https://www.youtube.com/watch?v=yszygk1cpEc]
