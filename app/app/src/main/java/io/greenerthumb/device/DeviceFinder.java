package io.greenerthumb.device;

/**
 * DeviceFinder finds Devices.
 */
@FunctionalInterface
public interface DeviceFinder {
    /**
     * addDevicesHandler adds a DevicesHandler that is triggered whenever
     * the devices found changes.
     *
     * @param handler to trigger.
     */
    void addDevicesHandler(DevicesHandler handler);
}
