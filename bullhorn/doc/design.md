# `bullhorn` Design

`bullhorn` contains program pairs for networked communication. Included methods
are pub/sub, broadcast, and reliable.

## pub/sub

pub/sub allows messages to be sent from publishers to subscribers.

![Sequence Diagram](sequence.png)

This is done via a pub/sub system which operates over UDP with a TCP trunk. The
TCP trunk allows the publisher to know when to stop publishing to subscribers
and lets the subscriber know when it needs to try to reconnect to a publisher if
reconnect is enabled. The unreliable UDP connection is fine because mostly
periodic statuses are sent through the system.

The publisher will publish all newline-separated lines it receives over STDIN to
every subscriber until STDIN is closed.

The subscribers print all newline-separated lines they receive from the
publisher until the publisher is closed if reconnect isn't enabled. Subscribers
never close if reconnect is enabled and will just periodically attempt
reconnects. Subscribers will always exit with a failure to connect
unless terminated because they will either try to reconnect forever or fail to
connect to a terminated publisher.

## broadcast

broadcast messages to all clients.

This is done via the broadcast address.

The server will send all newline-separated lines it receives over STDIN to every
client until STDIN is closed. The clients print all newline-separated lines they
receive until they are terminated.

## Reliable

Reliable allows messages to be reliably sent from talkers to a listener.

This is done via TCP.

The clients will connect to the servers, write all their input from STDIN to the
connections, and the server will echo the messages to STDOUT. Clients run until
STDIN is closed or the connection is closed. The servers run until they're
terminated.
