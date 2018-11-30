# `log` Design

`log` messages from STDIN to a file.

## Program

`log` appends messages to rotated generated file.

An optional duration flag can be specified which sets the duration a log file is
used before being rotated.

## Example

```
echo 'line' | ./log
```

```
cat log-1543537416.log

line
```
