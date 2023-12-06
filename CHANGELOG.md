# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- More comprehensive tests for `Task.JoinWithoutPanicking()`

### Changed

- The error returned by `Task.JoinWithoutPanicking()` now wraps the panic value if the panic value is an error

## [1.1.1] - 2023-11-29

### Changed

- Refactored internals
  - Simpler structure
  - Reduced memory usage

## [1.1.0] - 2023-11-29

### Added

- Testing with GitHub Actions
- `Task.JoinWithoutPanicking()` to return panics as errors rather than actually panicking

## [1.0.0] - 2023-03-07

### Added

- `Task` struct representing a shunted function execution
- `Do` function to start a new `Task`
- Simple example

[1.1.1]: https://github.com/Quantaly/shunt/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/Quantaly/shunt/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/Quantaly/shunt/releases/tag/v1.0.0
