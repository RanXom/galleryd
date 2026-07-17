# Changelog

## [v0.1.0] - 2026-07-13

### Added

- Directory scanning
- EXIF metadata extraction
- Thumbnail generation and caching
- HTTP API
- CLI configuration
- API tests

## [v0.2.0] - 2026-07-17

### Added

- Automatic filesystem watching using fsnotify.
- Recursive directory watching.
- Debounced gallery reloads after filesystem changes.
- Manual `POST /api/reload` endpoint.
- Service reload tests.
- Watcher debounce tests.

### Changed

- Gallery loading refactored into reloadable service.
- Gallery now updates without restarting the server.
