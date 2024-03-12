# network-function

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] network-function`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree network-function`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init network-function
kpt live apply network-function --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
