# `message` Design

`message` is where `greenerthumb` messages from the ICD are defined and the
bytes-JSON conversion is implemented. Converters are provided for both
directions. The converters run for continuous input instead of one message at a
time so it is easier to pipe with other programs. Errors in input are ignored
with a log so the converter can continue running.
