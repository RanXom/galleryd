# Shared Test Fixtures

This directory contains image fixtures shared across multiple packages.

## Canon_40D.jpg

Source: ianare/exif-samples

Contains:

- EXIF DateTimeOriginal
- EXIF Orientation
- Camera metadata

## Nikon_D70.jpg

Source: ianare/exif-samples

Contains:

- EXIF DateTimeOriginal
- Camera metadata

## no-exif.png

Generated using:

```bash
magick -size 512x512 xc:white "$OUT/no-exif.png"
```

Contains no EXIF metadata.
