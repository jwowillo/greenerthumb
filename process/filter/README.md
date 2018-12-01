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
./filter A 1 --lt 4 --gt 2

< {"Name": "A", "Timestamp": 0, "1": 1}
< {"Name": "A", "Timestamp": 1, "1": 2}

< {"Name": "A", "Timestamp": 2, "1": 3}
{"Name": "A", "Timestamp": 2, "1": 3}

< {"Name": "A", "Timestamp": 3, "1": 4}
< {"Name": "A", "Timestamp": 4, "1": 5}
```

## Testing

```
cd build && ./test
```
