# Journal

Producdive terminal journaling.

## Install

Install from a release binary or build from source.

## Config

```sh
# Journal path defaults to "$HOME/.journal". Override with:
export JOURNAL_PATH="$HOME/MyJournal"

# Set the editor command
export EDITOR="code -w"
```

> You should add this configuration to your `.profile`/`.bashrc`/`.zshrc`/etc to make it perminent.

### Seperate Journals

You can have multiple journals by setting aliases along with your [environment settings](#config):

```sh
alias work-journal='journal -D "$HOME/MyWorkJournal"'
```

### Multi Machine Syncing

Journal does not support file sync directly, but is built to be friendly to
other file-syncing solutions. To sync your journal between machines use a
file sync service (some popular ones mentioned below) and sync your full
`$JOURNAL_PATH`.

- [Sync Thing](https://syncthing.net)
- [Mega Sync](https://mega.io/sync)
- [Microsoft OneDrive](https://www.microsoft.com)

## Use

```
Journal is an oppinionated journaling utility for keeping daily
and archival notes - all from your terminal.

Usage:
  journal [command]

Available Commands:
  archives    List archive files.
  edit        Edit a journal page.
  help        Help about any command
  index       Re-index all journal pages for searching.
  info        Show info/stats about your journal.
  memo        Add a memo to a journal page with a timestamp
  search      Run a search across all journal pages.
  version     Print version info
  view        Read journal pages.

Flags:
  -h, --help                  help for journal
  -D, --journal-path string   Directory journal entries are stored in.

Use "journal [command] --help" for more information about a command.
```