# `filter`

`filter` instances of data-types by specifying a list of ANDing conditions in
the set of less than or equal to, less than, equal, greater than, and greater
than or equal to and filtering STDIN according to the conditions.

An epsilon value for comparisons can also optionally be passed. The system
epsilon is used otherwise.

## Building

* `make` builds all targets.
* `make filter` builds `filter`.
* `make test` builds `filter`'s tests.

## Running

```
./filter <name> <key> \
    ?--epsilon <epsilon> \
    ?--e <equal> \
    ?--lt <less_than> \
    ?--lte <less_than_or_equal_to> \
    ?--gt <greater_than> \
    ?--gte <greater_than_or_equal_to>
```

An example is:

```
./filter A 1 --lt 4 --gt 2

< {"Header": {}, "1": 1}
< {"Header": {}, "1": 2}

< {"Header": {}, "1": 3}
{"Header": {}, "1": 3}

< {"Header": {}, "1": 4}
< {"Header": {}, "1": 5}
```

## Testing

```
cd build && ./test
```
