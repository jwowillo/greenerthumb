package io.greenerthumb.receiver;

/**
 * Receiver receives Ts.
 *
 * @param <T> Type to receive.
 */
public interface Receiver<T> {
    /**
     * addReceiveHandler adds a ReceiveHandler<T> to be invoked whenever a
     * T is received.
     *
     * @param handler to call with received Ts.
     */
    void addReceiveHandler(ReceiveHandler<T> handler);
}