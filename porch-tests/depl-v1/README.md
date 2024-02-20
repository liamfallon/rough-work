# bp

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] bp`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree bp`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init bp
kpt live apply bp --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
