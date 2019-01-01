package io.greenerthumb.device;

import java.util.ArrayList;
import java.util.Collection;
import java.util.HashSet;
import java.util.Set;

import io.greenerthumb.message.data.Disclosure;
import io.greenerthumb.receiver.Receiver;

/**
 * DisclosureDeviceFinder finds Devices via the Disclosure message.
 *
 * Order is preserved as Devices are found.
 */
public class DisclosureDeviceFinder implements DeviceFinder {
    private final Collection<DevicesHandler> handlers = new ArrayList<>();
    private final DisclosureDeviceFactory factory;

    private final Set<Disclosure> seen = new HashSet<>();
    private final Collection<Device> devices = new ArrayList<>();

    /**
     * @param receiver of Disclosure messages.
     * @param factory that converts Disclosures to Devices.
     */
    public DisclosureDeviceFinder(Receiver<Disclosure> receiver, DisclosureDeviceFactory factory) {
       this.factory = factory;

       receiver.addReceiveHandler(this::receiveDisclosure);
    }

    @Override
    public void addDevicesHandler(DevicesHandler handler) {
        handlers.add(handler);
    }

    private void receiveDisclosure(Disclosure disclosure) {
        if (seen.contains(disclosure)) {
            return; // Already seen this disclosure.
        }
        seen.add(disclosure);
        devices.add(factory.create(disclosure));
        for (DevicesHandler handler : handlers) {
            handler.handle(devices);
        }
    }
}
