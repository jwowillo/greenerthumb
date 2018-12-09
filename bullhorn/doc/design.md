# `bullhorn` Design

## Network

![Sequence Diagram](sequence.png)

## Programs

### Publisher

`./publish <port>`

Runs `publish` which publishes newline-delimited messages from STDIN to
subscribers received on `<port>`.

`publish` exits when STDIN is closed.

### Subscriber

```
./subscribe <publish_host> <publish_port> ?--reconnect-delay <delay>
```

Runs `subscribe` which prints messages to STDOUT published from `publish`
running at the `<publish_host>` and `<publish_port>`.

An optional reconnect delay will cause suscribers to attempt to reconnect to the
publisher.

## Examples

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
