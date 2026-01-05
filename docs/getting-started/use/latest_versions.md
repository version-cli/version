# Retrieve Latest Versions

`version` can retrieve the latest version of a dependency from a datasource.

## Usage

```bash
version latest [flags] DEPNAME
```

## Flags

| Flag | Shorthand | Description | Required | Default |
|---|---|---|---|---|
| `--datasource` | `-d` | Datasource to use | Yes | |
| `--major` | `-M` | Set version level to major | No | `false` |
| `--minor` | `-m` | Set version level to minor | No | `false` |
| `--patch` | `-p` | Set version level to patch | No | `true` |
| `--count` | `-c` | Number of versions to retrieve | No | `1` |
| `--semver` | `-S` | Use semver as the versioning type | No | `true` |

## Examples

### Retrieve latest version

```bash
version latest --datasource docker ghcr.io/version-cli/version
# Output: 0.1.3
```

### Retrieve latest 3 versions

```bash
version latest --datasource docker --count 3 ghcr.io/version-cli/version
# Output:
# 0.1.1
# 0.1.2
# 0.1.3
```

### Retrieve latest minor version

```bash
version latest --datasource docker --minor ghcr.io/version-cli/version
# Output: 0.1.3
```

### Retrieve latest major version

```bash
version latest --datasource docker --major ghcr.io/version-cli/version
# Output: 0.1.3
```
