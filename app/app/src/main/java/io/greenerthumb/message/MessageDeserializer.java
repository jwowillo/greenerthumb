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
        long byte0 = data.at(0);
        long byte1 = data.at(1);
        long byte2 = data.at(2);
        long byte3 = data.at(3);
        long byte4 = data.at(4);
        long byte5 = data.at(5);
        long byte6 = data.at(6);
        long byte7 = data.at(7);
        if (byte0 < 0) {
            byte0 += 0xff + 1;
        }
        if (byte1 < 0) {
            byte1 += 0xff + 1;
        }
        if (byte2 < 0) {
            byte2 += 0xff + 1;
        }
        if (byte3 < 0) {
            byte3 += 0xff + 1;
        }
        if (byte4 < 0) {
            byte4 += 0xff + 1;
        }
        if (byte5 < 0) {
            byte5 += 0xff + 1;
        }
        if (byte6 < 0) {
            byte6 += 0xff + 1;
        }
        if (byte7 < 0) {
            byte7 += 0xff + 1;
        }
        return byte7 |
                byte6 << 8 |
                byte5 << 16 |
                byte4 << 24 |
                byte3 << 32 |
                byte2 << 40 |
                byte1 << 48 |
                byte0 << 56;
    }

    private static Optional<MessageType> typeFor(byte x) {
        switch (x) {
            case 0x02:
                return Optional.of(MessageType.DISCLOSURE);
            default:
                return Optional.empty();
        }
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
