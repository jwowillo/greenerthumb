package io.greenerthumb.device;

import io.greenerthumb.message.data.Disclosure;

/**
 * DisclosureDeviceFactory is a factory which creates Devices from Disclosures.
 */
@FunctionalInterface
public interface DisclosureDeviceFactory {
    Device create(Disclosure disclosure);
}
