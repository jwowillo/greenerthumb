# `bullhorn`

`bullhorn` contains program pairs for networked communication. Included methods
are `broadcast`, `listen`, and `pubsub`.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make bullhorn` builds all `bullhorn` programs.
* `make broadcast` builds `broadcast` programs.
* `make listen` builds `listen` programs.
* `make pubsub` builds `pubsub` programs.
* `make test` builds `bullhorn`'s tests.

## `broadcast`

`broadcast` messages to all clients.

```
./broadcast/server <port>
```

```
./broadcast/client <port>
```

An example is:

Machine 1 (192.168.1.50):

```
./broadcast/server 5050
```

Machine 2 (192.168.1.80):

```
./broadcast/client 5050
```

Machine 3 (192.168.1.81):

```
./broadcast/client 5050
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

## `listen`

`listen` allows messages to be reliably sent from talkers to a listener.

```
./listen/server ?--host <host>
```

```
./listen/client <host>
```

An example is:

Machine 1 (192.168.1.50):

```
./listen/server --host :5050
```

Machine 2 (192.168.1.80):

```
./listen/client :5050

< a
< b
```

Machine 1 (192.168.1.50):

```
a
b
```

## `pubsub`

`pubsub` allows messages to be sent from publishers to subscribers.

An optional reconnect delay will cause subscribers to attempt to reconnect to
the publisher.

```
./pubsub/server ?--host <host>
```

```
./pubsub/client <publish_host> ?--reconnect-delay <delay>
```

An example is:

Machine 1 (192.168.1.50):

```
./pubsub/server --host :5050
```

Machine 2 (192.168.1.80):

```
./pubsub/client 192.168.1.50:5050
```

Machine 3 (192.168.1.81):

```
./pubsub/client 192.168.1.50:5050
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
