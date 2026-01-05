# Installation

## Homebrew

```bash
brew tap version-cli/version
brew install version
```

## Docker

You can run `version` using Docker:

```bash
docker run ghcr.io/version-cli/version:latest ...
```

## Go

If you have Go 1.24+ installed, you can install `version` using `go install`:

```bash
go install github.com/version-cli/version@latest
```

## Manual

You can download the pre-compiled binaries from the [releases page](https://github.com/version-cli/version/releases) and copy them to a location in your `$PATH`.

## Shell Completion

`version` provides completion scripts for Bash, Zsh, Fish, and PowerShell.

### Bash

To load completions for every new session, execute once:

#### Linux

```bash
version completion bash > /etc/bash_completion.d/version
```

#### macOS

```bash
version completion bash > $(brew --prefix)/etc/bash_completion.d/version
```

### Zsh

If shell completion is not already enabled in your environment, you will need to enable it. You can execute the following once:

```zsh
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

To load completions for each session, execute once:

```zsh
version completion zsh > "${fpath[1]}/_version"
```

You will need to start a new shell for this setup to take effect.

### Fish

To load completions for each session, execute once:

```fish
version completion fish > ~/.config/fish/completions/version.fish
```

### PowerShell

To load completions in your current shell session:

```powershell
goreleaser completion powershell | Out-String | Invoke-Expression
```

To load completions for every new session, add the output of the above command to your powershell profile.
