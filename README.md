Codeship Golang Best Practices
=============================

## Motivation

To achieve consistency and simplicity throughout our Go codebases. This is to aid in readability, maintainability and efficiency when onboarding and during code review.

This document is built upon and inspired by a number of articles and talks given over the past few years by the Go community. These practices are not unique, [everything is borrowed](https://www.youtube.com/watch?v=j8BHL5SWX0Q).

## Contents

- [testing](./testing)
- [errors](./errors)
- [concurrency](./concurrency)
- [makefiles](./makefiles)
- [dependency management](./dependencies)

## Resources

## Extra Goodies

We run markdown-spellcheck in CI on all contributions. Failure to amend such errors will result in changes being rejected.

`make spellcheck` to get quick validation

`make fix-spelling` to use `mdspell` interactive correction functionality and to maintain the `.spelling` file.

see output and see [mdspell](https://github.com/lukeapage/node-markdown-spellcheck) for details on how to maintain the `.spelling` file.

## Contributing

Everyone interacting in the project and its sub-projects' codebases, issue trackers, chat rooms, and mailing lists is expected to follow the [Code of Conduct](CODE_OF_CONDUCT.md).
