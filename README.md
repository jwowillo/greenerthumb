# `greenerthumb`

`greenerthumb` is a garden automation tool that allows multiple devices to
cooperate and effectively manage garden operations as packages.

`greenerthumb` is in active development.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` build all targets.
* `make <component>` builds everything for the component, where component is one
  of air, soil, logger, or plotter.
* `make test` builds the programs' tests.

`make clean` should always be run between building a component and other builds.
This is because building a component might cause dependencies to be built for a
different target architecture.

## Testing

All the programs' tests can be run with:

```
./test
```

## Utilities

Utilities are provided in the 'util' directory and are expected to be run from
the project root.

## Running

`util/activate.sh` must be sourced in the project root before running any of the
components to set up aliases. Running instructions are described in the
'README.md' in the build directory after building the component.

## Deployment

Deployment is facilitated by 3 scripts.

* `util/copy-keys`: Copies SSH keys to a passed remote user and host.
* `util/deploy`: Builds a passed target, copies all files for the target to a
  passed remote user and host, and restarts the target.
* `util/cat-error-log`: cats the error-log for a deployed `greenerthumb` program
  on a passed remote user and host.
* `util/rm-error-log`: rms the error-log for a deployed `greenerthumb` program
  on a passed remote user and host.
