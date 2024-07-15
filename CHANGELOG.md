# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic
Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- New `--bind` option which allows you to specify what address to bind to, defaults to `0.0.0.0`.
### Changed
- Updated to Go version `1.22`.

## [v1.1.1] - 2023-12-04
### Fixed
- Fixed issue with incorrect headers and response code for browser/system level proxy settings
([#1](https://github.com/jthomperoo/simple-proxy/issues/1)).

## [v1.1.0] - 2022-04-08
### Added
- Basic auth support, can provide `--basic-auth 'username:password'` which the proxy then checks for valid auth
provided in the `Proxy-Authentication` header.
- Can choose if auth should be logged with the `--log-auth` option.
- Can choose to log all request headers using the `--log-headers` option.

## [v1.0.0] - 2022-04-07
### Added
- Initial release, self-contained binary that allows hosting a simple proxy.
    - Supports HTTP and HTTPS.
    - Supports choosing which port.
    - Supports printing binary version number.
    - Supports specifying paths to certificate and private key file to use.
    - Logs each proxied connection.
    - Supports log options can be supplied using `glog`.
        - Can choose the log verbosity with the `-v` flag.
        - Can choose to log to a file.

[Unreleased]: https://github.com/jthomperoo/simple-proxy/compare/v1.1.1...HEAD
[v1.1.1]:https://github.com/jthomperoo/simple-proxy/compare/v1.1.0...v1.1.1
[v1.1.0]:https://github.com/jthomperoo/simple-proxy/compare/v1.0.0...v1.1.0
[v1.0.0]: https://github.com/jthomperoo/simple-proxy/releases/tag/v1.0.0
