# `greenerthumb` ICD

* Timestamps are 8-byte unix timestamps.
* Sender is the name of the package who sent the message. Multiple hosts can
  share a sender. This means the hosts are in the same package.
* Checksums are cyclic sums of all of a message's bytes excluding the checksum.
* Multi-byte fields are big-endian.
* All messages start with the header.

## Links

Provided links are:
* `broadcast`
* `listen`
* `pub/sub`

These are described more in `bullhorn`. Each link associates metadata with
messages while over the network. `broadcast` and `pub/sub` prefix messages with
a 4 byte checksum. `listen` prefixes messages with a 2 byte message-length. This
are stripped after being checked at either end of the link.

## Header

| Byte       | Name              | Type           |
| ---------- | ----------------- | -------------- |
| 1          | ID                | Byte           |
| 2          | Timestamp         | Long           |
| 10         | Sender Length (n) | Byte           |
| 11         | Sender            | Byte Sequence  |

## Air Status Message (pub/sub, ID = 0x00)

| Byte | Name        | Type  |
| ---- | ----------- | ----- |
| 1    | Temperature | Float |

* Temperature is in degrees fahrenheit.

## Soil Status Message (pub/sub, ID = 0x01)

| Byte | Name      | Type  |
| ---- | --------- | ----- |
| 1    | Moisture  | Float |

* Moisture is the ratio of water to soil.

## Disclosure Message (Broadcast on port 35053 by default, ID = 0x02)

| Byte   | Name            | Type          |
| ------ | --------------- | ------------- |
| 1      | Host Length (n) | Byte          |
| 2      | Host            | Byte Sequence |

* Host is the host the device publishes to.

## JSON

JSON is used for internal messaging. Every message has a corresponding JSON
format with the structure:

```
{
  "Header": {
    "Name": <message_name>,
    "Timestamp": <timestamp>,
    "Sender": <sender>
  },
  <name>: <value>,...
}
```

IDs are swapped with names for friendlier use in applications. Names and values
correspond to the non-ID, non-time, and non-checksum fields in the messages.
