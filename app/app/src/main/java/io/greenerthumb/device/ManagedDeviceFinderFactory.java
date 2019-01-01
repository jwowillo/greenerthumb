package io.greenerthumb.device;

/**
 * ManagedDeviceFinderFactory creates ManagedDeviceFinders.
 *
 * Throws a CreationException if the ManagedDeviceFinder couldn't be created.
 */
@FunctionalInterface
public interface ManagedDeviceFinderFactory {
    ManagedDeviceFinder create() throws CreationException;
}
