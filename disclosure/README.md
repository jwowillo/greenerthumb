# `disclosure`

`disclosure` prints the disclosure message body with the passed values
periodically at the given rate in hertz.

The host can't be longer than 255 characters.

The default rate is 5 hertz.

## Building

* `make` builds all targets.
* `make disclosure` builds `disclosure`.
* `make test` builds `disclosure`'s tests.

## Running

```
./disclosure <host> ?--rate <rate>
```

An example after one second is:

```
./disclosure :8080 --rate 1

{"Host":":8080"}
```

## Testing

```
cd build && ./test
```
