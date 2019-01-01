package io.greenerthumb.device;

/**
 * DevicesHandler handles an Iterable of devices.
 */
@FunctionalInterface
public interface DevicesHandler {
    void handle(Iterable<Device> devices);
}
