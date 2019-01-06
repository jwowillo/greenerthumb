package io.greenerthumb.message.data;

import org.junit.Test;

import static org.junit.Assert.*;

public class DisclosureUnitTest {
    @Test
    public void hostIsInjected() {
        Disclosure disclosure = new Disclosure("a");
        assertEquals(disclosure.host(), "a");
    }

    @Test
    public void equalsIsSame() {
        Disclosure a = new Disclosure("a");
        Disclosure b = new Disclosure("a");
        assertEquals(a, b);
        assertEquals(b, a);
    }

    @Test
    public void equalsIsNotSame() {
        Disclosure a = new Disclosure("a");
        Disclosure b = new Disclosure("1");
        assertNotEquals(a, b);
        assertNotEquals(b, a);
    }

    @Test
    public void hashCodeIsSame() {
        Disclosure a = new Disclosure("a");
        Disclosure b = new Disclosure("a");
        assertEquals(a.hashCode(), b.hashCode());
    }

    @Test
    public void hashCodeIsNotSame() {
        Disclosure a = new Disclosure("a");
        Disclosure b = new Disclosure("1");
        assertNotEquals(a.hashCode(), b.hashCode());
    }
}