# `greenerthumb`

`greenerthumb` is a garden automation package.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` and `make greenerthumb` build all targets.
* `make air` builds programs for the air sensor that run on PIs.
* `make soil` builds programs for the soil sensor that run on PIs.
* `make test` builds the subcomponents' tests.

## Testing

All the sub-components' tests can be run with:

```
./test
```

## Utilities

`activate.sh` must be run in the project root before running any of the
utilities to set up aliases. `deactivate.sh` removes the aliases.

* `greenerthumb-run-air`: Runs `sense/air` into `bullhorn/publish`.
* `greenerthumb-run-soil`: Runs `sense/soil` into `bullhorn/publish`.
* `greenerthumb-run-logger`: Runs `bullhorn/subscribe` into `log`.
* `greenerthumb-run-plotter`: Runs `bullhorn/subscribe` into `log` and `plot`.

## Deployment

Deployment is facilitated by 3 scripts.

* `copy-keys`: Copies SSH keys to a passed remote user and host.
* `deploy`: Builds a passed target, copies all files for the target to a passed
  remote user and host, and restarts the target.
* `cat-error-log`: Cats the error-log for a deployed `greenerthumb` program on a
  passed remote user and host.
