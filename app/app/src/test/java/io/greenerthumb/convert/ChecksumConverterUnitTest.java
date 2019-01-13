package io.greenerthumb.convert;

import org.junit.Test;

import java.util.Optional;

import io.greenerthumb.collection.ArrayView;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;
import static org.junit.Assert.assertEquals;

public class ChecksumConverterUnitTest {
    @Test
    public void bytesTooShort() {
        Converter<ArrayView<Byte>, ArrayView<Byte>> deserializer = new ChecksumConverter();
        Optional<ArrayView<Byte>> result = deserializer.convert(new ArrayView<>(new Byte[]{
                0x00, 0x01, 0x02}));
        assertFalse(result.isPresent());
    }

    @Test
    public void checksumNoMessage() {
        Converter<ArrayView<Byte>, ArrayView<Byte>> deserializer = new ChecksumConverter();
        Optional<ArrayView<Byte>> result = deserializer.convert(new ArrayView<>(new Byte[]{
                0x00, 0x00, 0x00, 0x00}));
        assertTrue(result.isPresent());
        assertEquals(0, result.get().size());
    }

    @Test
    public void goodChecksum() {
        Converter<ArrayView<Byte>, ArrayView<Byte>> deserializer = new ChecksumConverter();
        Optional<ArrayView<Byte>> result = deserializer.convert(new ArrayView<>(new Byte[]{
                0x00, 0x00, 0x00, 0x09, 0x01, 0x03, 0x05}));
        assertTrue(result.isPresent());
        assertEquals(3, result.get().size());
        assertEquals(0x01, (byte)result.get().at(0));
        assertEquals(0x03, (byte)result.get().at(1));
        assertEquals(0x05, (byte)result.get().at(2));
    }

    @Test
    public void badChecksum() {
        Converter<ArrayView<Byte>, ArrayView<Byte>> deserializer = new ChecksumConverter();
        Optional<ArrayView<Byte>> result = deserializer.convert(new ArrayView<>(new Byte[]{
                0x00, 0x00, 0x00, 0x08, 0x01, 0x03, 0x05}));
        assertFalse(result.isPresent());
    }
}
