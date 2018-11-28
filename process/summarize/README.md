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
,/summarize

< {"Name": "A", "Timestamp": 0, "1": 1}
< {"Name": "A", "Timestamp": 1, "1": 2}
< {"Name": "A", "Timestamp": 2, "1": 3}
< {"Name": "A", "Timestamp": 3, "1": 4}
< {"Name": "A", "Timestamp": 4, "1": 5}

{"A": {"1": {"N": 5, "Minimum": 1, "Q1": 1.5, "Median": 3, "Q2": 4.5, "Maximum": 5}}}
```

## Testing

```
cd build && ./test
```
