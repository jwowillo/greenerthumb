package io.greenerthumb;

import android.content.Intent;
import android.content.SharedPreferences;
import android.support.v4.widget.SwipeRefreshLayout;
import android.support.v7.app.AlertDialog;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.support.v7.preference.PreferenceManager;

import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.Menu;
import android.view.MenuInflater;
import android.view.MenuItem;

import io.greenerthumb.app.BroadcastDisclosureManagedDeviceFinderFactory;
import io.greenerthumb.device.ManagedDeviceFinder;
import io.greenerthumb.mock.MockManagedDeviceFinder;

/**
 * MainActivity starts the app.
 */
public class MainActivity extends AppCompatActivity {
    private SwipeRefreshLayout refresher; // Set at onCreate.
    private RecyclerView recyclerView; // Set at onCreate.

    private ManagedDeviceFinder deviceFinder = new MockManagedDeviceFinder(); // Avoids NPEs.

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        PreferenceManager.setDefaultValues(this, R.xml.settings, false);

        refresher = findViewById(R.id.refresher);
        refresher.setOnRefreshListener(() -> {
            stopDeviceFinder();
            startDeviceFinder();
            refresher.setRefreshing(false);
        });

        recyclerView = findViewById(R.id.recycler_view);
        recyclerView.setHasFixedSize(true);
        RecyclerView.LayoutManager manager = new LinearLayoutManager(this);
        recyclerView.setLayoutManager(manager);
    }

    @Override
    protected void onStart() {
        super.onStart();

        startDeviceFinder();
    }

    @Override
    protected void onStop() {
        super.onStop();

        stopDeviceFinder();
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        MenuInflater inflater = getMenuInflater();
        inflater.inflate(R.menu.menu, menu);
        return super.onCreateOptionsMenu(menu);
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case R.id.action_settings:
                Intent intent = new Intent(this, SettingsActivity.class);
                startActivity(intent);
                return true;
            default:
                return super.onOptionsItemSelected(item);
        }
    }

    /**
     * startDeviceFinder attempts to start the DeviceFinder and shows an error if not possible.
     */
    private void startDeviceFinder() {
        try {
            SharedPreferences preferences = PreferenceManager.getDefaultSharedPreferences(this);
            String rawPort = preferences.getString("port", "");
            int port;
            if (rawPort == null) {
                port = -1;
            } else {
                port = Integer.parseInt(rawPort);
            }

            deviceFinder = new BroadcastDisclosureManagedDeviceFinderFactory(port).create();

            RecyclerView.Adapter adapter = new DeviceFinderAdapter(
                    deviceFinder,
                    this::runOnUiThread);
            recyclerView.setAdapter(adapter);

            new Thread(deviceFinder::start).start();
        } catch (Exception exception) {
            AlertDialog dialog = new AlertDialog.Builder(MainActivity.this).create();
            dialog.setTitle("Error creating device-finder:");
            dialog.setMessage(exception.getMessage());
            dialog.show();
        }
    }

    /**
     * stopDeviceFinder stops the DeviceFinder.
     */
    private void stopDeviceFinder() {
        deviceFinder.stop();
    }
}
