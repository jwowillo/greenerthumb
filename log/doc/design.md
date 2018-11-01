# `log` Design

`log`s JSON messages from STDIN to a file.

## Program

`log` appends JSON messages to a specified file.

## Example

```
echo 'line' | ./log log.txt
```

```
cat log.txt

line
```
