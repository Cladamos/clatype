# clatype

A terminal-based typing speed test written in Go using [Bubbletea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss).

## Installation

Make sure you have Go installed (>=1.20).

```bash
go install github.com/cladamos/clatype@latest
```

## Usage

Run the typing test with default settings (English words, 30 seconds):

```bash
clatype
```

### Command-Line Flags

- **`-l <language>`** - Set the language/mode for typing practice
  - `english` or `en` - English words (default)
  - `go` or `golang` - Go code snippets
  - `javascript` or `js` - JavaScript code snippets

- **`-t <duration>`** - Set the test duration (default: 30s)
  - Examples: `10s`, `60s`, `2m`, `1m30s`

### Examples

Test with Go snippets for 60 seconds:

```bash
clatype -l go -t 60s
```

Test with JavaScript snippets for 2 minutes:

```bash
clatype -l javascript -t 2m
```

Quick 10-second English test:

```bash
clatype -t 10s
```
