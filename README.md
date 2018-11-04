# `greenerthumb`

`greenerthumb` is a garden automation package.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make greenerthumb` builds the sub-components.
* `make pi` builds programs that run on PIs.
* `make host` builds programs for hosts.
* `make test` builds `builds the sub-components' tests.

## Testing

All the sub-components' tests can be run with:

```
./test
```

## Utilities

* `run_sensors`: Runs all sensors into `bullhorn/publish`.
* `run_air`: Runs `sense/air` into `bullhorn/publish`.
* `run_soil`: Runs `sense/soil` into `bullhorn/publish`.
* `run_logger`: Runs `bullhorn/subscribe` into `log`.
* `run_plotter`: Runs `bullhorn/subscribe` into `log` and `plot`.
