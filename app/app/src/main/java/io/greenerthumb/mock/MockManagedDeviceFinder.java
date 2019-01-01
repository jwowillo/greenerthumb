package io.greenerthumb.mock;

import io.greenerthumb.device.DevicesHandler;
import io.greenerthumb.device.ManagedDeviceFinder;

/**
 * MockManagedDeviceFinder is a ManagedDeviceFinder implemented with stubs.
 */
public class MockManagedDeviceFinder implements ManagedDeviceFinder {
    @Override
    public void addDevicesHandler(DevicesHandler handler) { }

    @Override
    public void start() { }

    @Override
    public void stop() { }
}
