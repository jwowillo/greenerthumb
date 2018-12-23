# `store` Design

`store` messages with unique names.

`read` will read messages from a store and `write` will write unique messages to
a store updating repeats.

`read` will write all messages in the store to STDOUT. `write` will accept all
messages from STDIN and put them in the store. Duplicated messages are detected
based on the name.
