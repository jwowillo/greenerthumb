# `plot` Design

`plot`s `greenerthumb` JSON messages from STDIN.

## UI

`plot` will convert input messages to points on a line-graph. The expected
message format is:

```
{
  "name": <message_name>,
  "timestamp": <timestamp>,
  <name>: <value>,...
}
```

Messages don't necessarily need to be `greenerthumb` messages.  They just need
to fit this format.

Each non-ID and non-timestamp field will become a line in the graph. The name of
the line will be determined by concatenating the `<message_name>` and the
field's `<name>`. The line's will be assigned unique random colors which will be
displayed in a legend with the message name's on the right side of the plot.

Each line will be overlayed to allow trend comparison. To do this ,each line
will have units normalized to eachother. The x-axis will have units of hours
scaled to the period of all the received messages.

If received messages have the same timestamp, the newest message will overwrite
the older messages.

## Program

```
./plot
```

## Example

```
./plot

< {"name": "Soil", "timestamp": 0, "Moisture": 0.37}
< {"name": "Air", "timestamp": 3600, "Temperature": 84, "Humidity": 0.54}
< {"name": "Soil", "timestamp": 3600, "Moisture": 0.35}
< {"name": "Air": "timestamp": 7200, "Temperature": 82, "Humidity": 0.51}
```

This will plot 3 lines labelled "Soil Moisture", "Air Temperature", and "Air
Humidity". Each will have 2 points. The "Soil"-line will start at hour 0 and
finish at hour 1. The "Air"-lines will start at hour 1 and finish at hour 2. The
entire plot will occupy 2 hours.
