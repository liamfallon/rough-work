# bp-rs-f

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] bp-rs-f`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree bp-rs-f`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init bp-rs-f
kpt live apply bp-rs-f --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
