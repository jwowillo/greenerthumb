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

