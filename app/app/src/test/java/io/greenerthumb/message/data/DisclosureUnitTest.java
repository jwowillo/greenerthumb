package io.greenerthumb.message.data;

import org.junit.Test;

import static org.junit.Assert.*;

public class DisclosureUnitTest {
    @Test
    public void deviceNameIsInjected() {
        Disclosure disclosure = new Disclosure("a", "b", "c");
        assertEquals(disclosure.deviceName(), "a");
    }

    @Test
    public void publishHostIsInjected() {
        Disclosure disclosure = new Disclosure("a", "b", "c");
        assertEquals(disclosure.publishHost(), "b");
    }

    @Test
    public void commandHostIsInjected() {
        Disclosure disclosure = new Disclosure("a", "b", "c");
        assertEquals(disclosure.commandHost(), "c");
    }

    @Test
    public void equalsIsSame() {
        Disclosure a = new Disclosure("a", "b", "c");
        Disclosure b = new Disclosure("a", "b", "c");
        assertEquals(a, b);
        assertEquals(b, a);
    }

    @Test
    public void equalsIsNotSame() {
        Disclosure a = new Disclosure("a", "b", "c");
        Disclosure b = new Disclosure("1", "2", "3");
        assertNotEquals(a, b);
        assertNotEquals(b, a);
    }

    @Test
    public void hashCodeIsSame() {
        Disclosure a = new Disclosure("a", "b", "c");
        Disclosure b = new Disclosure("a", "b", "c");
        assertEquals(a.hashCode(), b.hashCode());
    }

    @Test
    public void hashCodeIsNotSame() {
        Disclosure a = new Disclosure("a", "b", "c");
        Disclosure b = new Disclosure("1", "2", "3");
        assertNotEquals(a.hashCode(), b.hashCode());
    }
}