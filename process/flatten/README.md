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
```

An example is:

```
./flatten

< {"Header": {"Name": "A"}, "1": 1, "2": 7}
< {"Header": {"Name": "A"}, "1": 2, "2": 3}

{"Header": {"Name": "A"}, "1": 1.16667}
{"Header": {"Name": "A"}, "2": 6.33334}

< {"Header": {"Name": "B"}, "3": 4}
< {"Header": {"Name": "A"}, "2": 5}

{"Header": {"Name": "B"}, "3": 4}
{"Header": {"Name": "A"}, "1": 1.83333}
{"Header": {"Name": "A"}, "2": 4}
{"Header": {"Name": "A"}, "2": 4.66667}
```

## Testing

```
cd build && ./test
```
