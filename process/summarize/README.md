# `summarize`

`summarize` reads all input until STDIN is closed and then reports a
5-number-summary for each data-type along with how many instances of that
data-type were included.

## Building

* `make` builds all targets.
* `make summarize` builds `summarize`.
* `make test` builds `summaraize`'s tests.

## Running

```
./summarize
```

An example is:

```
./summarize

< {"Header": {"Name": "A"}, "1": 1}
< {"Header": {"Name": "A"}, "1": 2}
< {"Header": {"Name": "A"}, "1": 3}
< {"Header": {"Name": "A"}, "1": 4}
< {"Header": {"Name": "A"}, "1": 5}

{"A": {"1": {"N": 5, "Minimum": 1, "Q1": 1.5, "Median": 3, "Q2": 4.5, "Maximum": 5}}}
```

## Testing

```
cd build && ./test
```
