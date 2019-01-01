package io.greenerthumb.mock;

import java.util.ArrayList;
import java.util.Collection;

import io.greenerthumb.receiver.ReceiveHandler;
import io.greenerthumb.receiver.Receiver;

/**
 * MockReceiver is a Receiver<T> that has Ts received manually.
 */
public class MockReceiver<T> implements Receiver<T> {
    private final Collection<ReceiveHandler<T>> handlers = new ArrayList<>();

    public void receive(T t) {
        for (ReceiveHandler<T> handler : handlers) {
            handler.receive(t);
        }
    }

    @Override
    public void addReceiveHandler(ReceiveHandler<T> handler) {
        handlers.add(handler);
    }
}
