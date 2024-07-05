# CI Bump

A cli tool for incrementing semver versions in files in CI pipelines

## Usage

### YAML

To update versions in a yaml file, you can pass `yq` formatted selectors of the field to update, e.g.:

To update the patch version of the field `appVersion` in file `test.yaml`:

```
ci-bump yaml --patch '.appVersion' test.yaml
```

To update the patch version of `appVersion` and `version`:

```
ci-dump yaml --patch '.appVersion' --patch '.verison' test.yaml
```

To update the patch of `appversion`, the minor version of `app.version` and the major version of `app.otherVersion`:

```
ci-dump yaml --patch '.appVersion' --minor '.app.version' --major '.app.otherVersion' test.yaml
```
