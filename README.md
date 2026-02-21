# clipwipe

Clipboard URL cleaner - automatically removes UTM parameters and other query parameters from URLs in your clipboard.

## Features

- Monitors clipboard and automatically cleans URLs
- Removes tracking parameters by default (`utm_source`, `utm_medium`, `utm_campaign`, `utm_content`, `utm_term`, `fbclid`, `gclid`)
- Configurable parameter list via `-params` flag
- Configurable polling interval via `-interval` flag
- Debug mode for troubleshooting

## Installation

### Via `go install` (recommended)

```bash
go install github.com/damirm/clipwipe@latest
```

The binary will be installed to `$GOPATH/bin` (or `$HOME/go/bin` by default).

### From source

```bash
go build -o clipwipe
```

## Usage

```bash
# Run with defaults (500ms interval, UTM parameters)
clipwipe

# Custom interval
clipwipe -interval 100ms

# Custom parameters to remove
clipwipe -params "utm_source,utm_medium,ref"

# Debug mode
clipwipe -debug -interval 100ms
```

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-interval` | `500ms` | Clipboard polling interval |
| `-params` | `utm_source,utm_medium,utm_campaign,utm_content,utm_term,fbclid,gclid` | Comma-separated list of query parameters to remove |
| `-debug` | `false` | Enable debug output |

## Examples

```bash
# Remove only utm_source and utm_medium
clipwipe -params "utm_source,utm_medium"

# Remove tracking parameters (default)
clipwipe

# Fast polling for responsive cleaning
clipwipe -interval 50ms
```

## How it works

1. Program polls the clipboard at the specified interval
2. When content changes, it checks if it's a URL with `http://` or `https://` prefix
3. If URL contains specified query parameters, they are removed
4. Cleaned URL is written back to clipboard

## Requirements

- Go 1.24+
- Cross-platform: Windows, macOS, Linux
- **Linux/Wayland users**: Requires XWayland. Make sure `DISPLAY` is set or run in an X11 session.

## License

MIT
