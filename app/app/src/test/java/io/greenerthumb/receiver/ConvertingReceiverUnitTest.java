package io.greenerthumb.receiver;

import org.junit.Test;

import java.util.Optional;
import java.util.concurrent.atomic.AtomicBoolean;

import io.greenerthumb.mock.MockReceiver;

import static org.junit.Assert.*;


public class ConvertingReceiverUnitTest {
    @Test
    public void convertsValid() {
        AtomicBoolean wasCalled = new AtomicBoolean(false);
        MockReceiver<Integer> rawReceiver = new MockReceiver<>();
        Receiver<Integer> receiver = new ConvertingReceiver<>(rawReceiver, this::convert);
        Integer initial = 1;
        receiver.addReceiveHandler(x -> {
            assertEquals(x, new Integer(initial + 1));
            wasCalled.set(true);
        });
        rawReceiver.receive(initial);
        assertTrue(wasCalled.get());
    }

    @Test
    public void doesNotConvertInvalid() {
        AtomicBoolean wasCalled = new AtomicBoolean(false);
        MockReceiver<Integer> rawReceiver = new MockReceiver<>();
        Receiver<Integer> receiver = new ConvertingReceiver<>(rawReceiver, this::convert);
        receiver.addReceiveHandler(x -> wasCalled.set(true));
        rawReceiver.receive(0);
        assertFalse(wasCalled.get());
    }

    /**
     * convert is a Converter<int, int> helpful for tests.
     *
     * @param x to convert.
     * @return x++ if x != 0 and an empty Optional otherwise.
     */
    private Optional<Integer> convert(Integer x) {
        if (x == 0) {
           return Optional.empty();
        }
        return Optional.of(x + 1);
    }
}
