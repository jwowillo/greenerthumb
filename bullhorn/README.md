# `bullhorn`

`bullhorn` allows data to be sent on a network from publishers to subscribers.

An optional reconnect delay will cause subscribers to attempt to reconnect to
the publisher.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make bullhorn` builds `publish` and `subscribe`.
* `make publish` builds `publish`.
* `make subscribe` builds `subscribe`.
* `make test` builds `bullhorn`'s tests.

## Running

```
./publish <port>
```

```
./subscribe <publish_host> <publish_port> ?--reconnect-delay <delay>
```

An example is:

Machine 1 (192.168.1.50):

```
./publish 5050
```

Machine 2 (192.168.1.80):

```
./subscribe 192.168.1.50 5050
```

Machine 3 (192.168.1.81):

```
./subscribe 192.168.1.50 5050
```

Machine 1 (192.168.1.50):

```
< message1
< message2
```

Machine 2 (192.168.1.80):

```
message1
message2
```

Machine 3 (192.168.1.81):

```
message1
message2
```

## Testing

```
cd build && ./test
```
