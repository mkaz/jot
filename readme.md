# jot

A simple command-line tool for taking notes stored in plain text files.

## Install

1. Build from source.

2. Create jot.conf, see [Configuration below](#configuration)

    - `default_path` is only required config

## Usage

See `jot --help` for help.

### Creating notes

1. Use `jot` opens new note in EDITOR (default: vim)

2. Add note inline as a command-line argument

```
$ jot "Here is my note"
```

3. Use STDIN by piping or slurping in data

```
$ git log -1 | jot
$ jot < file-to-import.txt
```

### Types of Notes

There are five types of notes: default, monthly, weekly, daily, new. This is a means for filing notes for different uses. I use the weekly notes for my weekly status updates. My monthly notes are for on-going projects through the month. I use daily or new note types for one-offs.

The note type is specified by using the command-line flag, the default note is when no flag is present.

| Flag      | Config Name    | Format String   | Example           |
| --------- | -------------- | --------------- | ----------------- |
| --monthly | monthly_format | %Y-%m-%b.md     | 2021-02-Feb.md    |
| --weekly  | weekly_format  | %Y-week-%U.md   | 2021-week-08.md   |
| --daily   | daily_format   | %Y-%m-%d.md     | 2021-02-25.md     |
| --new     | new_format     | %Y%m%d%H%M%S.md | 20210225153423.md |
|           | default_format | %Y%m%d%H%M.md   | 202102251534.md   |

See https://docs.rs/chrono-wasi/0.4.10/chrono/format/strftime/[Chrono strftime documentation] for format parameters.

Each note type can also have its own path, see config below. If the file does not exist, it will be created, otherwise opened or appended to.

## Configuration

You need to set the `default_path` in the `jot.conf` config file.

The config file location can be specified various ways, `jot` will look in the following order to determine where the config file is:

1. Command-line argument. Use `jot --config /path/to/jot.conf`

2. Environment variable: `JOT_CONFIG_FILE`

3. Look for platform config directory

a. Linux: `${XDG_CONFIG_HOME}/jot.conf`
b. Windows: `${APPDATA}/jot.conf`

4. Look for `${HOME}/.config/jot.conf`

If not specified or found in any of the above locations, `jot` will error out with a message to set the configuration file.

The config file is in TOML format, example:

```toml

# jot config file

# base directory that all notes are stored
default_path = '~/Documents/jot'

# What format to use for notes
# See: https://docs.rs/chrono/0.4.0/chrono/format/strftime/index.html
default_format = '%Y%m%d%H%M.md'

# 2021-02-Feb.md
monthly_format = "%Y-%m-%b.md"
# monthly_path = '/full/path/to/monthly/notes'

# 2021-week-08.md
weekly_format = "%Y-week-%U.md"
# weekly_path = '/full/path/to/weekly/notes'

# 2021-02-25.md
daily_format = "%Y-%m-%d.md"
# daily_path = '/full/path/to/daily/notes'

# 20210225153423.md
new_format = "%Y%m%d%H%M%S.md"
# new_path = '/full/path/to/new/notes'
```

## Errata

**Syncing data** - You can use a cloud or sync provider to keep your notes in-sync across computers. I'm happy with the open-source https://syncthing.net/[Syncthing] to keep my Documents folder in sync across computers. I hear Google Drive, Dropbox, and others might be able to do the same.

## History

A bit of a toy project, a place to apply programming languages I'm learning, so it has gone through a couple iterations. It was first created in Python (https://github.com/mkaz/jot/tree/python[branch]) and then rewritten in Go (https://github.com/mkaz/jot/tree/golang[branch]). Now, it is Rust's turn.

## License

Jot is open-source and free to use, it is licensed under a [MIT License](https://opensource.org/licenses/MIT)

I welcome contributions, but as a side project I may not always respond promptly. Please feel free to open an issues to report a bug, submit a feature, or even a pull request.
