Errors - Definition and Handling
================================

1. We strive to create sensibly named root error types which adheres to `golint` standards for errors

For example:

```go
var (
    // ErrSomethingBadHappened is returned when something bad happened
    ErrSomethingBadHappened = errors.New("something bad happened")
)
```

2. We use enrich errors with contextual information where possible. This is done via the `errors.Wrap` family of operations. The wrapping message should contain information with regards to the surrounding function, rather than focusing on the function being called. With an expectation that the error being returned contains all the necessary context.

For example:

```go
func Process(thing Thing) error {
    if err := util.DoSomething(thing.Property); err != nil {
        return errors.Wrapf(err, "processing %q", thing)
    }
}
```

The expectation is that the error returned from `util.DoSomething(...)` contains any relevant context with regards to the `DoSomething` function. Where at the point of collecting the error in `Process(...)` we have the opportunity to enrich the error with context regarding the fact we are processing a particular `Thing` when we encountered this error.

3. All error messages are lower-case

```go
return errors.New("always in lower-case")
```

## Third Party Tools

- [pkg/errors](https://github.com/pkg/errors) - Dave Cheney's error handling primitives package
