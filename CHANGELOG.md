# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic
Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/jthomperoo/simple-proxy/commits/main
