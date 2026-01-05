# GitHub Action

`version` can be used in GitHub Actions to bump version numbers or retrieve the latest version of a dependency.

## Usage

```yaml
steps:
  - name: Bump version
    id: bump
    uses: version-cli/action@v1
    with:
      command: bump
      args: 1.2.3 --minor
```

## Inputs

| Input | Description | Required | Default |
|---|---|---|---|
| `command` | Command to run (`bump` or `latest`) | Yes | |
| `args` | Arguments to pass to the command | Yes | |

## Outputs

| Output | Description |
|---|---|
| `version` | The output of the command |

## Examples

### Bump version

```yaml
steps:
  - name: Bump version
    id: bump
    uses: version-cli/action@v1
    with:
      command: bump
      args: 1.2.3 --minor
  - name: Print version
    run: echo ${{ steps.bump.outputs.version }}
    # Output: 1.3.0
```

### Retrieve latest version

```yaml
steps:
  - name: Retrieve latest version
    id: latest
    uses: version-cli/action@v1
    with:
      command: latest
      args: --datasource docker ghcr.io/version-cli/version
  - name: Print version
    run: echo ${{ steps.latest.outputs.version }}
    # Output: 0.1.3
```

### Update workflow with latest versions

This example shows how to use `version` in a GitHub Action to automatically update a workflow file with the latest 3 versions of a Docker image (in this case, `kindest/node`). This is useful for matrix builds where you want to test against the latest versions of a dependency.

```yaml
name: Update Kindest Node Versions

on:
  schedule:
    - cron: '0 0 * * 0' # Run weekly
  workflow_dispatch:

jobs:
  update-versions:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v6

      - name: Get latest versions
        id: versions
        uses: version-cli/action@<version>
        with:
          command: latest
          args: --datasource docker --count 3 kindest/node

      - name: Update workflow
        run: |
          # Convert newlines to JSON array
          VERSIONS_JSON=$(echo "${{ steps.versions.outputs.version }}" | jq -R -s -c 'split("\n")[:-1]')
          
          # Update the target workflow file using yq
          yq -i ".jobs.test.strategy.matrix.k8s-version = $VERSIONS_JSON" .github/workflows/test.yaml
          
      - name: Create Pull Request
        uses: peter-evans/create-pull-request@v6
        with:
          title: "ci: update kindest/node versions"
          body: "Updates the kindest/node versions in the test workflow to the latest 3 versions."
          branch: "update-kindest-node-versions"
```
