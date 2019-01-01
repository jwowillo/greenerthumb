package io.greenerthumb.receiver;

/**
 * ReceiveHandler handles a received T.
 *
 * @param <T> Type to handle.
 */
@FunctionalInterface
public interface ReceiveHandler<T> {
    void receive(T message);
}
