# Project Design

## Overview

Cleanfreak is a cross-platform CLI tool written in Go that maintains a clean and orderly working environment. It monitors key user directories and provides commands to clean, list, and plan cleanup operations.

## CLI Structure

- `cleanfreak init` — Initialises a cleanfreak project directory with configured subdirectories.
- `cleanfreak diagnose` — Diagnoses the current system environment.
- `cleanfreak clean <item>` — Executes cleanup for a specified item.
  - `clean all` — Runs cleanup across all items.
  - `clean desktop` — Cleans the Desktop directory.
  - `clean downloads` — Cleans the Downloads directory.
  - `clean trash` — Empties the Trash.
  - `clean cache` — Removes cache files exceeding the configured threshold.
- `cleanfreak list <item>` — Lists the contents of a specified item.
- `cleanfreak plan <item>` — *(Planned, not yet implemented)* Constructs a deferred cleanup plan for selective execution later.

## Monitored Items

- **Desktop** — No files should live on the user's Desktop. They should be filed in an appropriate location.
- **Downloads** — No files should live in the user's Downloads folder. They should be filed in an appropriate location.
- **Trash** — The trash folder should be empty.
- **Cache** — No large and unnecessary cache files should be present. The user configures a size threshold (in GB) beyond which cache directories are flagged for cleanup.

## Design Principles

1. **User control** — The user decides when to initiate a cleanup operation. Cleanfreak never acts without explicit instruction.
2. **Non-destructive by default** — The `plan` command allows users to review and selectively approve actions before execution, rather than blindly deleting.
3. **Configurability** — Key settings are user-configurable, including the cleanfreak config directory path and the cache size threshold.

## Configuration

Cleanfreak stores its configuration in a user-specified directory. Configurable values include:

- **Config root** — The path where the cleanfreak configuration directory is created.
- **Cache size threshold** — The allowed cache size (in GB) before cleanup is recommended.
- **Subdirectories and cache directories** — Lists of directories to monitor.

## Platform Support

Cleanfreak targets cross-platform support. Notable caveats:

- **Windows** — Supported, with virtual terminal processing enabled for console output.
- **macOS and Linux** — Full support, including trash operations.
- **Trash operations** — Currently only supported on `darwin` and `linux`.

## Future Considerations

These tenets and features are open to negotiation, extension, and improvement.
