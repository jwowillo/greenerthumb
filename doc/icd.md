# `greenerthumb` ICD

* Timestamps are 8-byte unix timestamps.
* Checksums are cyclic sums of all of a message's bytes excluding the checksum.
* Multi-byte fields are big-endian.
* All messages are newline-terminated across the network for simplicity.

## Air Status Message (pub/sub)

| Byte | Name        | Type  |
| ---- | ----------- | ----- |
| 1    | ID (0x00)   | Byte  |
| 2    | Timestamp   | Long  |
| 10   | Temperature | Float |
| 14   | Checksum    | Byte  |

* Temperature is in degrees fahrenheit.

## Soil Status Message (pub/sub)

| Byte | Name      | Type  |
| ---- | --------- | ----- |
| 1    | ID (0x01) | Byte  |
| 2    | Timestamp | Long  |
| 10   | Moisture  | Float |
| 14   | Checksum  | Byte  |

* Moisture is the ratio of water to soil.

## Disclosure Message (Broadcast on port 35053 by default)

| Byte                   | Name                    | Type          |
| ---------------------- | ----------------------- | ------------- |
| 1                      | ID (0x02)               | Byte          |
| 2                      | Timestamp               | Long          |
| 10                     | Device Name Length (l)  | Byte          |
| 11                     | Device Name             | Byte Sequence |
| 11 + l                 | Publish Host Length (m) | Byte          |
| 11 + l + 1             | Publish Host            | Byte Sequence |
| 11 + l + 1 + m         | Command Host Length (n) | Byte          |
| 11 + l + 1 + m + 1     | Command Host            | Byte Sequence |
| 11 + l + 1 + m + 1 + n | Checksum                | Byte          |

* Device name is the name of the device.
* Publish host is the host the device publishes to.
* Command host is the host the device receives commands from.

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
