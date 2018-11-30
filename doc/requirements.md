# `greenerthumb` Requirements

1. Each hardware component must be optional.
2. Sense:
    a. Air temperature.
    b. Soil moisture.
3. Publish sensed data on a connected network.
4. Log sensed data from the network into rotating logs.
5. Plot sensed data from the network.
6. Allow plots to be saved.
7. Remove anomalous data from sensors.
8. Include an option in `bullhorn/subscribe` to automatically reconnect.
9. Embed the icon in the `plot` binary.
10. Don't hide errors in the utility scripts.
11. Emulate sensors if real sensors aren't detected.
12. Provide a log processor.
13. Provide install tools.
14. Include data ranges in the plot legend.

## Future Requirements

* Sense soil pH.
* Push a notification to a connected phone when the soil moisture from the
  network is below a user set threshold.
* Water plants when the soil moisture from the network is below the user set
  threshold.
* Make sensors wireless.
* Provide IC packages for sensors.
* Monitor networks.
* Provide a web interface.
* Discover addresses and ports automatically.
