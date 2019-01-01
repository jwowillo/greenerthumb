package io.greenerthumb.device;

import org.junit.Test;

import io.greenerthumb.manage.Manager;

import static org.junit.Assert.assertTrue;

class MockManager implements Manager {
    boolean startCalled = false;
    boolean stopCalled = false;

    @Override
    public void start() {
        this.startCalled = true;
    }

    @Override
    public void stop() {
        this.stopCalled = true;
    }
}

class MockDeviceFinder implements DeviceFinder {
    boolean addDevicesHandlerCalled = false;

    @Override
    public void addDevicesHandler(DevicesHandler handler) {
        this.addDevicesHandlerCalled = true;
    }
}

public class ManagerAndDeviceFinderUnitTest {
    @Test
    public void composesMethods() {
        MockManager manager = new MockManager();
        MockDeviceFinder finder = new MockDeviceFinder();
        ManagedDeviceFinder composed = new ManagerAndDeviceFinder(manager, finder);
        composed.start();
        composed.stop();
        composed.addDevicesHandler(null);
        assertTrue(manager.startCalled);
        assertTrue(manager.stopCalled);
        assertTrue(finder.addDevicesHandlerCalled);
    }
}
