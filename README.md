# version

`version` is a small cli tool that bumps your version number for you and retrieves the latest version of a specified
datasource.

## Supported versioning schemes

Currently, `version` only supports [SemVer 2.0](https://semver.org/). SemVer 2.0 is one of the most used versioning schemes.

## Supported datasources

Currently, `version` only supports Docker registries as a datasource. At the moment it doesn't support authentication, but
other than that, if `docker`-cli pulls from it, `version` should support it. This means that among others, `ghcr.io` `hub.docker.com`,
`gcr.io`, and, `docker.pkg.dev` should be supported.
