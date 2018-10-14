# `fan` Design

`fan` connects STDOUTs from listed out-programs to STDINs of listed in-programs.

## Program

```
./fan [--out '<out>'...] [--in '<in>'...]
```

## Example

```
./fan --out 'echo a' --out 'echo b' --in 'cat' --in 'cat'

a
a
b
b
```
