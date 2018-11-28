# `process` Design

`process` `greenerthumb` data.

## Programs

All programs accept `greenerthumb` ICD JSON messages from STDIN and report
results to STDOUT. Each program terminates with a message printed to STDERR if
any JSON message is malformed.

### `summarize`

`summarize` reads all input until STDIN is closed and then reports a
5-number-summary for each data-type along with how many instances of that
data-type were included.

### `flatten`

`flatten` smooths data by keeping a sliding window of 3 instances of a data-type
and replacing it with a weighted average of the 3 instances biased towards the
middle instance. The first instance and last instance have a copy of themselves
used as the instance to the left and right of them.

The left and right values are weighted by 1/6 each while the middle value is
weighted 2/3.

### `filter`

`fitler` instances of data-types by specifying a list of ANDing conditions in
the set of less than or equal to, less than, equal, greater than, and greater
than or equal to as a comma-separated list of `<NAME,KEY,VALUE>` and filtering
STDIN according to the conditions.

An epsilon value for comparisons can also optionally be passed. The system
epsilon should be used otherwise.

### `clean`

`clean` reads all input until STDIN is closed and filters instances that are
more than a passed number of standard deviations away from the mean.

## Examples

### `summarize`

```
./summarize

< {"Name": "A", "Timestamp": 0, "1": 1}
< {"Name": "A", "Timestamp": 1, "1": 2}
< {"Name": "A", "Timestamp": 2, "1": 3}
< {"Name": "A", "Timestamp": 3, "1": 4}
< {"Name": "A", "Timestamp": 4, "1": 5}

{"A": {"1": {"N": 5, "Minimum": 1, "Q1": 1.5, "Median": 3, "Q2": 4.5, "Maximum": 5}}}
```

### `flatten`

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

### `filter`

```
./filter A 1 --lt 4 --gt 2

< {"Name": "A", "Timestamp": 0, "1": 1}
< {"Name": "A", "Timestamp": 1, "1": 2}

< {"Name": "A", "Timestamp": 2, "1": 3}
{"Name": "A", "Timestamp": 2, "1": 3}

< {"Name": "A", "Timestamp": 3, "1": 4}
< {"Name": "A", "Timestamp": 4, "1": 5}
```

```
./filter A 1 --e 3 --epsilon 0.01


< {"Name": "A", "Timestamp": 2, "1": 2.991}
{"Name": "A", "Timestamp": 2, "1": 2.991}
< {"Name": "A", "Timestamp": 2, "1": 2.99}
```

### `clean`

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
