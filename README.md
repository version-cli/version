<div style="text-align: center">
    <div>
        <img alt="logo" src="./docs/assets/images/logo.svg"/>
    </div>
</div>

<!-- # --8<-- [start:description] -->
<div style="text-align: center">
    <h1>Version</h1>
    <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/version-cli/version?style=for-the-badge&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGVuYWJsZS1iYWNrZ3JvdW5kPSJuZXcgMCAwIDI0IDI0IiBoZWlnaHQ9IjI0cHgiIHZpZXdCb3g9IjAgMCAyNCAyNCIgd2lkdGg9IjI0cHgiIGZpbGw9IiNGRkZGRkYiPjxnPjxwYXRoIGQ9Ik0wLDBoMjR2MjRIMFYweiIgZmlsbD0ibm9uZSIvPjwvZz48Zz48Zz48cGF0aCBkPSJNMjEuNDEsMTEuNDFsLTguODMtOC44M0MxMi4yMSwyLjIxLDExLjcsMiwxMS4xNywySDRDMi45LDIsMiwyLjksMiw0djcuMTdjMCwwLjUzLDAuMjEsMS4wNCwwLjU5LDEuNDFsOC44Myw4LjgzIGMwLjc4LDAuNzgsMi4wNSwwLjc4LDIuODMsMGw3LjE3LTcuMTdDMjIuMiwxMy40NiwyMi4yLDEyLjIsMjEuNDEsMTEuNDF6IE0xMi44MywyMEw0LDExLjE3VjRoNy4xN0wyMCwxMi44M0wxMi44MywyMHoiLz48Y2lyY2xlIGN4PSI2LjUiIGN5PSI2LjUiIHI9IjEuNSIvPjwvZz48L2c%2BPC9zdmc%2B">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/version-cli/version?style=for-the-badge&logo=go&color=00ADD8">
    <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/version-cli/version/release.yaml?logo=github&style=for-the-badge"/>
    <a href="https://sonarcloud.io/summary/new_code?id=version-cli_version"><img alt="Sonar Quality Gate" src="https://img.shields.io/sonar/quality_gate/version-cli_version?server=https%3A%2F%2Fsonarcloud.io&logo=sonarcloud&style=for-the-badge&color=F3702A"/></a>
    <a href="https://snyk.io/test/github/version-cli/version"><img alt="Snyk Security" src="https://img.shields.io/badge/monitored-%23914dc2?style=for-the-badge&logo=snyk&label=snyk%20security&color=4C4A73"></a>
</div>

`version` is a small cli tool that bumps your version number for you and retrieves the latest version of a specified
datasource.

## Supported versioning schemes

Currently, `version` only supports [SemVer 2.0](https://semver.org/). SemVer 2.0 is one of the most used versioning schemes.

## Supported datasources

Currently, `version` only supports Docker registries as a datasource.
At the moment `version` doesn't support authentication,
but other than that if the `docker`-cli pulls it, `version` supports it.
This means that among others, `ghcr.io` `hub.docker.com`, `gcr.io`, and `docker.pkg.dev` are supported.

<!-- # --8<-- [end:description] -->
