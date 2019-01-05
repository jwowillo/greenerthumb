package io.greenerthumb.message.data;

import java.util.Objects;

/**
 * Disclosure message.
 */
public class Disclosure {
    private final String host;

    public Disclosure(String host) {
        this.host = host;
    }

    public String host() {
        return host;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Disclosure that = (Disclosure) o;
        return Objects.equals(host, that.host);
    }

    @Override
    public int hashCode() {
        return host.hashCode();
    }
}
