# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.0.1] - 2025-07-05

### Fixed

- Removed vestigial reference to panicking behavior in `Task.Join()` documentation
- Fixed version tag links in changelog

## [2.0.0] - 2025-07-05

### Removed

- Panics are no longer automatically handled. A panic in a task will crash the program, unless you recover from it on your own.
  - Removed `Task.JoinWithoutPanicking()`

### Changed

- Updated GitHub Actions workflow
  - Tests on both "stable" and "oldstable" Go versions
  - Simplify

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

[Unreleased]: https://github.com/Quantaly/shunt/compare/v2.0.1...HEAD
[2.0.1]: https://github.com/Quantaly/shunt/compare/v2.0.0...v2.0.1
[2.0.0]: https://github.com/Quantaly/shunt/compare/v1.1.1...v2.0.0
[1.1.1]: https://github.com/Quantaly/shunt/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/Quantaly/shunt/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/Quantaly/shunt/releases/tag/v1.0.0
