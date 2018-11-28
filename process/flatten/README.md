# `flatten`

`flatten` smooths data by keeping a sliding window of 3 instances of a data-type
and replacing it with a weighted average of the 3 instances biased towards the
middle instance.

## Building

* `make` builds all targets.
* `make flatten` builds `flatten`.
* `make test` builds `flatten`'s tests.

## Running

```
./flatten

< {"Name": "A", "Timestamp": 0, "1": 1, "2": 7}
< {"Name": "A", "Timestamp": 1, "1": 2, "2": 3}

{"Name": "A", "Timestamp": 0, "1": 1.16667}
{"Name": "A", "Timestamp": 0, "2": 6.33334}

< {"Name": "B", "Timestamp": 0, "3": 4}
< {"Name": "A", "Timestamp": 2, "2": 5}

{"Name": "B", "Timestamp": 0, "3": 4}
{"Name": "A", "Timestamp": 1, "1": 1.83333}
{"Name": "A", "Timestamp": 1, "2": 4}
{"Name": "A", "Timestamp": 2, "2": 4.66667}
```

## Testing

```
cd build && ./test
```
