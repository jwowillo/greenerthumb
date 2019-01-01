package io.greenerthumb.message;

public enum MessageType {
    DISCLOSURE((byte)0x02);

    private final byte id;

    MessageType(byte id) {
        this.id = id;
    }

    public byte id() {
        return this.id;
    }
}
