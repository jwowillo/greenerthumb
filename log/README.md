# `log`

`log` messages from STDIN to a file.

An optional duration flag can be specified which sets the duration a log file is
used before being rotated.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make log` builds `log`.

## Running

```
./log ?--duration <duration>
```

An example is:

```
echo 'line' | ./log
```

```
cat log-1543537416.log

line
```
