
# zk

A simple command-line tool for taking notes stored in plain text files.

## Install

1. Build from source for now.

2. Create zk.conf specify notesdir, see Configuration below

## Usage

### Creating notes

1. Use `zk` opens new note in editor (default: vim)

2. Add note inline as a command-line argument
```
$ zk "Here is my note"
```

3. Use STDIN by piping or slurping in data
```
$ git log -1 | zk
$ zk < file-to-import.txt
```

You can create a tag using @tag which can then be used to search later.
```
$ zk "Here is my note with a @tag"
```

### Viewing Notes

Show the past week of notes
```
$ zk --week
```

Show today's notes
```
$ zk --today
```

Show a specific date's notes
```
$ zk --date 2017-10-15
```

Search by term
```
$ zk -s searchterm
```

Use @tag to search. If you only include tags on command-line, then zk assumes you are searching by tag. You can include multiple tags and it will OR search for all of them.
```
$ zk @tag
```

### Editing zks

Edit today's note
```
$ zk --today --edit
```

Edit a specific date
```
$ zk --date 2017-11-15 --edit
```

## Configuration

zk creates a single text file a day, `zk-2017-01-06.txt` organizing them in monthly and yearly folders. You need to set the `notesdir` in the `zk.conf` config file.

The config file location be be specified various ways, zk will look for the following to determine where the config file is:

1. Command-line argument. Use `zk --config /path/to/zk.conf`

2. Environment variable: `ZK_CONFIG_FILE`

3. Look for `${XDG_CONFIG_HOME}/zk.conf`

4. Look for `${HOME}/.config/zk.conf`

If not specified or found in any of the above locations, zk will error out with a message on to set the configuration file.

The config file is in TOML format, example:

```toml
# zk config file

# base directory that all zks are stored
notes_dir = '~/Documents/Zks'

# What timestamp to use for default notes
# This allow you to create monthly, weekly, or daily notes.
# See: https://docs.rs/chrono/0.4.0/chrono/format/strftime/index.html
filename_format = '%y%m%d%H%M.md'
```

## Errata

**Syncing data** - You can use a cloud or sync provider to keep your notes in-sync across computers. I'm happy with the open-source [Syncthing](https://syncthing.net/) to keep my Documents folder in sync across computers. I hear Google Drive, Dropbox, and others might be able to do the same.

## History

A bit of a toy project, a place to apply programming languages I'm learning, so it has gone through a couple iterations.

The utlity was initially called **jot** but to reduce one character and hop on to the Zettlekasten train, I changed it to **zk**.

Plus it was first created in Python ([branch](https://github.com/mkaz/zk/tree/python)) and then rewritten in Go ([branch](https://github.com/mkaz/zk/tree/golang)).  Now, it is Rust's turn.

## License

zk is open source and free to use, it is licensed under a <a rel="license" href="https://opensource.org/licenses/MIT">MIT License</a>.

I welcome any contributions. Please feel free to open an issues to report a bug, submit a feature. Pull requests are also welcome.

An [mkaz](https://mkaz.blog/) contrivance.

