Codeship Golang Best Practices
=============================

## Motivation

To achieve consistency and simplicity throughout our Go codebases. This is to aid in readability, maintainability and efficiency when onboarding and during code review.

This document is built upon and inspired by a number of articles and talks given over the past few years by the Go community. These practices are not unique, [everything is borrowed](https://www.youtube.com/watch?v=j8BHL5SWX0Q).

## Contents

- [testing](./testing)
- [concurrency](./concurrency)
- [makefiles](./makefiles)
- [dependency management](./dependencies)

## Resources

### Articles

### Talks

## Extra Goodies

We run markdown-spellcheck in CI on all contributions. Failure to amend such errors will result in changes being rejected.

run `npm test` to get quick validation

run `npm run-script fix` to use `mdspell` interactive correction functionality and to maintain the `.spelling` file.

see output and see [mdspell](https://github.com/lukeapage/node-markdown-spellcheck) for details on how to maintain the `.spelling` file.
