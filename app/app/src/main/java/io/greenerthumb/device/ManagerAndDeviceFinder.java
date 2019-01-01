package io.greenerthumb.device;

import io.greenerthumb.manage.Manager;

/**
 * ManagerAndDeviceFinder helps implement ManagedDeviceFinder by composing the Manager's methods and
 * the DeviceFinder's methods.
 */
public class ManagerAndDeviceFinder implements ManagedDeviceFinder {
    private final Manager manager;
    private final DeviceFinder deviceFinder;

    public ManagerAndDeviceFinder(Manager manager, DeviceFinder deviceFinder) {
        this.manager = manager;
        this.deviceFinder = deviceFinder;
    }

    @Override
    public void addDevicesHandler(DevicesHandler handler) {
        deviceFinder.addDevicesHandler(handler);
    }

    @Override
    public void start() {
        manager.start();
    }

    @Override
    public void stop() {
        manager.stop();
    }
}
