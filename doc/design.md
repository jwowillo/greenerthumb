# `greenerthumb` Design

## `fan`

Fans connects multiple programs' STDOUTs and STDINs.

## `bullhorn`

Allows data to be sent on a network from publishers to subscribers (3).

### `publish`

Publishes data from STDIN to all subscribers.

### `subscribe`

Receives data from a publisher to STDOUT.

## `sense`

Sensors write messages from the `greenerthumb` ICD to STDOUT. These are fanned
into a `bullhorn/publisher` with `fan/in`. This allows sensors to be seamlessly
excluded (1). Testing is also simplified since emulators that print to STDOUT
can replace sensors or `bullhorn/subscriber`.

### `air`

Senses the 'Air Status Message' (2a).

### `soil`

Senses the 'Soil Moisture Status Message' (2b).

## `emulate`

Emulators allow simple testing the by excluding the network and sensors.

### `air`

Emulates the 'Air Status Message'.

### `soil`

Emulates the 'Soil Moisture Status Message'.

## `log`

Logs data from STDIN to a file. This can be used with `bullhorn/subscriber` and
`fan/out` to log sensed data (4).

## `plot`

### `replay`

Plots ICD messages stored in a log file (5).
