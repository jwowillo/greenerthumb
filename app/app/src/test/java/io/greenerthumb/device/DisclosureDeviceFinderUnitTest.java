package io.greenerthumb.device;

import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

import io.greenerthumb.message.data.Disclosure;
import io.greenerthumb.mock.MockDisclosureDeviceFactory;
import io.greenerthumb.mock.MockReceiver;

import static org.junit.Assert.*;

public class DisclosureDeviceFinderUnitTest {
    @Test
    public void handlersAreTriggered() {
        boolean[] didTrigger = {false, false, false};

        MockReceiver<Disclosure> receiver = new MockReceiver<>();
        DisclosureDeviceFactory factory = new MockDisclosureDeviceFactory();
        DeviceFinder finder = new DisclosureDeviceFinder(receiver, factory);

        finder.addDevicesHandler(devices -> didTrigger[0] = true);
        finder.addDevicesHandler(devices -> didTrigger[1] = true);
        finder.addDevicesHandler(devices -> didTrigger[2] = true);

        receiver.receive(new Disclosure("a"));

        assertTrue(didTrigger[0]);
        assertTrue(didTrigger[1]);
        assertTrue(didTrigger[2]);
    }

    @Test
    public void handlersAreNotReTriggered() {
        int[] triggeredCount = {0, 0, 0};

        MockReceiver<Disclosure> receiver = new MockReceiver<>();
        DisclosureDeviceFactory factory = new MockDisclosureDeviceFactory();
        DeviceFinder finder = new DisclosureDeviceFinder(receiver, factory);

        finder.addDevicesHandler(devices -> triggeredCount[0]++);
        finder.addDevicesHandler(devices -> triggeredCount[1]++);
        finder.addDevicesHandler(devices -> triggeredCount[2]++);

        receiver.receive(new Disclosure("a"));
        receiver.receive(new Disclosure("a"));

        assertEquals(triggeredCount[0], 1);
        assertEquals(triggeredCount[1], 1);
        assertEquals(triggeredCount[2], 1);
    }

    @Test
    public void devicesAddedAreCorrect() {
        MockReceiver<Disclosure> receiver = new MockReceiver<>();
        DisclosureDeviceFactory factory = new MockDisclosureDeviceFactory();
        List<Device> devices = new ArrayList<>();
        DeviceFinder finder = new DisclosureDeviceFinder(receiver, factory);
        finder.addDevicesHandler(received -> {
            devices.clear();
            for (Device device : received) {
                devices.add(device);
            }
        });

        Disclosure a = new Disclosure("a1");
        Disclosure b = new Disclosure("b1");
        Disclosure c = new Disclosure("c1");

        receiver.receive(a);

        assertEquals(1, devices.size());
        assertEquals("a1", devices.get(0).host());

        receiver.receive(b);

        assertEquals(2, devices.size());
        assertEquals("a1", devices.get(0).host());
        assertEquals("b1", devices.get(1).host());

        receiver.receive(c);

        assertEquals(3, devices.size());
        assertEquals("a1", devices.get(0).host());
        assertEquals("b1", devices.get(1).host());
        assertEquals("c1", devices.get(2).host());
    }

    @Test
    public void devicesAreNotReAdded() {
        MockReceiver<Disclosure> receiver = new MockReceiver<>();
        DisclosureDeviceFactory factory = new MockDisclosureDeviceFactory();
        DeviceFinder finder = new DisclosureDeviceFinder(receiver, factory);
        finder.addDevicesHandler(devices -> assertEquals(1, size(devices)));

        receiver.receive(new Disclosure("a"));
        receiver.receive(new Disclosure("a"));
    }

    private static <T> int size(Iterable<T> ts) {
        int size = 0;
        for (T ignored : ts) {
            size++;
        }
        return size;
    }
}