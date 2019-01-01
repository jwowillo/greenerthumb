package io.greenerthumb.mock;

import io.greenerthumb.device.Device;

/**
 * MockDevice is a Device with all getters injected in the constructor.
 */
public class MockDevice implements Device {
    private final String name;
    private final String publishHost;
    private final String commandHost;

    MockDevice(String name, String publishHost, String commandHost) {
        this.name = name;
        this.publishHost = publishHost;
        this.commandHost = commandHost;
    }

    @Override
    public String name() {
        return name;
    }

    @Override
    public String publishHost() {
        return publishHost;
    }

    @Override
    public String commandHost() {
        return commandHost;
    }
}
