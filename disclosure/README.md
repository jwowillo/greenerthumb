# `disclosure`

`disclosure` prints the disclosure message with the passed values periodically
at the given rate in hertz.

The default rate is 5 hertz.

## Building

* `make` builds all targets.
* `make disclosure` builds `disclosure`.
* `make test` builds `disclosure`'s tests.

## Running

```
./disclosure <device_name> <publish_host> <command_host> ?--rate <rate>
```

An example after one second is:

```
./disclosure device :8080 :8081 --rate 1

{"Name":"disclosure","Timestamp":0,"DeviceName":"device","PublishHost":":8080","CommandHost":":8081"}
```

## Testing

```
cd build && ./test
```
