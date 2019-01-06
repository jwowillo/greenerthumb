# `clean`

`clean` reads all input until STDIN is closed and filters instances that are
more than a passed number of standard deviations away from the mean.

## Building

* `make` builds all targets.
* `make clean` builds `clean`.
* `make test` builds `clean`'s tests.

## Running

```
./clean <standard_deviation_limit>
```

An example is:

```
./clean 1

< {"Header": {"Name": "A"}, "1": 1}
< {"Header": {"Name": "A"}, "1": 2}
< {"Header": {"Name": "A"}, "1": 3}
< {"Header": {"Name": "A"}, "1": 4}
< {"Header": {"Name": "A"}, "1": 5}

{"Header": {"Name": "A"}, "1": 2}
{"Header": {"Name": "A"}, "1": 3}
{"Header": {"Name": "A"}, "1": 4}
```

## Testing

```
cd build && ./test
```
