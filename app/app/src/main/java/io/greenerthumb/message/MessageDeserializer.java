package io.greenerthumb.message;

import java.time.Instant;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.Optional;

import io.greenerthumb.collection.ArrayView;
import io.greenerthumb.convert.Converter;

/**
 * MessageDeserializer tries to parse Messages from ArrayView<Byte>s.
 *
 * An empty Optional is returned if the Message isn't a Disclosure.
 */
public class MessageDeserializer implements Converter<ArrayView<Byte>, Message> {
    @Override
    public Optional<Message> convert(ArrayView<Byte> data) {
        if (data.size() < 1) {
            return Optional.empty();
        }

        byte actual = sumOf(data.viewOf(0, data.size()-1));
        byte expected = data.at(data.size()-1);
        if (actual != expected) {
            return Optional.empty(); // Bad checksum.
        }
        data = data.viewOf(0, data.size()-1);

        if (data.size() < 1) {
            return Optional.empty();
        }
        Optional<MessageType> type = typeFor(data.at(0));
        if (!type.isPresent()) {
            return Optional.empty();
        }
        data = ArrayView.advance(data, 1);

        if (data.size() < 8) {
            return Optional.empty();
        }
        OffsetDateTime time = OffsetDateTime.ofInstant(
                Instant.ofEpochSecond(longFrom(data)),
                ZoneOffset.UTC);
        data = ArrayView.advance(data, 8);

        Optional<String> sender = parseString(data);
        if (!sender.isPresent()) {
            return Optional.empty();
        }
        data = ArrayView.advance(data, sender.get().length()+1);

        return Optional.of(new Message(
                type.get(),
                time,
                sender.get(),
                data));
    }

    private static long longFrom(ArrayView<Byte> data) {
        return data.at(7) |
                data.at(6) << 8 |
                data.at(5) << 16 |
                data.at(4) << 24 |
                data.at(3) << 32 |
                data.at(3) << 40 |
                data.at(2) << 48 |
                data.at(1) << 56 |
                data.at(0) << 64;
    }

    private static Optional<MessageType> typeFor(byte x) {
        switch (x) {
            case 0x02:
                return Optional.of(MessageType.DISCLOSURE);
            default:
                return Optional.empty();
        }
    }

    private static byte sumOf(ArrayView<Byte> data) {
        int sum = 0;
        for (int i = 0; i < data.size(); i++) {
            int value = data.at(i);
            if (value < 0) {
                value += 0xff + 1; // Java uses signed bytes.
            }
            sum += value;
        }
        return (byte)sum;
    }

    /**
     * parseString parses a String from the ArrayView<Byte> by checking for a length at the first
     * index and then reading the length in the next part of the ArrayView<Byte> into a String.
     *
     * @param view containing the String.
     * @return The parsed String or an empty Optional if no String could be parsed.
     */
    private static Optional<String> parseString(ArrayView<Byte> view) {
        if (view.size() < 1) {
            return Optional.empty();
        }
        byte length = view.at(0);
        if (view.size() < 1 + length) {
            return Optional.empty();
        }
        return Optional.of(new String(primitive(view.viewOf(1, length))));
    }

    private static byte[] primitive(ArrayView<Byte> view) {
        byte[] primitive  = new byte[view.size()];
        for (int i = 0; i < view.size(); i++) {
            primitive[i] = view.at(i);
        }
        return primitive;
    }
}
