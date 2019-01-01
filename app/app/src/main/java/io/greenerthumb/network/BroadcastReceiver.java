package io.greenerthumb.network;

import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.SocketException;
import java.net.UnknownHostException;
import java.util.ArrayList;
import java.util.Collection;

import io.greenerthumb.collection.ArrayView;
import io.greenerthumb.manage.Manager;
import io.greenerthumb.receiver.ReceiveHandler;
import io.greenerthumb.receiver.Receiver;

/**
 * BroadcastReceiver receives ArrayView<Byte>s sent via broadcast.
 *
 * The Receiver can only be started and stopped once.
 */
public class BroadcastReceiver implements Receiver<ArrayView<Byte>>, Manager {
    private final DatagramSocket socket;
    private final Collection<ReceiveHandler<ArrayView<Byte>>> handlers = new ArrayList<>();
    private boolean isStarted = false;

    /**
     * @param port to receive broadcasts on.
     * @throws SocketException if the DatagramSocket couldn't be constructed.
     */
    public BroadcastReceiver(int port) throws SocketException, UnknownHostException {
        this.socket = new DatagramSocket(port, InetAddress.getByName("0.0.0.0"));
        this.socket.setBroadcast(true);
    }

    @Override
    public void addReceiveHandler(ReceiveHandler<ArrayView<Byte>> handler) {
        handlers.add(handler);
    }

    @Override
    public void start() {
        if (socket.isClosed() || isStarted) {
            return;
        }
        isStarted = true;
        byte[] buffer = new byte[256];
        while (!socket.isClosed()) {
            DatagramPacket packet = new DatagramPacket(buffer, buffer.length);
            try {
                socket.receive(packet);
            } catch (IOException exception) {
                socket.disconnect();
                break;
            }
            // Newlines are removed here for simplicity. If this BroadcastReceiver needs to be
            // used for other types of messages in the future, a better design would be to
            // have a NewlineRemovingConverter that checks if the last byte is a newline.
            ArrayView<Byte> view = new ArrayView<>(box(packet.getData()))
                    .viewOf(0, packet.getLength() - 1); // Messages end in newlines.
            for (ReceiveHandler<ArrayView<Byte>> handler : handlers) {
                handler.receive(view);
            }
        }
    }

    @Override
    public void stop() {
        socket.close();
    }

    private static Byte[] box(byte[] data) {
        Byte[] boxed = new Byte[data.length];
        for (int i = 0; i < data.length; i++) {
            boxed[i] = data[i];
        }
        return boxed;
    }
}
