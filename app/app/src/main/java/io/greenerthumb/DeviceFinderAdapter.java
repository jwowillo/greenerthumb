package io.greenerthumb;

import android.support.annotation.NonNull;
import android.support.v7.widget.CardView;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.ViewGroup;
import android.widget.TextView;

import java.util.ArrayList;
import java.util.List;
import java.util.function.Consumer;

import io.greenerthumb.device.Device;
import io.greenerthumb.device.DeviceFinder;

/**
 * DeviceFinderAdapter adapts Devices found from a DeviceFinder to a RecyclerView.
 */
public class DeviceFinderAdapter extends RecyclerView.Adapter<DeviceFinderAdapter.ViewHolder> {
    /**
     * ViewHolder holds the cards used to display Devices.
     */
    static class ViewHolder extends RecyclerView.ViewHolder {
        TextView deviceName;
        TextView publishHost;
        TextView commandHost;

        ViewHolder(CardView view, TextView deviceName, TextView publishHost, TextView commandHost) {
            super(view);

            this.deviceName = deviceName;
            this.publishHost = publishHost;
            this.commandHost = commandHost;
        }
    }

    private final List<Device> devices = new ArrayList<>();

    DeviceFinderAdapter(DeviceFinder finder, Consumer<Runnable> threadCommunicator) {
        finder.addDevicesHandler(received -> {
            devices.clear();
            received.forEach(devices::add);

            threadCommunicator.accept(this::notifyDataSetChanged);
        });
    }

    @Override
    public @NonNull ViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        CardView view = (CardView) LayoutInflater.from(parent.getContext())
                .inflate(R.layout.device_card_view, parent, false);
        TextView deviceName = view.findViewById(R.id.device_name);
        TextView publishHost = view.findViewById(R.id.publish_host);
        TextView commandHost = view.findViewById(R.id.command_host);
        return new ViewHolder(view, deviceName, publishHost, commandHost);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        Device device = devices.get(position);
        holder.deviceName.setText(device.name());
        holder.publishHost.setText(device.publishHost());
        holder.commandHost.setText(device.commandHost());
    }

    @Override
    public int getItemCount() {
        return devices.size();
    }
}