# Journal

A unix-based jornal.

## Install

Install from a release binary or build from source.

## Config

Copy and update the following file to `$HOME/.config/journal/config.yaml`

```yaml
editor: "code -w"
journalPath: "/Users/ME/.journal"
```

## Use

```
Journal is a fully featured journal for keeping daily and archival notes, all from your terminal.

Usage:
  journal [command]

Available Commands:
  archives    List archive files.
  edit        Edit a journal page.
  help        Help about any command
  index       Re-index all journal pages for searching.
  memo        Add a memo to a journal page with a timestamp
  search      Run a search across all journal pages.
  version     Print version info
  view        Read journal pages.

Flags:
  -h, --help   help for journal

Use "journal [command] --help" for more information about a command.
```