# `clean`

`clean` reads all input until STDIN is closed and filters instances that are
more than a passed number of standard deviations away from the mean.

## Building

* `make` builds all targets.
* `make clean` builds `clean`.
* `make test` builds `clean`'s tests.

## Running

```
./clean 1

< {"Name": "A", "Timestamp": 0, "1": 1}
< {"Name": "A", "Timestamp": 1, "1": 2}
< {"Name": "A", "Timestamp": 2, "1": 3}
< {"Name": "A", "Timestamp": 3, "1": 4}
< {"Name": "A", "Timestamp": 4, "1": 5}

{"Name": "A", "Timestamp": 1, "1": 2}
{"Name": "A", "Timestamp": 2, "1": 3}
{"Name": "A", "Timestamp": 3, "1": 4}
```

## Testing

```
cd build && ./test
```
