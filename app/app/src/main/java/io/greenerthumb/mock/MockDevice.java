package io.greenerthumb.mock;

import io.greenerthumb.device.Device;

/**
 * MockDevice is a Device with all getters injected in the constructor.
 */
public class MockDevice implements Device {
    private final String host;

    MockDevice(String host) {
        this.host = host;
    }

    @Override
    public String host() {
        return host;
    }
}
