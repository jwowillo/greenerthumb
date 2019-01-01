package io.greenerthumb.message.data;

import java.util.Objects;

/**
 * Disclosure message.
 */
public class Disclosure {
    private final String deviceName;
    private final String publishHost;
    private final String commandHost;

    public Disclosure(String deviceName, String publishHost, String commandHost) {
        this.deviceName = deviceName;
        this.publishHost = publishHost;
        this.commandHost = commandHost;
    }

    public String deviceName() {
        return deviceName;
    }

    public String publishHost() {
        return publishHost;
    }

    public String commandHost() {
        return commandHost;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Disclosure that = (Disclosure) o;
        return Objects.equals(deviceName, that.deviceName) &&
                Objects.equals(publishHost, that.publishHost) &&
                Objects.equals(commandHost, that.commandHost);
    }

    @Override
    public int hashCode() {
        return Objects.hash(deviceName, publishHost, commandHost);
    }
}
