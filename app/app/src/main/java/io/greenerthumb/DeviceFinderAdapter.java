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
        TextView host;

        ViewHolder(CardView view, TextView host) {
            super(view);

            this.host = host;
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
        TextView host = view.findViewById(R.id.host);
        return new ViewHolder(view, host);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolder holder, int position) {
        Device device = devices.get(position);
        holder.host.setText(device.host());
    }

    @Override
    public int getItemCount() {
        return devices.size();
    }
}