package io.greenerthumb.message;

import org.junit.Test;

import java.time.Instant;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.Optional;

import io.greenerthumb.collection.ArrayView;
import io.greenerthumb.convert.Converter;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

public class MessageDeserializerUnitTest {
    @Test
    public void deserialize() {
        Converter<ArrayView<Byte>, Message> deserializer = new MessageDeserializer();
        ArrayView<Byte> data = new ArrayView<>(new Byte[]{
                0x02,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x02});
        Optional<Message> message = deserializer.convert(data);
        assertTrue(message.isPresent());
        assertEquals(MessageType.DISCLOSURE, message.get().type());
        assertEquals(0, message.get().data().size());
        assertEquals(
                OffsetDateTime.ofInstant(Instant.ofEpochSecond(0), ZoneOffset.UTC),
                message.get().timestamp());
    }

    @Test
    public void deserializeBad() {
        Converter<ArrayView<Byte>, Message> deserializer = new MessageDeserializer();
        ArrayView<Byte> data = new ArrayView<>(new Byte[]{
                0x02,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Bad length
                0x02});
        Optional<Message> message = deserializer.convert(data);
        assertFalse(message.isPresent());
    }

    @Test
    public void deserializeBadChecksum() {
        Converter<ArrayView<Byte>, Message> deserializer = new MessageDeserializer();
        ArrayView<Byte> data = new ArrayView<>(new Byte[]{
                0x02,
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x03 // Bad checksum.
        });
        Optional<Message> message = deserializer.convert(data);
        assertFalse(message.isPresent());
    }

    @Test
    public void deserializeBadMessageType() {
        Converter<ArrayView<Byte>, Message> deserializer = new MessageDeserializer();
        ArrayView<Byte> data = new ArrayView<>(new Byte[]{
                (byte)0xFF, // Bad MessageType
                0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
                0x02});
        Optional<Message> message = deserializer.convert(data);
        assertFalse(message.isPresent());
    }
}
