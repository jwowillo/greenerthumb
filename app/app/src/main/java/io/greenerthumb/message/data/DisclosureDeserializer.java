package io.greenerthumb.message.data;

import java.util.Optional;

import io.greenerthumb.collection.ArrayView;
import io.greenerthumb.convert.Converter;
import io.greenerthumb.message.Message;
import io.greenerthumb.message.MessageType;

/**
 * DisclosureDeserializer tries to parse Disclosures from Messages.
 *
 * An empty Optional is returned if the Message isn't a Disclosure.
 */
public class DisclosureDeserializer implements Converter<Message, Disclosure> {
    @Override
    public Optional<Disclosure> convert(Message message) {
        if (message.type() != MessageType.DISCLOSURE) {
            return Optional.empty();
        }

        ArrayView<Byte> data = message.data();

        Optional<String> host = parseString(data);
        if (!host.isPresent()) {
            return Optional.empty();
        }
        data = ArrayView.advance(data,1 + host.get().length());

        if (data.size() != 0) {
            return Optional.empty();
        }

        return Optional.of(new Disclosure(host.get()));
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
