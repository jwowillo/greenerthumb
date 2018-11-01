# `log`

`log`s JSON messages from STDIN to a file.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make log` builds `log`.

## Running

```
echo 'line' | ./log log.txt
```

```
cat log.txt

line
```
