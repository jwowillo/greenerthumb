package io.greenerthumb.device;

/**
 * Device has a name, a host that it publishes from, and a host it can be commanded from.
 */
public interface Device {
    String name();
    String publishHost();
    String commandHost();
}
