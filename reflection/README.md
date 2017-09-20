Reflection
==========

1. Mostly just NO.
2. Reflection is useful for things which unmarshal (like `encoding/json`) and parsing of struct tags, but we rarely ever implement this stuff.
3. A third party tool can use reflection, but if we see them using reflection, we have to ask ourselves why are they using it? Does this give us more or less confidence in that tool.
