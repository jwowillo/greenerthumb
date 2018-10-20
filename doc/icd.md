# `greenerthumb` ICD

* Checksums are cyclic sums of all of a message's bytes excluding the checksum.
* Multi-byte fields are big-endian.

## Air Status Message

| Byte | Name        | Type  |
| ---- | ----------- | ----- |
| 1    | ID (0x00)   | Byte  |
| 2    | Temperature | Float |
| 6    | Humidity    | Float |
| 10   | Checksum    | Byte  |

* Temperature is in degrees fahrenheit.
* Humidity is the ratio of water to air.

## Soil Moisture Status Message

| Byte | Name      | Type   |
| ---- | --------- | ------ |
| 1    | ID (0x01) | Byte   |
| 2    | Moisture  | Float  |
| 6    | Checksum  | Byte   |

* Moisture is the ratio of water to soil.

## JSON

JSON is used for internal messaging. Every message has a corresponding JSON
format with the structure:

```
{
  "name": <message_name>,
  <name>: <value>,...
}
```

IDs are swapped with names for friendlier use in applications. Checksums are
excluded since the messages don't need to be sent over a network. Names and
values correspond to the non-ID and non-checksum fields in the messages.
