package io.greenerthumb.mock;

import io.greenerthumb.device.Device;
import io.greenerthumb.device.DisclosureDeviceFactory;
import io.greenerthumb.message.data.Disclosure;

/**
 * MockDisclosureDeviceFactory is a DisclosureDeviceFactory that makes
 * MockDevices out of Disclosures.
 */
public class MockDisclosureDeviceFactory implements DisclosureDeviceFactory {
    @Override
    public Device create(Disclosure disclosure) {
        return new MockDevice(
                disclosure.deviceName(),
                disclosure.publishHost(),
                disclosure.commandHost());
    }
}
