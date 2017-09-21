Makefiles
=========

Use a common [Makefile](https://en.wikipedia.org/wiki/Makefile) to reduce repetitive tasks.

This can be iterated on/added to as needed on a per project basis, but having a generic one as a good foundation will make it easier for developers to move between projects.

An example can be found at: [examples/Makefile](examples/Makefile)

```
❯ make help
setup                          Install all the build and lint dependencies
dep                            Run dep ensure and prune
test                           Run all the tests
cover                          Run all the tests and opens the coverage report
fmt                            Run goimports on all go files
lint                           Run all the linters
ci                             Run all the tests and code checks
build                          Build a version
clean                          Remove temporary files
```
