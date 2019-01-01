package io.greenerthumb.app;

import io.greenerthumb.device.CreationException;
import io.greenerthumb.device.DeviceFinder;
import io.greenerthumb.device.DisclosureDeviceFinder;
import io.greenerthumb.device.ManagedDeviceFinder;
import io.greenerthumb.device.ManagedDeviceFinderFactory;
import io.greenerthumb.device.ManagerAndDeviceFinder;
import io.greenerthumb.message.Message;
import io.greenerthumb.message.MessageDeserializer;
import io.greenerthumb.message.data.Disclosure;
import io.greenerthumb.message.data.DisclosureDeserializer;
import io.greenerthumb.mock.MockDisclosureDeviceFactory;
import io.greenerthumb.network.BroadcastReceiver;
import io.greenerthumb.receiver.ConvertingReceiver;
import io.greenerthumb.receiver.Receiver;

/**
 * BroadcastDisclosureManagedDeviceFinderFactory makes ManagedDeviceFinders that receives
 * Disclosures via broadcast.
 */
public class BroadcastDisclosureManagedDeviceFinderFactory implements ManagedDeviceFinderFactory {
    private final int port;

    /**
     * @param port to receive broadcasts.
     */
    public BroadcastDisclosureManagedDeviceFinderFactory(int port) {
        this.port = port;
    }

    @Override
    public ManagedDeviceFinder create() throws CreationException {
        try {
            BroadcastReceiver byteReceiver = new BroadcastReceiver(port);
            Receiver<Message> messageReceiver = new ConvertingReceiver<>(
                    byteReceiver,
                    new MessageDeserializer());
            Receiver<Disclosure> receiver = new ConvertingReceiver<>(
                    messageReceiver,
                    new DisclosureDeserializer());
            DeviceFinder finder = new DisclosureDeviceFinder(
                    receiver,
                    new MockDisclosureDeviceFactory());

            return new ManagerAndDeviceFinder(byteReceiver, finder);
        } catch (Exception exception) {
            throw new CreationException(exception.getMessage());
        }
    }
}
