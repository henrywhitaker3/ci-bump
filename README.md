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

To update the patch of `appVersion`, the minor version of `app.version` and the major version of `app.otherVersion`:

```
ci-dump yaml --patch '.appVersion' --minor '.app.version' --major '.app.otherVersion' test.yaml
```

To set the `appVersion` field to the value `bongo`:

```
ci-bump yaml --set '.appVersion=bongo' test.yaml
```

## CI

### GitHub

```yaml
- uses: henrywhitaker3/ci-bump@main
  with:
    cmd: ci-bump yaml --patch '.versions.app' --minor '.versions.chart' --major '.versions.lock' demo.yaml
```

### GitLab

```yaml
update helm version:
  image: ghcr.io/henrywhitaker3/ci-bump:latest
  script:
    - ci-bump --patch '.appVersion' chart/Chart.yaml
```
