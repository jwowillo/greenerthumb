package io.greenerthumb.message;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class MessageTypeUnitTest {
    @Test
    public void disclosureIdIs0x02() {
        assertEquals(0x02, MessageType.DISCLOSURE.id());
    }
}
