# `greenerthumb` Design

## `fan`

`fan` connects STDOUTs from listed out-programs to STDINs of listed in-programs.

## `message`

`message` is where `greenerthumb` messages from the ICD are defined and the
bytes-JSON conversion is implemented.

### `bytes`

`bytes` converts JSON messages from STDIN to bytes written to STDOUT.

### `json`

`json` converts bytes messages from STDIN to JSON written to STDOUT.

## `bullhorn`

`bullhorn` allows data to be sent on a network from publishers to subscribers
(3).

### `publish`

`publish`es data from STDIN to all subscribers.

### `subscribe`

`subscribe`s to a publisher and write data to STDOUT.

## `sense`

`sense` writes `greenerthumb` JSON messages from sensors to STDOUT. These can be
`fan`ned into `message/bytes` piped into `bullhorn/publish`. This allows sensors
to be excluded (1).

### `air`

`air` senses the 'Air Status Message' (2a).

### `soil`

`soil` senses the 'Soil Moisture Status Message' (2b).

## `log`

`log`s JSON messages from STDIN to a file. This can be used with
`bullhorn/subscribe` piped into `message/json` to log sensed data (4).

## `plot`

`plot`s `greenerthumb` JSON messages from STDIN. This can be used with
`bullhorn/subscribe` piped into `message/json` to plot sensed data (5).
