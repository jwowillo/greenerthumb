# `sense` Design

`sense` programs write `greenerthumb` JSON messages from sensors to STDOUT.

These can be `fan`ned into `message/bytes` piped into `bullhorn/publish`.

## Sensors

### `air`

`air` senses the 'Air Status Message' at 0.1 hertz.

![Air Schematic](air.png)

### `soil`

`soil` senses the 'Soil Status Message' at 0.1 hertz.

![Soil Schematic](soil.png)

## Emulators

Emulators are provided for all programs and each accepts an optional rate flag.
