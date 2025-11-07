# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A command-line tool for working with the 24 Elder Futhark runes, organized by their traditional three aetts (Freya's, Heimdall's, and Tyr's). The tool allows users to display all runes or search by keyword/name.

## Development Commands

### Build
```bash
go build -o runes
```

### Run
```bash
# Display all runes
./runes

# Search by keyword
./runes love
./runes protection

# Search by rune name
./runes fehu

# Search by symbol (paste UTF-8 runic character)
./runes ᚠ
./runes ᛉ
```

### Test Build
```bash
go build -o runes && ./runes
```

## Architecture

**Single-file architecture**: The entire application is contained in `main.go` with:

- **Data structures**:
  - `Rune` struct: Represents a single rune with Symbol, Name, Meaning, and Keywords
  - `Aett` struct: Groups 8 runes together with a name
  - Three package-level variables (`freyaAett`, `heimdallAett`, `tyrAett`) define all 24 runes with their metadata

- **Core functions**:
  - `getAllRunes()`: Flattens all runes from the three aetts into a single slice
  - `displayAllRunes()`: Displays all runes organized by aett with formatted output
  - `searchRunes(keyword)`: Case-insensitive search matching either rune names or keywords (supports partial matching)
  - `main()`: Entry point that routes to display or search based on command-line arguments

- **Search behavior**: The search is flexible and matches:
  - Exact UTF-8 rune symbol match (e.g., ᚠ, ᛉ, ᛞ) - checked first for exact match
  - Exact or partial rune names (case-insensitive)
  - Any keyword in the Keywords slice (using substring matching in both directions)
  - Multiple search terms are joined with spaces before matching

## Important Conventions

- **Unicode rune symbols**: Each rune uses its actual Elder Futhark Unicode character (e.g., ᚠ, ᚢ, ᚦ)
- **Keywords array**: Each rune has an extensive array of related keywords for flexible searching
- **Output formatting**: Uses box-drawing characters for visual hierarchy and readability
- **Aett organization**: Maintains the traditional three-aett structure (8 runes each) with thematic groupings
