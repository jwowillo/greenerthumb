package io.greenerthumb.receiver;

import io.greenerthumb.convert.Converter;

/**
 * ConvertingReceiver converts a Receiver<U> to a Receiver<T> using a
 * Converter<T, U>.
 *
 * Conversions that fail are skipped.
 *
 * @param <T> Type to convert from.
 * @param <U> Type to convert to.
 */
public class ConvertingReceiver<T, U> implements Receiver<U> {
    private final Receiver<T> receiver;
    private final Converter<T, U> converter;

    /**
     * @param receiver to convert.
     * @param converter to convert with.
     */
    public ConvertingReceiver(
            Receiver<T> receiver,
            Converter<T, U> converter) {
        this.receiver = receiver;
        this.converter = converter;
    }

    @Override
    public void addReceiveHandler(ReceiveHandler<U> handler) {
        receiver.addReceiveHandler(
                t -> converter.convert(t).ifPresent(handler::receive));
    }
}
