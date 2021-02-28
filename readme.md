
# zk

A simple command-line tool for taking notes stored in plain text files.

## Install

1. Build from source.

2. Create zk.conf specify `notes_dir`, see [Configuration below](#configuration)

## Usage

See `zk --help` for help.

### Creating notes

1. Use `zk` opens new note in EDITOR (default: vim)

2. Add note inline as a command-line argument
```
$ zk "Here is my note"
```

3. Use STDIN by piping or slurping in data
```
$ git log -1 | zk
$ zk < file-to-import.txt
```

### Types of Notes

There are five types of notes: default, monthly, weekly, daily, new.

The note type is specified by using the command-line flag,  the default note is when no flag is present.

| Flag      | Config Name       | Format String   | Example           |
|-----------|-------------------|-----------------|-------------------|
| --monthly | monthly_format	| %Y-%m-%b.md     | 2021-02-Feb.md    |
| --weekly  | weekly_format     | %Y-week-%U.md   | 2021-week-08.md   |
| --daily   | daily_format      | %Y-%m-%d.md     | 2021-02-25.md     |
| --new     | new_format        | %Y%m%d%H%M%S.md | 20210225153423.md |
|           | default_format    | %Y%m%d%H%M.md   | 202102251534.md   |

See [Chrono strftime documentation](https://docs.rs/chrono-wasi/0.4.10/chrono/format/strftime/index.html) for format parameters.

If the file does not exist, it will be created, otherwise opened or appended to.

## Configuration

You need to set the `notes_dir` in the `zk.conf` config file.

The config file location be be specified various ways, zk will look for the following to determine where the config file is:

1. Command-line argument. Use `zk --config /path/to/zk.conf`

2. Environment variable: `ZK_CONFIG_FILE`

3. Look for platform config directory
a) Linux: `${XDG_CONFIG_HOME}/zk.conf`
b) Windows: `${APPDATA}/zk.conf`

4. Look for `${HOME}/.config/zk.conf`

If not specified or found in any of the above locations, zk will error out with a message on to set the configuration file.

The config file is in TOML format, example:

```toml
# zk config file

# base directory that all zks are stored
notes_dir = '~/Documents/Zettelkasten'

# What format to use for notes
# See: https://docs.rs/chrono/0.4.0/chrono/format/strftime/index.html
default_format = '%Y%m%d%H%M.md'

# 2021-02-Feb.md
monthly_format="%Y-%m-%b.md"

# 2021-week-08.md
weekly_format="%Y-week-%U.md"

# 2021-02-25.md
daily_format="%Y-%m-%d.md"

# 20210225153423.md
new_format="%Y%m%d%H%M%S.md"
```

## Errata

**Syncing data** - You can use a cloud or sync provider to keep your notes in-sync across computers. I'm happy with the open-source [Syncthing](https://syncthing.net/) to keep my Documents folder in sync across computers. I hear Google Drive, Dropbox, and others might be able to do the same.

## History

A bit of a toy project, a place to apply programming languages I'm learning, so it has gone through a couple iterations.  It was first created in Python ([branch](https://github.com/mkaz/zk/tree/python)) and then rewritten in Go ([branch](https://github.com/mkaz/zk/tree/golang)).  Now, it is Rust's turn.

The utlity was renamed to **zk** during the Rust rewrite, it was previously  called **jot**. I started creating a new command-line Zettlekasten tool and it was basically the same as my existing utility, so I merged into one and I prefer the shorter **zk** command.

## License

zk is open source and free to use, it is licensed under a <a rel="license" href="https://opensource.org/licenses/MIT">MIT License</a>.

I welcome any contributions. Please feel free to open an issues to report a bug, submit a feature. Pull requests are also welcome.

An [mkaz](https://mkaz.blog/) contrivance.

