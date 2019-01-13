# `message` Design

`message` is where `greenerthumb` messages from the ICD are defined, the
bytes-JSON conversion is implemented, and header wrapping is provided.
Converters are provided for both directions. All programs run for continuous
input instead of one message at a time so it is easier to pipe with other
programs. Errors in input are ignored with a log so the converter can continue
running. Bytes are handled in base-16 to facilitate ease of reading and writing
messages between programs.

The "Sender" to greenerthumb-bullhorn-header can't be longer than 255
characters.
