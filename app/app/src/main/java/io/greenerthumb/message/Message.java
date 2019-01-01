package io.greenerthumb.message;

import java.time.OffsetDateTime;

import io.greenerthumb.collection.ArrayView;

/**
 * Message containing its wrapper data and internal data.
 */
public class Message {
    private final MessageType type;
    private final OffsetDateTime timestamp;
    private final ArrayView<Byte> data;

    public Message(MessageType type, OffsetDateTime timestamp, ArrayView<Byte> data) {
        this.type = type;
        this.timestamp = timestamp;
        this.data = data;
    }

    public MessageType type() {
        return type;
    }

    OffsetDateTime timestamp() {
        return timestamp;
    }

    public ArrayView<Byte> data() {
        return data;
    }
}
