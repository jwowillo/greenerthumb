# `bullhorn`

`bullhorn` contains program pairs for networked communication. Included methods
are pub/sub, broadcast, and reliable.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make bullhorn` builds all `bullhorn programs.
* `make pubsub` builds `publish` and `subscribe`.
* `make publish` builds `publish`.
* `make subscribe` builds `subscribe`.
* `make broadcast` builds `yell` and `snoop`.
* `make yell` builds `yell`.
* `make snoop` builds `snoop`.
* `make reliable` builds `talk` and `listen`.
* `make talk` builds `talk`.
* `make listen` builds `listen`.
* `make test` builds `bullhorn`'s tests.

## pub/sub

pub/sub allows messages to be sent from publishers to subscribers.

An optional reconnect delay will cause subscribers to attempt to reconnect to
the publisher.

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

## Broadcast

broadcast messages to all clients.

```
./yell <port>
```

```
./snoop <port>
```

An example is:

Machine 1 (192.168.1.50):

```
./yell 5050
```

Machine 2 (192.168.1.80):

```
./snoop 5050
```

Machine 3 (192.168.1.81):

```
./snoop 5050
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

## Reliable

Reliable allows messages to be reliably sent from talkers to a listener.

```
./listen ?--host <host>
```

```
./talk <host>
```

An example is:

Machine 1 (192.168.1.50):

```
./listen
```

Machine 2 (192.168.1.80):

```
./talk

< a
< b
```

Machine 1 (192.168.1.50):

```
a
b
```

## Testing

```
cd build && ./test
```
