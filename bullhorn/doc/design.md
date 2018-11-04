# `bullhorn` ICD

## Network

![Sequence Diagram](sequence.png)

## Program

### Publisher

`./publish <port>`

Runs `publish` which publishes newline-delimited messages from `STDIN` to
subscribers received on `<port>`.

### Subscriber

```
./subscribe <port> <publish_host> <publish_port>
```

Runs `subscribe` which prints messages to `STDOUT` published from `publish`
running at the `<publish_host>` and `<publish_port>`.

An option is also provided to attempt to automatically reconnect to
`bullhorn/publish` after a disconnect.

## Example

Machine 1 (192.168.1.50):

```
./publish 5050
```

Machine 2 (192.168.1.80):

```
./subscribe 8080 192.168.1.50 5050
```

Machine 3 (192.168.1.81):

```
./subscribe 8081 192.168.1.50 5050
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
