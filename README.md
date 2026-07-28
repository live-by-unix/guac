# Guac 🥑

Guac is a cross-platform CLI package manager wrapper that unifies APT, Homebrew, and winget under a single, playful interface. It provides consistent commands across Linux, macOS, and Windows, with platform-aware error handling and integrated logging.

## Features

- **Unified Interface**: Single command set for APT, Homebrew, and winget
- **Platform-Aware**: Automatically detects your platform and enables appropriate package managers
- **Consistent Commands**: Same syntax across all platforms
- **Integrated Logging**: Structured logging with logrus
- **Configurable**: Customizable via `~/.guac/config.json`
- **Cross-Platform**: Builds for Linux, macOS, and Windows

## Installation

### From Source

```bash
git clone https://github.com/live-by-unix/guac.git
cd guac
go build -o guac ./cmd
sudo mv guac /usr/local/bin/  # Linux/macOS
# or add to PATH on Windows
```

### From Homebrew (macOS/Linux)

```bash
brew tap live-by-unix/guac
brew install guac
```

### From Release Binaries

Download the latest release from [GitHub Releases](https://github.com/live-by-unix/guac/releases).

## Configuration

Guac creates a configuration file at `~/.guac/config.json` with default settings:

```json
{
  "preferred_manager": "brew",
  "use_sudo": true
}
```

- `preferred_manager`: Default package manager for your platform (apt/brew/winget)
- `use_sudo`: Whether to use sudo for operations (applies to APT)

## Usage

### Install a Package

```bash
guac install <package-manager> <package>

# Examples:
guac install apt nginx
guac install brew git
guac install winget VSCode
```

**Platform Rules:**
- APT: Always runs `sudo apt update` before install, always uses sudo
- Brew: Runs `brew install` directly
- Winget: Runs `winget install` with auto-accept flags

### List Installed Packages

```bash
guac list
```

Lists all installed packages for each supported package manager on your platform.

### Search for Packages

```bash
# Search across all managers
guac search <package>

# Search specific manager
guac search <package-manager> <package>

# Examples:
guac search python
guac search brew python
guac search apt python
```

### Remove a Package

```bash
guac remove <package-manager> <package>

# Examples:
guac remove apt nginx
guac remove brew git
guac remove winget VSCode
```

### Upgrade a Package

```bash
guac upgrade <package-manager> <package>

# Examples:
guac upgrade apt nginx
guac upgrade brew git
guac upgrade winget VSCode
```

### Version

```bash
guac --version
guac -v
```

### Help

```bash
guac --help
guac -h
```

## Platform Support

| Platform | Supported Managers |
|----------|-------------------|
| Linux    | APT, Brew         |
| macOS    | Brew              |
| Windows  | Winget            |

If you try to use an unsupported package manager on your platform, you'll see:
```
Not supported for your platform.
```

## Development

### Prerequisites

- Go 1.26+
- Make (optional)

### Building

```bash
go build -o guac ./cmd
```

### Running Tests

```bash
go test ./...
```

### Building with GoReleaser

```bash
goreleaser build --snapshot
```

## CI/CD

Guac uses GitHub Actions for:
- **CI**: Automated testing on Linux, macOS, and Windows
- **Release**: Automated releases via GoReleaser when tags are pushed

## Roadmap

- [ ] Add support for additional package managers (yum, dnf, pacman)
- [ ] Interactive mode for package selection
- [ ] Package dependency visualization
- [ ] Configuration wizard
- [ ] Shell completion (bash, zsh, fish, powershell)

## License

MIT License - see LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
