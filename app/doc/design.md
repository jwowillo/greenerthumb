# `app` Design

`app` shows devices on the network.

The major portion of the work the app does is deserializing messages from the
network. To accomodate this, a layered approach is taken where base receivers
are defined for the different protocols and converters and converting receivers
can get the received data all the way to the correct message with high code
reuse. Higher level receivers can share lower level receivers to improve
resource utilization. Testing is also simplified because most of the testing can
occur at the highest level of data. Each receiver only receives one kind of
message. This allows models to modularly choose which receivers they need and
further simplifies testing.

One con of the layered approach is that it's difficult to construct the actual
instances used for the app. Factories are provided in the app package to
facilitate the full construction.

The mock, network, and app packages aren't tested. The mock package isn't tested
because the gains would be marginal. The network package would be difficult to
test and can be reasonably verified just by running the app since all its
receivers are immutable and don't have branches. The app package is difficult to
test because it uses lots of real resources. The Android app itself isn't tested
because it is difficult, even though it probably should be.
