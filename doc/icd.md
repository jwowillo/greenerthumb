# `greenerthumb` ICD

* Timestamps are 8-byte unix timestamps.
* Checksums are cyclic sums of all of a message's bytes excluding the checksum.
* Multi-byte fields are big-endian.

## Air Status Message

| Byte | Name        | Type  |
| ---- | ----------- | ----- |
| 1    | ID (0x00)   | Byte  |
| 2    | Timestamp   | Long  |
| 10   | Temperature | Float |
| 14   | Checksum    | Byte  |

* Temperature is in degrees fahrenheit.

## Soil Status Message

| Byte | Name      | Type   |
| ---- | --------- | ------ |
| 1    | ID (0x01) | Byte   |
| 2    | Timestamp | Long   |
| 10   | Moisture  | Float  |
| 14   | Checksum  | Byte   |

* Moisture is the ratio of water to soil.

## JSON

JSON is used for internal messaging. Every message has a corresponding JSON
format with the structure:

```
{
  "name": <message_name>,
  "timestamp": <timestamp>,
  <name>: <value>,...
}
```

IDs are swapped with names for friendlier use in applications. Checksums are
excluded since the messages don't need to be sent over a network. Names and
values correspond to the non-ID, non-time, and non-checksum fields in the
messages.
