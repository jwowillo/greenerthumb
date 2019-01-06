package io.greenerthumb.message.data;

import org.junit.Test;

import java.time.OffsetDateTime;
import java.util.Optional;

import io.greenerthumb.collection.ArrayView;
import io.greenerthumb.convert.Converter;
import io.greenerthumb.message.Message;
import io.greenerthumb.message.MessageType;

import static org.junit.Assert.*;

public class DisclosureDeserializerUnitTest {
    @Test
    public void deserialize() {
        Converter<Message, Disclosure> deserializer = new DisclosureDeserializer();
        ArrayView<Byte> data = new ArrayView<>(new Byte[]{
                0x05,
                0x3a, 0x38, 0x30, 0x38, 0x30});
        Message message = new Message(
                MessageType.DISCLOSURE,
                OffsetDateTime.now(),
                "sender",
                data);
        Optional<Disclosure> disclosure = deserializer.convert(message);
        assertTrue(disclosure.isPresent());
        assertEquals(":8080", disclosure.get().host());
    }

    @Test
    public void deserializeBad() {
        Converter<Message, Disclosure> deserializer = new DisclosureDeserializer();
        ArrayView<Byte> data = new ArrayView<>(new Byte[]{
                0x05, // Bad length
                0x3a, 0x38, 0x30, 0x38});
        Message message = new Message(
                MessageType.DISCLOSURE,
                OffsetDateTime.now(),
                "sender",
                data);
        Optional<Disclosure> disclosure = deserializer.convert(message);
        assertFalse(disclosure.isPresent());
    }
}
