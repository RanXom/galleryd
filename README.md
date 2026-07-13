<div align="center">

# galleryd

![Status](https://img.shields.io/badge/status-active-success)
![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-blue)

A fast, lightweight photo gallery backend written in Go.

galleryd scans one or more image directories, extracts EXIF metadata, generates cached WebP thumbnails, and exposes a simple HTTP API for powering web, desktop, or mobile gallery applications.

</div>

## Features

- Scan multiple gallery directories
- Extract EXIF metadata
- Stable photo IDs
- WebP thumbnail generation
- Disk-backed thumbnail cache
- In-memory gallery index
- HTTP API
- Configurable via command-line flags

## Installation

```bash
git clone ...
cd galleryd

go build ./cmd/galleryd
```

## Usage

```bash
galleryd \
  --dir ~/Pictures \
  --dir ~/Phone \
  --cache-dir ~/.cache/galleryd \
  --addr :8082
```

## API

| Method | Endpoint    | Description     |
| ------ | ----------- | --------------- |
| GET    | /health     | Health check    |
| GET    | /api/photos | List all photos |
| GET    | /photo/{id} | Original image  |
| GET    | /thumb/{id} | Thumbnail       |

## Architecture

```
Scanner

↓

Metadata Reader

↓

Gallery Builder

↓

Gallery Service

↓

HTTP API
```

## Roadmap

- Gallery reload
- Filesystem watching
- Search
- Frontend

## License

MIT
