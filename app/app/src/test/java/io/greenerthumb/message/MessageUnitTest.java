package io.greenerthumb.message;

import org.junit.Test;

import java.time.OffsetDateTime;

import io.greenerthumb.collection.ArrayView;

import static org.junit.Assert.assertEquals;

public class MessageUnitTest {
    @Test
    public void messageType() {
        MessageType expected = MessageType.DISCLOSURE;
        Message message = new Message(
                expected,
                OffsetDateTime.now(),
                "sender",
                new ArrayView<>(new Byte[]{}));
        assertEquals(expected, message.type());
    }

    @Test
    public void timestamp() {
        OffsetDateTime expected = OffsetDateTime.now();
        Message message = new Message(
                MessageType.DISCLOSURE,
                expected,
                "sender",
                new ArrayView<>(new Byte[]{}));
        assertEquals(expected, message.timestamp());
    }

    @Test
    public void sender() {
        String expected = "sender";
        Message message = new Message(
                MessageType.DISCLOSURE,
                OffsetDateTime.now(),
                expected,
                new ArrayView<>(new Byte[]{}));
        assertEquals(expected, message.sender());
    }

    @Test
    public void data() {
        ArrayView<Byte> expected = new ArrayView<>(new Byte[]{1, 2, 3});
        Message message = new Message(
                MessageType.DISCLOSURE,
                OffsetDateTime.now(),
                "sender",
                expected);
        ArrayView<Byte> actual = message.data();
        assertEquals(actual.size(), 3);
        assertEquals(expected.at(0), actual.at(0));
        assertEquals(expected.at(1), actual.at(1));
        assertEquals(expected.at(2), actual.at(2));
    }
}
