# `bullhorn` ICD

## Network

![Sequence Diagram](sequence.png)

## Program

### Server

`./server --port <port>`

Runs `server` which publishes newline-delimited messages from `STDIN` to
subscribers received on `<port>`.

### Client

```
./bclient --port <port> \
    --remote-host <remote_host> --remote-port <remote_port>
```

Runs `client` which prints messages to `STDOUT` published from the `server`
running at the `<remote_host>` and `<remote_port>`.

### Example

Machine 1 (192.168.1.50):

```
./server --port 5050
```

Machine 2 (192.168.1.80):

```
./client --port 8080 --remote-host 192.168.1.50 --remote-port 5050
```

Machine 3 (192.168.1.81):

```
./client --port 8081 --remote-host 192.168.1.50 --remote-port 5050
```

Machine 1 (192.168.1.50):

```
< message1
< message2
```

Machine 2 (192.168.1.80):

```
> message1
> message2
```

Machine 3 (192.168.1.81):

```
> message1
> message2
```
