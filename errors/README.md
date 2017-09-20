Errors - Definition and Handling
================================

1. We strive to create sensibly named root error types which adhead to `golint` standards for errors

e.g.

```go
var (
    // ErrSomethingBadHappened is returned when something bad happened
    ErrSomethingBadHappened = errors.New("something bad happened")
)
```

2. We use enrich errors with contextual information where possible

e.g.

```go
if err := process.DoSomething(5); err != nil {
    return errors.Wrapf(err, "doing something with argument %d", 5)
}
```

## Third Party Tools

- [pkg/errors](github.com/pkg/errors) - Dave Cheney's error handling primitives package
