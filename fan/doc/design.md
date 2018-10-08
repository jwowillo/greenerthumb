# `fan` Design

`fan` connects STDIN from each of the listed in-programs to STDOUT of all the
listed out-programs without mangling lines.

## Program

```
./fan [--in '<in>'...] [--out '<out>'...]
```

## Example

```
./fan --in 'echo a' --in 'echo b' --out 'cat' --out 'cat'

a
a
b
b
```
