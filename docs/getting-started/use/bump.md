# Bump

`version` can bump version numbers for you.

## Usage

```bash
version bump [flags] VERSION
```

## Aliases

You can also use `next` as an alias for `bump`.

```bash
version next [flags] VERSION
```

## Flags

| Flag | Shorthand | Description | Default |
|---|---|---|---|
| `--major` | `-M` | Set version level to major | `false` |
| `--minor` | `-m` | Set version level to minor | `false` |
| `--patch` | `-p` | Set version level to patch | `true` |
| `--count` | `-c` | Number of versions to bump | `1` |
| `--semver` | `-S` | Use semver as the versioning type | `true` |

## Examples

### Bump patch version

```bash
version bump 1.2.3
# Output: 1.2.4
```

### Bump minor version

```bash
version bump --minor 1.2.3
# Output: 1.3.0
```

### Bump major version

```bash
version bump --major 1.2.3
# Output: 2.0.0
```

### Bump multiple versions

```bash
version bump --count 2 1.2.3
# Output: 1.2.5
```
